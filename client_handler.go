package main

import (
	"bytes"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"io"
	"log"
)

type empty struct{}

type messageHandler func(c *Client, frame []byte)

var handlers = make(map[string]messageHandler)

func decodeFrame(frame []byte, v interface{}) {
	err := json.NewDecoder(bytes.NewReader(frame)).Decode(v)
	if err != nil {
		panic(err)
	}
}

func ClientHandler(ws *websocket.Conn) {
	done := make(chan empty, 1)
	defer func() {
		done <- empty{}
		err := recover()
		if err != io.EOF {
			websocket.JSON.Send(ws, errorMsg)
			ws.Close()
			panic(err)
		}
	}()

	client := &Client{}
	client.ws = ws
	client.sendQueue = make(chan interface{}, 1)
	client.session = &Session{}
	client.session.Id = getSessionId()
	clients[client.session.Id] = client
	client.sendQueue <- readyMsg

	go func() {
		for {
			select {
			case <-done:
				log.Println("write thread finishing")
				return
			case data := <-client.sendQueue:
				log.Println("sending packet meow")
				websocket.JSON.Send(ws, data)
			}
		}
	}()
	for {
		var frame []byte
		websocket.Message.Receive(ws, &frame)
		var msg BasicMessage
		decodeFrame(frame, &msg)
		log.Println(msg)
		if handler, ok := handlers[msg.Type]; ok {
			handler(client, frame)
		}
	}
}
