package main

func init() {
	handlers["session"] = sessionHandler
	handlers["session.updated"] = sessionUpdatedHandler
}

func sessionHandler(c *Client, frame hash) {
	oldSid, oldSecret := sessionMessage(frame)
	dbUser := UserWithSid(oldSid)
	newSid := c.session.Id
	if oldSid == "" {
		// Fake it if the client has no sid -- cleaner this way
		c.session.OldId = getSessionId()
	} else {
		c.session.OldId = oldSid
	}
	c.session.Secret = makeSessionSecret()
	var accepted bool
	if dbUser != nil && oldSecret == dbUser.LastSessionSecret {
		accepted = true
		c.UpdateUserFromDb(*dbUser)
	} else {
		accepted = false
		c.CreateFakeUser()
	}

	response := sessionResponseMessage(newSid, c.session.Secret, accepted)
	c.sendQueue <- response
}

func sessionUpdatedHandler(c *Client, frame hash) {
	if c.session.OldId != "" {
		UpdateUserSecret(c.session.OldId, c.session.Id, c.session.Secret)
		c.active = true
		c.SendUserInfo()
	}
}

// query db for user with sid as provided, if exists and secret matches then
// session is accepted, and the user is updated/changed appopriately
// otherwise a new fake user is created

// on login/register, query db for user with nick as provided, if exists and
// password matches then change the user as appropriate; if exists and password
// does not match, do nothing and fail; if not exists, change the user from fake
// to real and give it the new nick/pass hash
