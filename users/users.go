package users

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

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
