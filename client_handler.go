package main

import (
	"code.google.com/p/go.net/websocket"
	"io"
	"log"
)

type empty struct{}

type messageHandler func(c *Client, frame hash)

var handlers = make(map[string]messageHandler)

// ClientHandler handles a websocket connection for one client.  It initializes
// the client's Client object, dispatches incoming messages to their
// appropriate handlers, and writes outgoing messages as well.  Any error
// encountered while handling the client will cause the connection to be closed
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
				return
			case data := <-client.sendQueue:
				websocket.JSON.Send(ws, data)
			}
		}
	}()
	for {
		frame := make(hash)
		err := websocket.JSON.Receive(ws, &frame)
		if err != nil {
			// Hack: Handle IE10 being a special snowflake and sending Pong frames.
			// If IE10 doesn't receive a message within its deadline (default: 10s)
			// after sending the Pong frame, it will close the connection.  ENI is
			// only used in the case of the server receiving a Pong frame.
			if err == websocket.ErrNotImplemented {
				client.sendQueue <- goPongYourselfMsg
			} else {
				panic(err)
			}
		} else {
			msgType := frame["type"].(string)
			log.Println(msgType)
			if handler, ok := handlers[msgType]; ok {
				handler(client, frame)
			}
		}
	}
}
