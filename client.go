package main

import (
	"code.google.com/p/go.net/websocket"
)

var clients = make(map[string]*Client)

type Client struct {
	ws        *websocket.Conn
	session   *Session
	user      *User
	active    bool
	sendQueue chan interface{}
}
