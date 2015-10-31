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
	logger    *log.Logger
}

// Routes returns an http.Handler with the available RESTful routes for the service
func Routes(deps dependencies) *mux.Router {
	router := mux.NewRouter()

	// place all routes here to make it easier to find
	resources := map[string]string{
		"urlsEndpoint": "/v1/url",
		"urlEndpoint":  "/v1/url/{url_id}",
	}

	urlEndpoint := &URLEndpoint{DB: deps.dbAdaptor, Logger: deps.logger}
	URLsResource := resources["urlsEndpoint"]
	URLResource := resources["urlEndpoint"]

	router.HandleFunc(URLsResource, urlEndpoint.Get).Name(URLsResource).Methods("GET")
	router.HandleFunc(URLResource, urlEndpoint.Get).Name(URLResource).Methods("GET")

	router.HandleFunc(URLsResource, urlEndpoint.Post).Name(URLsResource).Methods("POST")

	router.HandleFunc(URLResource, urlEndpoint.Delete).Name(URLResource).Methods("DELETE")

	router.HandleFunc(URLResource, urlEndpoint.Patch).Name(URLResource).Methods("PATCH")

	router.NotFoundHandler = http.HandlerFunc(NotFound)

	return router
}

// NotFound JSON 404 page
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 - %s %s", r.Method, r.RequestURI)
	w.Header().Set("content-type", "application/json")
	handleErr(w, http.StatusNotFound, client.ErrorMessageResourceNotFound, "")
}
