package plans

import (
	"github.com/hjkelly/zbbapi/common"
	mgo "gopkg.in/mgo.v2"
)

type datastore struct {
	session *mgo.Session
}

func newDatastore() *datastore {
	return &datastore{common.GetMongoSession()}
}

func (ds datastore) C() *mgo.Collection {
	return ds.session.DB("plan").C("plans")
}
