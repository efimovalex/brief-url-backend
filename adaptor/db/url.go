package db

import (
	"encoding/base64"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/pborman/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type URLCollection struct {
	DB *mgo.Collection
}

var DefaultTTL = 2 * 7 * 24 * time.Hour

type URL struct {
	ID        bson.ObjectId `bson:"_id"`
	Route     string        `bson:"route"`
	Redirect  string        `bson:"redirect"`
	createdAt time.Time     `bson:"created_at"`
	updatedAt time.Time     `bson:"updated_at"`
	ExpiresAt time.Time     `bson:"expires_at"`
	Active    bool          `bson:"active"`
}

func GetURLCollection(DB *mgo.Database) *URLCollection {
	return &URLCollection{DB: DB.C("url")}
}

func (uc *URLCollection) GetAll() ([]URL, error) {
	var URLs []URL

	err := uc.DB.Find(nil).All(&URLs)

	return URLs, err
}

func (uc *URLCollection) Get(URLID bson.ObjectId) (URL, error) {
	var url URL

	err := uc.DB.Find(bson.M{"_id": URLID}).One(&url)

	return url, err
}

func (uc *URLCollection) Add(url *URL) error {
	if !url.ID.Valid() {
		url.ID = bson.NewObjectId()
	}

	byteData := []byte(strings.Replace(uuid.New(), "-", "", -1))

	url.Route = base64.StdEncoding.EncodeToString(byteData)
	url.Active = true
	url.createdAt = time.Now()
	url.updatedAt = time.Now()

	url.ExpiresAt = time.Now().Add(DefaultTTL)
	err := url.Validate([]string{})
	if err != nil {
		return err
	}

	err = uc.DB.Insert(url)

	return err
}

func (uc *URLCollection) Delete(URLID string) error {
	return uc.DB.RemoveId(bson.ObjectIdHex(URLID))
}

func (u *URL) Validate(forbiddenURLs []string) error {
	if u.Route == "" || u.Route == "/" {
		return errors.New("invalid route")
	}
	_, parseErr := url.Parse(u.Redirect)
	if u.Redirect == "" || parseErr != nil {
		return errors.New("invalid redirect url")
	}

	for _, forbiddenURL := range forbiddenURLs {
		if u.Redirect == forbiddenURL {
			return errors.New("forbidden url")
		}
	}

	return nil
}
