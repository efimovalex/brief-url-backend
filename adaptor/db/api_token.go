package db

import (
	"time"

	"gopkg.in/mgo.v2"
)

type APIKey struct {
	ID         bson.ObjectId `bson:"_id"`
	userID     bson.ObjectId `bson:"user_id"`
	scopes     string        `bson:"scopes"`
	token      string        `bson:"token"`
	name       string        `bson:"name"`
	created_at time.Time     `bson:"created_at"`
	revoked_at time.Time     `bson:"revoked_at"`
	expires_in time.Duration `bson:"expires_in"`
}

func GetAPITokenCollection(DB *mgo.Database) *mgo.Collection {
	return DB.C("api_token")
}
