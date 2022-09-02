package users

import (
	"database/sql"
	"fmt"
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
