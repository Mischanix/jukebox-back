package main

import (
	"log"
)

func init() {
	handlers["chat"] = chatHandler
}

func chatHandler(c *Client, frame hash) {
	message := chatMessage(frame)
	receiveMessage := chatReceiveMessage(c.user.Nick, message)
	for _, client := range clients {
		if client != c && client.active {
			client.sendQueue <- receiveMessage
		}
	}
	log.Println("client count", len(clients))
}
