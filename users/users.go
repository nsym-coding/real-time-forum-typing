package users

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var LoggedInUsers = make(map[string]string)

type AllUsers struct {
	Username string `json:"user"`
	Team     string `json:"team"`
}

// this func registers a users username, email, firstname, lastname, password(unhashed) and age
func RegisterUser(db *sql.DB, username string, age string, gender string, firstname string, lastname string, hash []byte, email string, team string) {
	// db, _ := sql.Open("sqlite3", "real-time-forum.db")
	stmt, err := db.Prepare("INSERT INTO users (username, age, gender, firstname, lastname, hash, email, team) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()
	result, _ := stmt.Exec(username, age, gender, firstname, lastname, hash, email, team)
	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", lastIns)
}

func UserExists(db *sql.DB, username string) bool {
	// check if username already exists
	userStmt := "SELECT userID FROM users WHERE username = ?"
	rowU := db.QueryRow(userStmt, username)
	var uIDs string
	error := rowU.Scan(&uIDs)
	if error != sql.ErrNoRows {
		fmt.Println("username already exists, err:", error)
		return true
	}
	return false
}

func EmailExists(db *sql.DB, email string) bool {
	userStmt := "SELECT userID FROM users WHERE email = ?"
	rowU := db.QueryRow(userStmt, email)
	var uIDs string
	error := rowU.Scan(&uIDs)
	if error != sql.ErrNoRows {
		fmt.Println("email already exists, err:", error)
		return true
	}
	return false
}

func CorrectPassword(db *sql.DB, username, password string) bool {
	// get user from db
	userStmt := "SELECT hash from users WHERE username = ? OR email = ?"
	rowU := db.QueryRow(userStmt, username, username)
	var hash string
	err := rowU.Scan(&hash)
	if err != nil {
		fmt.Println("Error in finding hash, ", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func CreateCookie(writer http.ResponseWriter, req *http.Request, username string) {
	// NO ACTIVE SESSION/FIRST TIME
	id := uuid.Must(uuid.NewV4())
	c := &http.Cookie{
		Name:  username,
		Value: id.String(),
	}
	http.SetCookie(writer, c)
	LoggedInUsers[username] = c.Value

	fmt.Println(LoggedInUsers)
	/////////////////////////////////////
}

func ValidEmail(email string) bool {
	i := strings.Index(email, "@")
	fmt.Println("i:", i)
	domain := email[i+1:]
	fmt.Println("Domain: ", domain)
	_, err := net.LookupMX(domain)
	// , err2 := mail.ParseAddress(email)
	if err != nil {
		fmt.Println("invalid email")
		return false
	}
	return true
}

func GetUserID(db *sql.DB, username string) int {
	// check if username already exists
	userStmt := "SELECT userID FROM users WHERE username = ?"
	rowU := db.QueryRow(userStmt, username)
	var uIDs int
	error := rowU.Scan(&uIDs)
	if error != sql.ErrNoRows {
		fmt.Println("username already exists, err:", error)
		// intUID, _ := strconv.Atoi(uIDs)
		return uIDs
	}
	return 0
}

// get all Users from DB for chat
func GetAllUsers(db *sql.DB) []AllUsers {
	var usernameBadges []AllUsers

	rows, err := db.Query(`SELECT username, team
    FROM users`)
	if err != nil {
		fmt.Println("Error with GetAllUsers func")
	}

	var allUsers AllUsers

	defer rows.Close()

	for rows.Next() {
		var user string
		var team string

		err2 := rows.Scan(&user, &team)

		allUsers.Username = user
		allUsers.Team = team

		usernameBadges = append(usernameBadges, allUsers)

		if err2 != nil {
			fmt.Println("Error appending users GetAllUsers()", err2)
		}
	}

	return usernameBadges
}

func GetUserName(db *sql.DB, email string) string {
	// check if username already exists
	userStmt := "SELECT username FROM users WHERE email = ?"
	rowU := db.QueryRow(userStmt, email)
	var uIDs string
	error := rowU.Scan(&uIDs)
	if error != sql.ErrNoRows {
		fmt.Println("username already exists, err:", error)
		// intUID, _ := strconv.Atoi(uIDs)
		return uIDs
	}
	return email
}
