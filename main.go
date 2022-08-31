package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"real-time-forum/db"

	"github.com/gorilla/websocket"
)

type T struct {
	TypeChecker
	*Posts
	*Comments
	*Register
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
	User      string `json:"username"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Tipo      string `json:"tipo"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcastChannelPosts = make(chan *Posts, 1)
var broadcastChannelComments = make(chan *Comments, 1)
var broadcastChannelRegister = make(chan *Register, 1)


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
	case "register":
		t.Register = &Register{}
		return json.Unmarshal(data, t.Register)
	default:
		return fmt.Errorf("unrecognized type value %q", t.Type)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func webSocketEndpoint(w http.ResponseWriter, r *http.Request) {
	go broadcastToAllClients()
	wsConn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("error when upgrading connection...")
	}

	defer wsConn.Close()

	clients[wsConn] = true

	for {
		_, infoType, _ := wsConn.ReadMessage()
		fmt.Println("----", string(infoType))

		var f T
		f.UnmarshalForumData(infoType)

		if f.Type == "post" {
			f.Posts.Tipo = "post"
			f.Posts.User = "yonas"
			broadcastChannelPosts <- f.Posts
		} else if f.Type == "comment" {
			f.Comments.User = "yonas"
			f.Comments.Tipo = "comment"
			broadcastChannelComments <- f.Comments
		} else if f.Type == "register" {
			fmt.Println("-----", f.Register)
			f.Register.User = "tols"
			f.Register.Tipo = "registration"
			
			//below solely for testing
			broadcastChannelRegister <- f.Register

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

func main() {
	db.CreateDB()
	// need to figure out how to serve these files on the homepage (/)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", webSocketEndpoint)
	log.Println("Listening on port :8080.....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
