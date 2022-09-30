package comments

import (
	"database/sql"
	"fmt"
)

type Comments struct {
	CommentContent string `json:"commentcontent"`
	User           string `json:"user"`
	Date           string `json:"commenttime"`
	Tipo           string `json:"tipo"`
	PostID         string `json:"postid"`
}

type CommentsFromPosts struct {
	ClickedPostID string `json:"clickedPostID"`
	Tipo          string `json:"getcommentsfrompost"`
}

// commentID integer primary key AUTOINCREMENT,
// 			username CHAR(50) REFERENCES users(username),
// 			postID integer REFERENCES post(postID),
// 			commentText CHAR(250),
// 			creationDate integer);`)

func StoreComment(db *sql.DB, user string, postID int, commentContent string) {
	stmt, err := db.Prepare("INSERT INTO comments (username, postID, commentText, creationDate) VALUES (?, ?, ?, strftime('%H:%M %d/%m/%Y','now', 'localtime'))")
	if err != nil {
		fmt.Println("error adding comment to DB")
		return
	}
	result, _ := stmt.Exec(user, postID, commentContent)
	rowsAff, _ := result.RowsAffected()
	LastIns, _ := result.LastInsertId()
	fmt.Println("rows affected: ", rowsAff)
	fmt.Println("last inserted: ", LastIns)
}
