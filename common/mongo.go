package common

import (
	"github.com/hjkelly/zbbapi/config"
	mgo "gopkg.in/mgo.v2"
)

// GetMongoSession connects to our single Mongo server. Eventually this will be split up by controller/service.
func GetMongoSession() *mgo.Session {
	url := config.GetConfig().MongoURL
	session, err := mgo.Dial(url)
	if err != nil {
		panic("Couldn't connect to the Mongo server at " + url)
	}
	return session
}
