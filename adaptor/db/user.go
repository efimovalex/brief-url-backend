package db

import (
	"time"

	"gopkg.in/mgo.v2"
)

type User struct {
	ID         mgo.ObjectId
	Email      string
	password   string
	Active     bool
	created_at time.Time
	updated_at time.Time
}

func GetUserCollection(DB *mgo.Database) *mgo.Collection {
	return DB.C("user")
}
