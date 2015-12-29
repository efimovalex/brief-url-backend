package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/efimovalex/brief_url/adaptor/db"
	"github.com/efimovalex/brief_url/client"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// UserEndpoints exists as an example
type UserEndpoints struct {
	DB     *db.Adaptor
	config *Config
	Logger *log.Logger
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ue *UserEndpoints) Post(w http.ResponseWriter, r *http.Request) {
	user := db.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errs := []client.Error{{Message: err.Error()}}
		handleErrs(w, http.StatusInternalServerError, errs)

		return
	}
	err, field := user.Validate(ue.DB.User)
	if err != nil {
		errs := []client.Error{{Message: err.Error(), Field: field}}
		handleErrs(w, http.StatusBadRequest, errs)

		return
	}

	err = ue.DB.User.Add(&user)
	if err != nil {
		errs := []client.Error{{Message: err.Error()}}
		handleErrs(w, http.StatusBadRequest, errs)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ue *UserEndpoints) Get(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]
	mongoID := bson.ObjectIdHex(userID)
	user, err := ue.DB.User.Get(mongoID)
	if err != nil {
		errs := []client.Error{{Message: err.Error()}}
		handleErrs(w, http.StatusInternalServerError, errs)

		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		ue.Logger.Println("problem encoding url")
		errs := []client.Error{{Message: err.Error()}}
		handleErrs(w, http.StatusInternalServerError, errs)

		return
	}

	return
}

func (ue *UserEndpoints) Authenticate(w http.ResponseWriter, r *http.Request) {
	var login LoginCredentials

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		ue.Logger.Println("problem decoding body")
		errs := []client.Error{{Message: err.Error()}}
		handleErrs(w, http.StatusInternalServerError, errs)

		return
	}

	user, err := ue.DB.User.GetByEmail(login.Email)
	if err != nil {
		ue.Logger.Println("user not found")
		errs := []client.Error{{Message: "user not found"}}
		handleErrs(w, http.StatusInternalServerError, errs)

		return
	}

	valid, err := user.CheckPassword(login.Password)
	if !valid {
		ue.Logger.Println("invalid credentials")
		handleErrs(w, http.StatusUnauthorized, []client.Error{{Message: "authentication failed"}})

		return
	}

	apiKey := db.APIKey{}
	apiKey.UserID = user.ID
	apiKey.ExpiresAt = time.Now().Add(time.Hour * 24)
	apiKey.CreatedAt = time.Now()
	apiKey.Name = fmt.Sprintf("login_token_%d", apiKey.CreatedAt)
	apiKey.Hidden = true
	apiKey.Scopes = "api"

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["name"] = apiKey.Name
	token.Claims["hidden"] = apiKey.Hidden
	token.Claims["scopes"] = apiKey.Scopes
	token.Claims["exp"] = apiKey.ExpiresAt.Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(ue.config.JWTSigningKey))
	if err != nil {
		ue.Logger.Println("problem signing token")
		errs := []client.Error{{Message: err.Error()}}
		handleErrs(w, http.StatusInternalServerError, errs)

		return
	}

	apiKey.Token = tokenString

	ue.DB.ApiToken.Add(&apiKey)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(TokenResponse{Token: tokenString}); err != nil {
		ue.Logger.Println("problem encoding token")
		errs := []client.Error{{Message: err.Error()}}
		handleErrs(w, http.StatusInternalServerError, errs)

		return
	}

	return
}
