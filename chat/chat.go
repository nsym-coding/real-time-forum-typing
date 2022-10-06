package chat

import (
	"database/sql"
	"fmt"
)

type Chat struct {
	ChatSender    string `json:"chatsender"`
	ChatRecipient string `json:"chatrecipient"`
	ChatMessage   string `json:"message"`
	Date          string `json:"chatDate"`
	Tipo          string `json:"tipo"`
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

		

//checking if a prior chat exists between the two users
func ChatHistoryValidation(db *sql.DB, user1 string, user2 string)bool{
	rows, err := db.Query("SELECT user1, user2 FROM chats")
	if err != nil {
		fmt.Println("Error from ChatHistoryV fn()", err)
	}

	var userone string
	var userone string
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userone, &usertwo)
		if err!= nil {
			log.Fatal("With ChatHisVal fn()", err)
		}
	}


}

//creates a chat entryy between two users on successful validation 
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
