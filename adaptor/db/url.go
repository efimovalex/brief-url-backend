package db

import (
	"time"

	"gopkg.in/mgo.v2"
)

type URL struct {
	ID         bson.ObjectId
	Route      string
	Redirect   string
	userID     bson.ObjectId
	domainID   bson.ObjectId
	Stats      []Stat.URLID
	created_at time.Time
	updated_at time.Time
	expires_At time.Time
	Active     bool
}

func GetURLCollection(DB *mgo.Database) *mgo.Collection {
	return DB.C("url")
}
