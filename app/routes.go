package app

import (
	"fmt"
	"net/http"

	"github.com/efimovalex/brief_url/adaptor/db"
	"github.com/efimovalex/brief_url/client"
	"github.com/gorilla/mux"
)

type dependencies struct {
	dbAdaptor *db.Adaptor
}

// Routes returns an http.Handler with the available RESTful routes for the service
func Routes(deps dependencies) *mux.Router {
	router := mux.NewRouter()
	var resource string

	// place all routes here to make it easier to find
	resources := map[string]string{
		"sampleEndpoint": "/path/to/toggle/controlled/endpoint",
	}

	sampleEndpoint := &SampleEndpoint{toggler: deps.dbAdaptor}
	resource = resources["sampleEndpoint"]
	router.HandleFunc(resource, sampleEndpoint.Get).Name(resource).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(NotFound)

	return router
}

// NotFound JSON 404 page
func NotFound(w http.ResponseWriter, r *http.Request) {
	ln.Info(fmt.Sprintf("404 - %s %s", r.Method, r.RequestURI), nil)
	w.Header().Set("content-type", "application/json")
	handleErr(w, http.StatusNotFound, client.ErrorMessageResourceNotFound, "")
}
