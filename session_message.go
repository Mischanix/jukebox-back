package main

type SessionMessage struct {
	BasicMessage
	Session string `json:"session"`
	Secret  string `json:"secret"`
}

type SessionResponseMessage struct {
	SessionMessage
	Accepted bool `json:"accepted"`
}
