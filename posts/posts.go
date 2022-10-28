package posts

import (
	"database/sql"
	"fmt"
	"real-time-forum/comments"
)

type Posts struct {
	PostID      int                 `json:"postid"`
	PostTitle   string              `json:"title"`
	PostContent string              `json:"postcontent"`
	Date        string              `json:"posttime"`
	Tipo        string              `json:"tipo"`
	Categories  string              `json:"categories"`
	Username    string              `json:"username"`
	Comments    []comments.Comments `json:"comments"`
}

func StorePosts(db *sql.DB, username string, title string, content string, categories string) {
	stmt, err := db.Prepare("INSERT INTO posts (username, postTitle, postContent, categories, creationDate) VALUES (?, ?, ?, ?, strftime('%H:%M %d/%m/%Y','now','localtime'))")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		return
	}
	// defer stmt.Close()
	result, _ := stmt.Exec(username, title, content, categories)
	// checking if the result has been added and the last inserted row
	rowsAff, _ := result.RowsAffected()
	LastIns, _ := result.LastInsertId()
	fmt.Println("rows affected:", rowsAff)
	fmt.Println("last inserted:", LastIns)
}

func SendPostsInDatabase(db *sql.DB) []Posts {

	rows, err := db.Query(`SELECT postID, postTitle, postContent, username, categories, creationDate FROM posts ;`)
	if err != nil {
		fmt.Println(err)
	}
	postdata := []Posts{}
	defer rows.Close()
	for rows.Next() {
		var p Posts
		// fmt.Println(&p.PostID)
		err2 := rows.Scan(&p.PostID, &p.PostTitle, &p.PostContent, &p.Username, &p.Categories, &p.Date)
		p.Comments = comments.DisplayAllComments(db, p.PostID)
		postdata = append(postdata, p)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return postdata

}

func SendLastPostInDatabase(db *sql.DB) Posts {

	rows, err := db.Query(`SELECT postID, postTitle, postContent, username, categories,creationDate FROM posts;`)
	if err != nil {
		fmt.Println(err)
	}
	postdata := []Posts{}
	defer rows.Close()
	for rows.Next() {
		var p Posts
		// fmt.Println(&p.PostID)
		err2 := rows.Scan(&p.PostID, &p.PostTitle, &p.PostContent, &p.Username, &p.Categories, &p.Date)
		p.Tipo = "post"
		postdata = append(postdata, p)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return postdata[len(postdata)-1]

}

//SELECT * FROM Table ORDER BY ID DESC LIMIT 1

func GetCommentData(db *sql.DB, postID int) Posts {
	rows, err := db.Query(`SELECT commentID, commentText, comments.creationDate as cmntDate, posts.username
    FROM comments
    INNER JOIN posts ON posts.postID = comments.postID
    WHERE comments.postID = ?;`, postID)
	if err != nil {
		fmt.Println(err)
	}
	post := Posts{}
	defer rows.Close()
	for rows.Next() {
		// var c Posts
		var comments comments.Comments
		err2 := rows.Scan(&comments.CommentID, &comments.CommentContent, &comments.Date, &comments.User)

		post.Comments = append(post.Comments, comments)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	fmt.Println("comments for post", post.Comments)
	return post
}
