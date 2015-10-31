package db

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DomainCollection struct {
	DB *mgo.Collection
}

type Domain struct {
	ID        string    `bson:"_id"`
	UserID    string    `bson:"user_id"`
	Domain    string    `bson:"domain"`
	Subdomain string    `bson:"subdomain"`
	createdAt time.Time `bson:"created_at"`
	updatedAt time.Time `bson:"updated_at"`
}

func GetDomainCollection(Database *mgo.Database) *DomainCollection {
	return &DomainCollection{DB: Database.C("domain")}
}

func (dc *DomainCollection) GetDomainsForUser(UserID string) ([]Domain, error) {
	domains := []Domain{}

	err := dc.DB.Find(bson.M{"user_id": UserID}).All(domains)

	return domains, err
}

func (dc *DomainCollection) GetDomain(UserID string, DomainID string) (Domain, error) {
	domain := Domain{}

	err := dc.DB.Find(bson.M{"user_id": UserID, "_id": DomainID}).One(domain)

	return domain, err
}

func (dc *DomainCollection) AddDomain(domain Domain) error {
	domain.createdAt = time.Now()
	domain.updatedAt = time.Now()

	err := dc.DB.Insert(&domain)

	return err
}

func (d *Domain) Check() error {
	if d.UserID == "" {
		return errors.New("user_id is missing")
	}
	if d.Domain == "" {
		return errors.New("domain is missing")
	}
	if d.Subdomain == "" {
		return errors.New("subdomain is missing")
	}

	return nil
}
