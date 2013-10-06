package main

func sessionMessage(frame hash) (session, secret string) {
	return frame.stringOrEmpty("session"), frame.stringOrEmpty("secret")
}

func sessionResponseMessage(session, secret string, accepted bool) hash {
	return hash{
		"type":     "session.response",
		"session":  session,
		"secret":   secret,
		"accepted": accepted,
	}
}
