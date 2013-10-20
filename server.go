package main

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

func main() {
	go player()

	server := websocket.Server{}
	server.Handler = websocket.Handler(ClientHandler)
	config, err := websocket.NewConfig("ws://:3343", "http://:3333")
	if err != nil {
		log.Fatal(err)
	}
	server.Config = *config
	http.Handle("/", server)
	err = http.ListenAndServe(":3343", nil)
	if err != nil {
		log.Fatal(err)
	}
}
