package chat

import (
	"database/sql"
	"fmt"
	"log"

	notification "real-time-forum/notifications"
)

type Chat struct {
	ChatSender       string                    `json:"chatsender"`
	ChatRecipient    string                    `json:"chatrecipient"`
	ChatMessage      string                    `json:"message"`
	MessageID        int                       `json:"messageID"`
	Date             string                    `json:"chatDate"`
	LastNotification notification.Notification `json:"livenotification"`
	Tipo             string                    `json:"tipo"`
}

type ChatHistory struct {
	ChatHist []Chat `json:"chathistory"`
	Tipo     string `json:"tipo"`
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

func StoreMessages(db *sql.DB, chatID int, message, chatSender, chatRecipient string) {
	stmt, err := db.Prepare(`INSERT INTO messages (chatID, message, sender, recipient, creationDate ) VALUES (?,?,?,?,strftime('%H:%M %d/%m/%Y','now', 'localtime') )`)
	if err != nil {
		fmt.Println("error adding message to DB", err)
		return
	}
	result, _ := stmt.Exec(chatID, message, chatSender, chatRecipient)
	rowsAff, _ := result.RowsAffected()
	LastIns, _ := result.LastInsertId()
	fmt.Println("chat rows affected: ", rowsAff)
	fmt.Println("chat last inserted: ", LastIns)
}

// Function that returns chats based on a chat id.
func GetAllMessageHistoryFromChat(db *sql.DB, chatID int) ChatHistory {
	rows, err := db.Query(`SELECT message, sender, recipient, creationDate FROM messages WHERE chatID = ?;`, chatID)
	if err != nil {
		fmt.Println(err)
	}
	messagedata := ChatHistory{}
	messagedata.Tipo = "messagehistoryfromgo"

	defer rows.Close()
	for rows.Next() {
		var m Chat
		// fmt.Println(&p.PostID)
		err2 := rows.Scan(&m.ChatMessage, &m.ChatSender, &m.ChatRecipient, &m.Date)
		// m.Tipo = "messagehistoryfromgo"
		messagedata.ChatHist = append(messagedata.ChatHist, m)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return messagedata
}

func GetChat(db *sql.DB, user string) []int {
	// check if chat already exists
	rows, err := db.Query("SELECT chatID FROM chats WHERE user1 = ? OR user2 = ?", user, user)
	if err != nil {
		fmt.Println(err)
	}

	var chatsID []int
	defer rows.Close()
	for rows.Next() {
		var uIDs int
		err2 := rows.Scan(&uIDs)
		chatsID = append(chatsID, uIDs)
		if err2 != nil {
			fmt.Println("chatID doesn't exist-----------")
		}
	}
	return chatsID
}

func GetLatestChat(db *sql.DB, chatID []int) []Chat {
	var latestChat []Chat
	for i := 0; i < len(chatID); i++ {

		// check latest chat
		userStmt := "SELECT sender, recipient, max(messageID) FROM messages WHERE chatID = ?;"
		rowU := db.QueryRow(userStmt, chatID[i])

		var c Chat
		error := rowU.Scan(&c.ChatSender, &c.ChatRecipient, &c.MessageID)
		if error != sql.ErrNoRows {
			fmt.Println("username already exists, err:", error)
			latestChat = append(latestChat, c)
		}
	}
	return latestChat
}
