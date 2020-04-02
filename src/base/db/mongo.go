package db

import (
	"gopkg.in/mgo.v2"
)

func NewMongo(mongo_server string) *mgo.Session {

	session, err := mgo.Dial(mongo_server)
	if err != nil {
		return nil
	}

	// Optional. Switch the session to a monotonic behavior.
	//session.SetMode(mgo.Monotonic, true)

	return session
}

func DelMongo(session *mgo.Session) {
	if session == nil {
		return
	}

	session.Close()
}
