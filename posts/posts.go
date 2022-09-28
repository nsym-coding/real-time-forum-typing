package posts

import (
	"database/sql"
	"fmt"
)

func StorePosts(db *sql.DB, userID int, title string, content string) {
	stmt, err := db.Prepare("INSERT INTO posts (userID, postTitle, postContent, creationDate) VALUES (?, ?, ?, strftime('%H:%M %d/%m/%Y','now','localtime'))")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()
	result, _ := stmt.Exec(userID, title, content)
	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	LastIns, _ := result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", LastIns)
}
