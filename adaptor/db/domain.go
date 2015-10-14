package db

import (
	"time"

	"gopkg.in/mgo.v2"
)

type Domain struct {
	userID     mgo.ObjectId
	domain     string
	subdomain  string
	created_at time.Time
	updated_at time.Time
}

func GetDomainCollection(DB *mgo.Database) *mgo.Collection {
	return DB.C("domain")
}
