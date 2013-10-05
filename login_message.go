package main

type LoginMessage struct {
	BasicMessage
	Nick string `json:"nick"`
	Pass string `json:"pass"`
}

type LoginResponseMessage struct {
	BasicMessage
	Status string `json:"status"` // ok, nok
	Reason string `json:"reason"`
}
