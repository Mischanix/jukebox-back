package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
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
		bson.M{"lastsessionid": newSid, "lastsessionsecret": newSecret})
}

func UpdateUser(sid string, user *User) error {
	return dbColl.Update(
		bson.M{"lastsessionid": sid}, *user)
}

func CreateFakeUser(sid, secret string) {
	log.Println("I be makin a fake user now mmhm (nah jk)")
}
