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
	var resource string

	// place all routes here to make it easier to find
	resources := map[string]string{
		"domainsEndpoint": "/domains/user/{user_id}",
		"domainEndpoint":  "/domains/user/{user_id}/domain/{domain_id}",
	}

	domainEndpoint := &DomainEndpoint{}
	domainsResource = resources["domainsEndpoint"]
	domainResource = resources["domainsEndpoint"]
	router.HandleFunc(resource, domainEndpoint.Get).Name(resource).Methods("GET")
	router.HandleFunc(resource, domainsEndpoint.Get).Name(resource).Methods("GET")

	router.HandleFunc(resource, domainEndpoint.Post).Name(resource).Methods("POST")

	router.HandleFunc(resource, domainEndpoint.Delete).Name(resource).Methods("DELETE")
	router.HandleFunc(resource, domainsEndpoint.Get).Name(resource).Methods("DELETE")

	router.HandleFunc(resource, domainEndpoint.Patch).Name(resource).Methods("PATCH")

	router.NotFoundHandler = http.HandlerFunc(NotFound)

	return router
}

// NotFound JSON 404 page
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("404 - %s %s", r.Method, r.RequestURI)
	w.Header().Set("content-type", "application/json")
	handleErr(w, http.StatusNotFound, client.ErrorMessageResourceNotFound, "")
}
