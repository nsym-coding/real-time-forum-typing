package users

import (
	"database/sql"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var LoggedInUsers = make(map[string]string)

// this func registers a users username, email, firstname, lastname, password(unhashed) and age
func RegisterUser(db *sql.DB, username string, age string, gender string, firstname string, lastname string, hash []byte, email string) {
	// db, _ := sql.Open("sqlite3", "real-time-forum.db")
	stmt, err := db.Prepare("INSERT INTO users (username, age, gender, firstname, lastname, hash, email) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()
	result, _ := stmt.Exec(username, age, gender, firstname, lastname, hash, email)
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
	//get user from db
	userStmt := "SELECT hash from users WHERE username = ?"
	rowU := db.QueryRow(userStmt, username)
	var hash string
	err := rowU.Scan(&hash)
	if err != nil {
		fmt.Println("Error in finding hash, ", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func CreateCookie(writer http.ResponseWriter, username string) {

	//NO ACTIVE SESSION/FIRST TIME
	id := uuid.Must(uuid.NewV4())
	c := &http.Cookie{
		Name:  username,
		Value: id.String(),
	}
	http.SetCookie(w, c)
	users.DbSessions[username] = c.Value
	// tpl.ExecuteTemplate(w, "loginauth.html", userID)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
	/////////remove///////////////////
	fmt.Println("sessionbool", users.SessionExists(username))
	for _, cookie := range r.Cookies() {
		fmt.Println()
		fmt.Println("Name : ", cookie.Name)
		fmt.Println("Value/UUID : ", cookie.Value)
	}
	fmt.Println("First time log-in successful")
	fmt.Println()
	/////////////////////////////////////
	return

}
