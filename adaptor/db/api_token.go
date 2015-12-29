package db

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type APIKey struct {
	ID        bson.ObjectId `bson:"_id"`
	UserID    bson.ObjectId `bson:"user_id"`
	Scopes    string        `bson:"scopes"`
	Token     string        `bson:"token"`
	Name      string        `bson:"name"`
	Hidden    bool          `bson:"hidden"`
	CreatedAt time.Time     `bson:"created_at"`
	updatedAt time.Time     `bson:"updated_at"`
	RevokedAt time.Time     `bson:"revoked_at"`
	ExpiresAt time.Time     `bson:"expires_in"`
}

type APIKeyCollection struct {
	DB *mgo.Collection
}

func GetAPITokenCollection(DB *mgo.Database) *APIKeyCollection {
	return &APIKeyCollection{DB: DB.C("api_token")}
}

func (ac *APIKeyCollection) Add(apiKey *APIKey) error {
	if !apiKey.ID.Valid() {
		apiKey.ID = bson.NewObjectId()
	}

	apiKey.updatedAt = apiKey.CreatedAt

	err := ac.DB.Insert(apiKey)

	return err
}
