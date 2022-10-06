package chat

import (
	"database/sql"
	"fmt"
	"log"
)

type Chat struct {
	ChatSender    string `json:"chatsender"`
	ChatRecipient string `json:"chatrecipient"`
	ChatMessage   string `json:"message"`
	Date          string `json:"chatDate"`
	Tipo          string `json:"tipo"`
}

type ChatExistsCheck struct {
	ChatID int
	Exists bool
}

// query
// rows, err := db.Query("SELECT * FROM userinfo")
// checkErr(err)
// var uid int
// var username string
// var department string
// var created time.Time

// for rows.Next() {
//     err = rows.Scan(&uid, &username, &department, &created)
//     checkErr(err)
//     fmt.Println(uid)
//     fmt.Println(username)
//     fmt.Println(department)
//     fmt.Println(created)
// }

// rows.Close() //good habit to close

// checking if a prior chat exists between the two users
func ChatHistoryValidation(db *sql.DB, user1 string, user2 string) ChatExistsCheck {
	rows, err := db.Query(`SELECT user1, user2, chatID FROM chats WHERE user1 = ? AND user2 =? OR user2 = ? AND user1 = ?;`, user1, user2, user1, user2)
	if err != nil {
		fmt.Println("Error from ChatHistoryV fn()", err)
	}

	// SELECT user1, user2 FROM chats WHERE user1 = "sancho" AND user2 = "royal" OR user1 = "royal" AND user2 = "sancho"

	var userone string
	var usertwo string
	var chatID int
	var chatExists ChatExistsCheck
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userone, &usertwo, &chatID)
		if err != sql.ErrNoRows {
			log.Println("With ChatHisVal fn()", err)
			log.Println("Users do have a chat")
			chatExists.ChatID = chatID
			chatExists.Exists = true
			return chatExists

		}
	}
	log.Println("Users don't have a chat")
	chatExists.ChatID = 0
	chatExists.Exists = false

	return chatExists
}

// creates a chat entryy between two users on successful validation
func StoreChat(db *sql.DB, chatsender string, chatrecipient string) {
	stmt, err := db.Prepare("INSERT INTO chats (user1, user2, creationDate) VALUES (?, ?, strftime('%H:%M %d/%m/%Y','now', 'localtime'))")
	if err != nil {
		fmt.Println("error adding chat to DB")
		return
	}
	result, _ := stmt.Exec(chatsender, chatrecipient)
	rowsAff, _ := result.RowsAffected()
	LastIns, _ := result.LastInsertId()
	fmt.Println("chat rows affected: ", rowsAff)
	fmt.Println("chat last inserted: ", LastIns)
}
