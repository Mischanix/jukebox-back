package main

func loginMessage(frame hash) (nick, pass string) {
	return frame.stringOrEmpty("nick"), frame.stringOrEmpty("pass")
}

func loginResponseMessage(status, reason string) hash {
	return hash{
		"type":   "login.response",
		"status": status,
		"reason": reason,
	}
}
