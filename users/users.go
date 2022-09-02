package users

import (
	"fmt"
	"real-time-forum/server"
)

// this func registers a users username, email, firstname, lastname, password(unhashed) and age
func (s *MyServer) RegisterUser(username string, age int, firstname string, lastname string, hash []byte, email string) {
	// db, _ := sql.Open("sqlite3", "real-time-forum.db")
	stmt, err := s.Db.Prepare("INSERT INTO users (username, age, firstname, lastname, hash, email) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()
	result, _ := stmt.Exec(username, age, firstname, lastname, hash, email)
	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", lastIns)
}
