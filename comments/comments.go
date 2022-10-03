package comments

import (
	"database/sql"
	"fmt"
)

type Comments struct {
	CommentID      int    `json:"commentid"`
	CommentContent string `json:"commentcontent"`
	User           string `json:"user"`
	Date           string `json:"commenttime"`
	Tipo           string `json:"tipo"`
	PostID         string `json:"postid"`
}

type CommentsFromPosts struct {
	ClickedPostID string `json:"clickedPostID"`
	Tipo          string `json:"tipo"`
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



func DisplayAllComments(db *sql.DB, postID int) []Comments {
	rows, err := db.Query(`SELECT commentID, commentText, creationDate, username FROM comments 
	WHERE postID = ?;`, postID)
	if err != nil {
		fmt.Println("ERROR getting comments from posts", err)
	}
	comment := []Comments{}
	defer rows.Close()
	for rows.Next() {
		var c Comments
		err2 := rows.Scan(&c.CommentID, &c.CommentContent, &c.Date, &c.User)
		comment = append(comment, c)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return comment
}

// INNER JOIN posts ON posts.postID = comments.postID
// INNER JOIN users ON users.username = comments.username
