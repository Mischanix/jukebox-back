package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Nick     string  `json:"nick"`
	Quarters float64 `json:"quarters"`
	SkipCost float64 `json:"skipCost"`
	Fake     bool    `json:"fake"`
}

type DbUser struct {
	Nick              string
	Quarters          float64
	SkipCost          float64
	Fake              bool
	Base64Key         string
	LastSessionId     string
	LastSessionSecret string
}

func UserWithSid(sid string) *DbUser {
	result := &DbUser{}
	err := dbColl.Find(bson.M{"lastsessionid": sid}).One(result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil
		} else {
			panic(err)
		}
	}
	return result
}

func UpdateUserSecret(oldSid, newSid, newSecret string) error {
	return dbColl.Update(
		bson.M{"lastsessionid": oldSid},
		bson.M{"$set": bson.M{
			"lastsessionid":     newSid,
			"lastsessionsecret": newSecret,
		}})
}

func UpdateUser(sid string, user *User) error {
	return dbColl.Update(
		bson.M{"lastsessionid": sid}, *user)
}

func (c *Client) CreateFakeUser() {
	var dbUser DbUser
	dbUser.Nick = FakeName()
	dbUser.SkipCost = 0.1
	dbUser.Fake = true
	dbUser.LastSessionId = c.session.OldId
	dbUser.LastSessionSecret = c.session.Secret
	err := dbColl.Insert(dbUser)
	if err != nil {
		panic(err)
	}

	c.UpdateUserFromDb(dbUser)
}

func (c *Client) UpdateUserFromDb(dbUser DbUser) {
	c.user = &User{}
	c.user.Fake = dbUser.Fake
	c.user.Nick = dbUser.Nick
	c.user.Quarters = dbUser.Quarters
	c.user.SkipCost = dbUser.SkipCost
}

func (c *Client) SendUserInfo() {
	c.sendQueue <- userMessage(c.user.Nick, c.user.Quarters, c.user.SkipCost)
}
