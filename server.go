package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", websocket.Handler(ClientHandler))
	err := http.ListenAndServe(":3343", nil)
	if err != nil {
		log.Fatal(err)
	}
}
