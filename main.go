package main

import (
	"log"
	"net/http"
	"real-time-forum/db"
	"real-time-forum/socket"
)

func main() {
	db.CreateDB()
	// need to figure out how to serve these files on the homepage (/)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/login", socket.GetLoginData)
	http.HandleFunc("/ws", socket.WebSocketEndpoint)

	log.Println("Listening on port :8080.....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
