package app

import (
	"log"
	"net/http"

	"github.com/efimovalex/brief_url/adaptor/db"
	"github.com/efimovalex/brief_url/client"
	"github.com/gorilla/mux"
)

type dependencies struct {
	dbAdaptor *db.Adaptor
	config    *Config
	logger    *log.Logger
}

// Routes returns an http.Handler with the available RESTful routes for the service
func Routes(deps dependencies) *mux.Router {
	router := mux.NewRouter()

	// place all routes here to make it easier to find
	resources := map[string]string{
		"urlCollectionEndpoint": "/v1/url",
		// [GET]  body: n/a
		// [POST] body: {"redirect_url":"http://www.google.com"}
		"urlEndpoint": "/v1/url/{url_id}",
		// GET    body: n/a
		// DELETE body: n/a
		// PATCH  body: {"redirect_url":"http://www.google.com"}
		"userCollenctionEndpoint": "/v1/user",
		// [POST] body: {"email":"blablabla@google.com","password":"passwd!1"}
		"authenticateEndpoint": "/v1/user/authenticate",
		// [POST] formdata: email=...&password=xxx
		"userEndpoint": "/v1/user/{user_id}",
		// [GET] body: n/a
	}

	urlEndpoint := &URLEndpoints{DB: deps.dbAdaptor, Logger: deps.logger}

	URLCollectionResource := resources["urlCollectionEndpoint"]
	URLResource := resources["urlEndpoint"]

	userEndpoint := &UserEndpoints{DB: deps.dbAdaptor, config: deps.config, Logger: deps.logger}
	UserCollectionResource := resources["userCollenctionEndpoint"]
	UserResource := resources["userEndpoint"]

	AuthenticationResource := resources["authenticateEndpoint"]

	// URL Endpoints
	router.HandleFunc(URLCollectionResource, urlEndpoint.Get).Name(URLCollectionResource).Methods("GET")
	router.HandleFunc(URLCollectionResource, urlEndpoint.Post).Name(URLCollectionResource).Methods("POST")
	router.HandleFunc(URLResource, urlEndpoint.Get).Name(URLResource).Methods("GET")
	router.HandleFunc(URLResource, urlEndpoint.Delete).Name(URLResource).Methods("DELETE")
	router.HandleFunc(URLResource, urlEndpoint.Patch).Name(URLResource).Methods("PATCH")

	// User Endpoints
	router.HandleFunc(UserCollectionResource, userEndpoint.Post).Name(UserCollectionResource).Methods("POST")
	router.HandleFunc(UserResource, userEndpoint.Get).Name(UserResource).Methods("GET")

	// Authentication
	router.HandleFunc(AuthenticationResource, userEndpoint.Authenticate).Name(AuthenticationResource).Methods("POST")

	router.NotFoundHandler = http.HandlerFunc(NotFound)

	return router
}

// NotFound JSON 404 page
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 - %s %s", r.Method, r.RequestURI)
	w.Header().Set("content-type", "application/json")
	handleErr(w, http.StatusNotFound, client.ErrorMessageResourceNotFound, "")
}
