package main

import (
	"log"
	"time"
)

func init() {
	handlers["session"] = sessionHandler
}

func sessionHandler(c *Client, frame []byte) {
	before := time.Now()
	var msg SessionMessage
	decodeFrame(frame, &msg)
	log.Println(msg)
	dbUser := UserWithSid(msg.Session)
	newSid := c.session.Id
	c.session.Secret = makeSessionSecret()
	if dbUser != nil && msg.Secret == dbUser.LastSessionSecret {
		UpdateUserSecret(msg.Session, newSid, c.session.Secret)
	} else {
		CreateFakeUser(newSid, c.session.Secret)
	}

	var response SessionResponseMessage
	response.Type = "session.response"
	response.Accepted = false
	response.Session = newSid
	response.Secret = c.session.Secret
	c.sendQueue <- response

	duration := time.Now().Sub(before)
	log.Println("session.response took", duration)
}

// query db for user with sid as provided, if exists and secret matches then
// session is accepted, and the user is updated/changed appopriately
// otherwise a new fake user is created

// on login/register, query db for user with nick as provided, if exists and
// password matches then change the user as appropriate; if exists and password
// does not match, do nothing and fail; if not exists, change the user from fake
// to real and give it the new nick/pass hash
