package common

import (
	"github.com/hjkelly/zbbapi/config"
	mgo "gopkg.in/mgo.v2"
)

func GetMongoSession() *mgo.Session {
	url := config.GetConfig().MONGO_URL
	session, err := mgo.Dial(url)
	if err != nil {
		panic("Couldn't connect to the Mongo server at " + url)
	}
	return session
}
