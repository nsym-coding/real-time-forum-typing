package db

import (
	"database/sql"
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

	db.Exec("create table if not exists users (userID integer primary key AUTOINCREMENT, email text, username text, hash CHAR(60))")
	db.Exec(`create table if not exists posts (
			postID integer primary key AUTOINCREMENT, 
			userID integer REFERENCES users(userID), 
			creationDate integer,
			postTitle CHAR(50),
			postContent CHAR(250), 
			image CHAR(100), 
			edited integer);`)
	db.Exec(`create table if not exists comments (
			commentID integer primary key AUTOINCREMENT, 
			userID integer REFERENCES users(userID), 
			postID integer REFERENCES post(postID), 
			commentText CHAR(250), 
			edited integer, 
			creationDate integer,
			notified integer,
			creatorID integer);`)

	db.Exec(`create table if not exists messages(messageID integer PRIMARY KEY AUTOINCREMENT, message text,
					sender text REFERENCES users(username), recepient text REFERENCES users(username) creationDate integer);`)
}
