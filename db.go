package main

import (
	"labix.org/v2/mgo"
	"log"
)

var dbColl *mgo.Collection

const dbName = "jukebox"
const collName = "alpha"

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	dbColl = session.DB(dbName).C(collName)
}
