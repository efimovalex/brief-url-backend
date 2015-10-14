package db

import (
	"time"

	"gopkg.in/mgo.v2"
)

type APIKey struct {
	userID     mgo.ObjectId
	scopes     string
	token      string
	name       string
	created_at time.Time
	revoked_at time.Time
	expires_in time.Duration
}

func GetAPIKeyCollection(DB *mgo.Database) *mgo.Collection {
	return DB.C("api_token")
}
