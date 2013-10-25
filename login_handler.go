package main

import (
	"log"
	"strings"
	"time"
)

func init() {
	handlers["login"] = loginHandler
}

func loginHandler(c *Client, frame hash) {
	before := time.Now()
	nick, pass := loginMessage(frame)
	if nick == "" || pass == "" {
		c.sendQueue <- loginResponseMessage("nok", "need nick and pass")
		return
	}
	if len(nick) > 50 {
		c.sendQueue <- loginResponseMessage("nok", "nick too long")
		return
	}
	if !c.user.Fake {
		c.sendQueue <- loginResponseMessage("nok", "already logged in")
		return
	}
	dbUser := UserWithNick(nick)
	if dbUser == nil { // register
		dbUser = UserWithSid(c.session.Id)
		if dbUser == nil {
			// unlikely
			c.sendQueue <- loginResponseMessage("nok", "victim not found")
			return
		}
		dbUser.Nick = nick
		dbUser.NickLower = strings.ToLower(nick)
		dbUser.Fake = false
		dbUser.NewKey(pass)
		c.UpdateUser(*dbUser)
		c.sendQueue <- loginResponseMessage("ok", "")
	} else { // login
		if dbUser.TestKey(pass) {
			c.sendQueue <- loginResponseMessage("ok", "")
			c.DestroyFakeUser()
			UpdateUserSecret(dbUser.LastSessionId, c.session.Id, c.session.Secret)
		} else {
			c.sendQueue <- loginResponseMessage("nok", "bad credentials")
			return
		}
	}
	c.UpdateUserFromDb(*dbUser)
	c.KillClones()
	c.SendUserInfo()
	log.Println("login took", time.Now().Sub(before))
}
