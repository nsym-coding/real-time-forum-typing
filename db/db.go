package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "real-time-forum.db"

func CreateDB() {
	var db *sql.DB
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {

		log.Fatal(err)
	}

	_, err1 := db.Exec(`create table if not exists users (
		userID integer primary key AUTOINCREMENT, 
		username CHAR(50), 
		age integer,
		email CHAR(50), 
		gender CHAR(50),
		firstname CHAR(50), 
		lastname CHAR(50), 
		hash CHAR(50)
		);`)
	fmt.Println("err1", err1)

	_, err2 := db.Exec(`create table if not exists posts (
			postID integer primary key AUTOINCREMENT, 
			username CHAR(50) REFERENCES users(username), 
			creationDate integer,
			postTitle CHAR(50),
			categories CHAR(50),
			postContent CHAR(250));`)
	fmt.Println("err2", err2)

	_, err3 := db.Exec(`create table if not exists comments (
			commentID integer primary key AUTOINCREMENT, 
			username CHAR(50) REFERENCES users(username), 
			postID integer REFERENCES post(postID), 
			commentText CHAR(250), 
			edited integer, 
			creationDate integer,
			notified integer,
			creatorID integer);`)
	fmt.Println("err3", err3)

	_, err4 := db.Exec(`create table if not exists messages(
		messageID integer PRIMARY KEY AUTOINCREMENT, 
		chatID integer REFERENCES chats(chatID), 
		message CHAR(250),
	    sender text REFERENCES users(username), 
		recepient text REFERENCES users(username),
		creationDate integer);`)
	fmt.Println("err4", err4)

	_, err5 := db.Exec(`create table if not exists chats(
		chatID integer PRIMARY KEY AUTOINCREMENT, 
	    user1 text REFERENCES users(username), 
		user2 text REFERENCES users(username),
		creationDate integer);`)
	fmt.Println("err5", err5)
}
