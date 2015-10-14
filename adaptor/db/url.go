package db

import (
	"time"

	"gopkg.in/mgo.v2"
)

type URL struct {
	Route      string
	Redirect   string
	userID     mgo.ObjectId
	domainID   mgo.ObjectId
	created_at time.Time
	updated_at time.Time
	expires_At time.Time
	Active     bool
}

func GetURLCollection(DB *mgo.Database) *mgo.Collection {
	return DB.C("url")
}
