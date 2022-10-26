package notification

import (
	"database/sql"
	"fmt"
	"log"
)

type Notification struct {
	NotificatonSender     string `json:"notificationsender"`
	NotificationRecipient string `json:"notificationrecipient"`
	NotificationCount     int    `json:"notificationcount"`
	Tipo                  string `json:"tipo"`
	Type                  string `json:"clientspecificnotifications"`
}

func CheckNotification(db *sql.DB, sender, recipient string) bool {
	rows, err := db.Query(`SELECT sender, recipient FROM notifications WHERE sender = ? AND recipient =?;`, sender, recipient)
	if err != nil {
		fmt.Println("Error from CheckNotification fn()", err)
	}

	// SELECT user1, user2 FROM notifications WHERE sender = "sancho" AND recipient = "royal"

	var userone string
	var usertwo string
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userone, &usertwo)
		if err != sql.ErrNoRows {
			log.Println("With ChatHisVal fn()", err)
			log.Println("User does have a notification")
			return true
		}
	}
	log.Println("User doesn't have a notification")

	return false
}

func AddFirstNotificationForUser(db *sql.DB, sender, recipient string) {
	stmt, err := db.Prepare("INSERT INTO notifications (sender, recipient, count) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("error adding notification to DB")
		return
	}

	result, _ := stmt.Exec(sender, recipient, 1)
	rowsAff, _ := result.RowsAffected()
	LastIns, _ := result.LastInsertId()
	fmt.Println("chat rows affected: ", rowsAff)
	fmt.Println("chat last inserted: ", LastIns)
}

func IncrementNotifications(db *sql.DB, sender, recipient string) {
	if _, err := db.Exec("UPDATE notifications SET count = count + 1  WHERE sender = ? AND recipient = ?;", sender, recipient); err != nil {
		fmt.Println("error incrementing notification to DB")
		return
	}
}

func RemoveNotifications(db *sql.DB, sender, recipient string) bool {
	if _, err := db.Exec("UPDATE notifications SET count = 0 WHERE sender = ? AND recipient = ?;", sender, recipient); err != nil {
		fmt.Println("error removing notification to DB")
		return false
	}
	return true
}

func NotificationQuery(db *sql.DB, recipient string) []Notification {
	rows, err := db.Query(`SELECT sender, count FROM notifications WHERE recipient =?;`, recipient)
	if err != nil {
		fmt.Println("Error from CheckNotification fn()", err)
	}

	var notificationData []Notification

	defer rows.Close()
	for rows.Next() {
		var m Notification
		m.NotificationRecipient = recipient
		m.Tipo = "notification"
		err2 := rows.Scan(&m.NotificatonSender, &m.NotificationCount)
		notificationData = append(notificationData, m)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	return notificationData
}

func SingleNotification(db *sql.DB, sender, recipient string) Notification {
	stmt := `SELECT count FROM notifications WHERE sender=? AND recipient =?;`
	row := db.QueryRow(stmt, sender, recipient)

	var n Notification

	n.NotificatonSender = sender
	n.NotificationRecipient = recipient
	err := row.Scan(&n.NotificationCount)

	if err != nil {
		fmt.Println("error from singleNotification", err)
	}
	return n
}
