package main

import (
	"log"
	"strings"
)

func init() {
	handlers["chat"] = chatHandler
}

func chatHandler(c *Client, frame hash) {
	message := strings.TrimSpace(chatMessage(frame))
	if len(message) == 0 {
		return
	}
	broadcast(chatReceiveMessage(c.user.Nick, message))
	log.Println("client count", len(clients))
}
