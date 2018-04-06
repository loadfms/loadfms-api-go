package utils

import (
	"log"

	"gopkg.in/mgo.v2"
)

func GetSession() *mgo.Session {
	var session *mgo.Session
	if session == nil {

		var err error
		session, err = mgo.Dial("localhost")

		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}

	return session
}
