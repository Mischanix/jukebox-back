package main

type BasicMessage struct {
	Type string `json:"type"`
}

var readyMsg = BasicMessage{"ready"}
var errorMsg = BasicMessage{"error"}
