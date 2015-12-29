package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/agnivade/easy-scrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID         bson.ObjectId   `bson:"_id" json:"id"`
	Email      string          `bson:"email" json:"email"`
	Password   string          `bson:"password" json:"password"`
	Active     bool            `bson:"active"`
	Domains    []bson.ObjectId `bson:"domains"`
	URLs       []bson.ObjectId `bson:"urls"`
	APIKeys    []bson.ObjectId `bson:"apikeys"`
	created_at time.Time       `bson:"created_at"`
	updated_at time.Time       `bson:"updated_at"`
}

type UserCollection struct {
	DB *mgo.Collection
}

func GetUserCollection(DB *mgo.Database) *UserCollection {
	return &UserCollection{DB: DB.C("user")}
}

func (uc *UserCollection) GetByEmail(email string) (User, error) {
	var user User

	err := uc.DB.Find(bson.M{"email": email}).One(&user)

	return user, err
}

func (uc *UserCollection) Add(user *User) error {
	if !user.ID.Valid() {
		user.ID = bson.NewObjectId()
	}

	err := user.EncryptPassword()
	if err != nil {
		return err
	}

	user.created_at = time.Now()
	user.updated_at = user.created_at
	user.Active = true

	err = uc.DB.Insert(user)

	return err
}

func (uc *UserCollection) Get(userID bson.ObjectId) (User, error) {
	var user User

	err := uc.DB.Find(bson.M{"_id": userID}).One(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *User) EncryptPassword() error {
	key, err := scrypt.DerivePassphrase(u.Password, 32)
	if err != nil {
		return fmt.Errorf("error encoding password: ", err.Error())
	}

	u.Password = fmt.Sprintf("%s", key)

	return nil
}

func (u *User) CheckPassword(passphrase string) (bool, error) {
	result, err := scrypt.VerifyPassphrase(passphrase, []byte(u.Password))
	if err != nil {
		return false, fmt.Errorf("error encoding password: %s", err.Error())
	}

	if !result {
		return false, fmt.Errorf("passphrase did not match")
	} else {
		return true, nil
	}
}

// Validate returns validation errors and the field in question
func (u *User) Validate(uc *UserCollection) (error, string) {
	if u.Email == "" {
		return errors.New("email is missing"), "email"
	}

	existingEmailUser := User{}
	err := uc.DB.Find(bson.M{"Email": u.Email}).One(&existingEmailUser)
	if err != nil && err != mgo.ErrNotFound {
		return fmt.Errorf("error validating user: %s", err.Error()), "email"
	} else if existingEmailUser.Email != "" {
		return fmt.Errorf("email is already in use: %s", existingEmailUser.Email), "email"
	}

	if u.Password == "" {
		return errors.New("password is missing"), "password"
	}

	if len(u.Password) < 6 {
		return errors.New("password is too short"), "password"
	}

	return nil, ""
}
