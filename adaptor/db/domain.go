package db

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

type Domain struct {
	ID         bson.ObjectId `bson:"_id"`
	userID     bson.ObjectId `bson:"user_id"`
	domain     string        `bson:"domain"`
	subdomain  string        `bson:"subdomain"`
	created_at time.Time     `bson:"created_at"`
	updated_at time.Time     `bson:"updated_at"`
}

func GetDomainCollection(DB *mgo.Database) *mgo.Collection {
	return DB.C("domain")
}
