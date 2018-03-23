package categories

import (
	"github.com/hjkelly/zbbapi/common"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Delete(id string) error {
	ds := newDatastore()
	err := ds.C().Remove(bson.M{
		"_id": uuid.FromStringOrNil(id),
	})
	if err != nil {
		if err == mgo.ErrNotFound {
			return common.NotFoundErr
		} else {
			return err
		}
	}
	return nil
}
