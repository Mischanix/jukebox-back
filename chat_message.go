package main

func chatMessage(frame hash) (message string) {
	return frame.stringOrEmpty("message")
}

func chatReceiveMessage(nick, message string) hash {
	return hash{
		"type":    "chat.receive",
		"nick":    nick,
		"message": message,
	}
}
