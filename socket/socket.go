package socket

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"real-time-forum/users"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/bcrypt"
)

type T struct {
	TypeChecker
	*Posts
	*Comments
	*Register
	*Login
}

type TypeChecker struct {
	Type string `json:"type"`
}

type Posts struct {
	Title       string `json:"title"`
	PostContent string `json:"postcontent"`
	Date        string `json:"posttime"`
	Tipo        string `json:"tipo"`
	User        string `json:"user"`
}

type Comments struct {
	CommentContent string `json:"commentcontent"`
	User           string `json:"user"`
	Date           string `json:"commenttime"`
	Tipo           string `json:"tipo"`
}

type Register struct {
	Username  string `json:"username"`
	Age       string `json:"age"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Tipo      string `json:"tipo"`
}

type Login struct {
	LoginUsername string `json:"loginUsername"`
	LoginPassword string `json:"loginPassword"`
	Tipo          string `json:"tipo"`
}

type formValidation struct {
	UsernameLength         bool   `json:"usernameLength"`
	UsernameSpace          bool   `json:"usernameSpace"`
	UsernameDuplicate      bool   `json:"usernameDuplicate"`
	EmailDuplicate         bool   `json:"emailDuplicate"`
	PasswordLength         bool   `json:"passwordLength"`
	AgeEmpty               bool   `json:"ageEmpty"`
	FirstNameEmpty         bool   `json:"firstnameEmpty"`
	LastNameEmpty          bool   `json:"lastnameEmpty"`
	EmailInvalid           bool   `json:"emailInvalid"`
	SuccessfulRegistration bool   `json:"successfulRegistration"`
	Tipo                   string `json:"tipo"`
}
type loginValidation struct {
	InvalidUsername bool   `json:"invalidUsername"`
	InvalidPassword bool   `json:"invalidPassword"`
	SuccessfulLogin bool   `json:"successfulLogin"`
	Tipo            string `json:"tipo"`
}

var (
	clients                  = make(map[*websocket.Conn]bool)
	loggedInUsers            = make(map[string]*websocket.Conn)
	broadcastChannelPosts    = make(chan *Posts, 1)
	broadcastChannelComments = make(chan *Comments, 1)
	broadcastChannelRegister = make(chan *Register, 1)
)

// unmarshall data based on type
func (t *T) UnmarshalForumData(data []byte) error {
	if err := json.Unmarshal(data, &t.TypeChecker); err != nil {
		log.Println("Error when trying to sort forum data type...")
	}

	switch t.Type {
	case "post":
		t.Posts = &Posts{}
		return json.Unmarshal(data, t.Posts)
	case "comment":
		t.Comments = &Comments{}
		return json.Unmarshal(data, t.Comments)
	case "signup":
		t.Register = &Register{}
		return json.Unmarshal(data, t.Register)
	case "login":
		t.Login = &Login{}
		return json.Unmarshal(data, t.Login)
	default:
		return fmt.Errorf("unrecognized type value %q", t.Type)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocketEndpoint(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "real-time-forum.db")

	go broadcastToAllClients()
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error when upgrading connection...")
	}
	fmt.Println("CONNECTION TO CLIENT")
	defer wsConn.Close()

	clients[wsConn] = true

	for {
		_, info, _ := wsConn.ReadMessage()
		fmt.Println("----", string(info))

		var f T
		f.UnmarshalForumData(info)

		if f.Type == "post" {
			f.Posts.Tipo = "post"
			fmt.Println("this is the post content       ", f.PostContent)
			// f.Posts.User = "yonas"
			broadcastChannelPosts <- f.Posts
		} else if f.Type == "comment" {
			// f.Comments.User = "yonas"
			f.Comments.Tipo = "comment"
			broadcastChannelComments <- f.Comments
		} else if f.Type == "register" {
			fmt.Println("-----", f.Register.Age)

			// var u formValidation
			// u.Tipo = "formValidation"
			// canRegister := true

			// if users.UserExists(db, f.Username) {
			// 	u.UsernameDuplicate = true
			// 	canRegister = false

			// }

			// if users.EmailExists(db, f.Email) {
			// 	u.EmailDuplicate = true
			// 	canRegister = false

			// }

			// if len(f.Password) < 5 {
			// 	u.PasswordLength = true
			// 	canRegister = false

			// }
			//wsConn.WriteJSON(u)

			// f.Register.Username = "tols"
			f.Register.Tipo = "registration"

			// below solely for testing
			broadcastChannelRegister <- f.Register

		} else if f.Type == "login" {
			var loginData loginValidation
			loginData.Tipo = "loginValidation"

			if !users.UserExists(db, f.Login.LoginUsername) {
				fmt.Println("Checking f.login.loginusername --> ", f.Login.LoginUsername)
				loginData.InvalidUsername = true
				wsConn.WriteJSON(loginData)
			} else if users.UserExists(db, f.Login.LoginUsername) {
				if !users.CorrectPassword(db, f.Login.LoginUsername, f.Login.LoginPassword) {
					loginData.InvalidPassword = true
					wsConn.WriteJSON(loginData)
				} else {
					loginData.SuccessfulLogin = true
					loggedInUsers[f.Login.LoginUsername] = wsConn
					fmt.Println(loggedInUsers)
					wsConn.WriteJSON(loginData)
					fmt.Println("SUCCESSFUL LOGIN")
				}
			}

			// Check username exists
			// Check the password matches
		}

		log.Println("Checking what's in f ---> ", f)
	}
}

func broadcastToAllClients() {
	for {
		select {
		case x, ok := <-broadcastChannelPosts:
			if ok {
				for client := range clients {
					client.WriteJSON(x)
					fmt.Printf("Value %v was read.\n", x)
				}
			}
		case y, ok := <-broadcastChannelComments:
			if ok {
				for client := range clients {
					client.WriteJSON(y)
				}
			}
		case z, ok := <-broadcastChannelRegister:
			if ok {
				for client := range clients {
					client.WriteJSON(z)
				}
			}
		}
	}
}

func GetLoginData(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "real-time-forum.db")
	fmt.Println(r.Method)

	var t T

	data, _ := io.ReadAll(r.Body)

	t.UnmarshalForumData(data)

	if t.Type == "signup" {

		var u formValidation
		u.Tipo = "formValidation"
		canRegister := true

		if len(t.Register.Username) < 5 {
			u.UsernameLength = true
			canRegister = false
		}

		intAge, _ := strconv.Atoi(t.Register.Age)
		if intAge < 16 {
			fmt.Println(t.Register.Age)
			fmt.Println("age invalid")
			u.AgeEmpty = true
			canRegister = false
		}
		if t.Register.FirstName == "" {
			fmt.Println("first name empty")
			u.FirstNameEmpty = true
			canRegister = false
		}
		if t.Register.LastName == "" {
			fmt.Println("last name empty")
			u.LastNameEmpty = true
			canRegister = false
		}

		if len(t.Register.Password) < 5 {
			u.PasswordLength = true
			canRegister = false
		}

		if strings.Contains(t.Register.Username, " ") {
			u.UsernameSpace = true
			canRegister = false
		}

		if len(t.Register.Password) < 5 {
			u.PasswordLength = true
			canRegister = false
		}

		if !users.ValidEmail(t.Register.Email) {
			u.EmailInvalid = true
			canRegister = false
		}
		if users.UserExists(db, t.Register.Username) {
			u.UsernameDuplicate = true
			canRegister = false
		}

		if users.EmailExists(db, t.Register.Email) {
			u.EmailDuplicate = true
			canRegister = false
		}

		// all validations passed
		if canRegister {
			// hash password
			var hash []byte
			hash, err := bcrypt.GenerateFromPassword([]byte(t.Password), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("bcrypt err:", err)
			}
			users.RegisterUser(db, t.Username, t.Register.Age, t.Gender, t.FirstName, t.LastName, hash, t.Email)

			// data gets marshalled and sent to client
			u.SuccessfulRegistration = true
			toSend, _ := json.Marshal(u)
			fmt.Println("toSend -- > ", toSend)
			w.Write(toSend)
			//	http.HandleFunc("/ws", WebSocketEndpoint)
		} else {

			toSend, _ := json.Marshal(u)
			w.Write(toSend)
		}
	}

	if t.Type == "login" {
		//validate values then
		var loginData loginValidation

		loginData.Tipo = "loginValidation"

		if !users.UserExists(db, t.Login.LoginUsername) {
			fmt.Println("Checking f.login.loginusername --> ", t.Login.LoginUsername)
			loginData.InvalidUsername = true
			toSend, _ := json.Marshal(loginData)
			w.Write(toSend)

		} else if users.UserExists(db, t.Login.LoginUsername) {
			fmt.Println("user exists")
			if !users.CorrectPassword(db, t.Login.LoginUsername, t.Login.LoginPassword) {
				loginData.InvalidPassword = true
				toSend, _ := json.Marshal(loginData)
				w.Write(toSend)

			} else {
				loginData.SuccessfulLogin = true
				toSend, _ := json.Marshal(loginData)
				w.Write(toSend)
				http.HandleFunc("/ws", WebSocketEndpoint)
				//loggedInUsers[f.Login.LoginUsername] = wsConn
				fmt.Println(loggedInUsers)

				fmt.Println("SUCCESSFUL LOGIN")
			}

			// Check username exists
			// Check the password matches
		}

		// data gets marshalled and sent to client

	}

}
