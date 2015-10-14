package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/efimovalex/brief_url/client"
)

// handleErrs is designed to conform to sendgrid's api guidelines for returning errors
func handleErrs(w http.ResponseWriter, statusCode int, errs []client.Error) {
	appErr := &client.ErrorResult{}
	appErr.Errors = errs

	jsonResponse, err := json.Marshal(appErr)
	if err != nil {
		jsonResponse = []byte(fmt.Sprintf(`{"errors":[{"message":"unable to format json response for error"]}`))
	}
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

// handleErr is designed to conform to sendgrid's api guidelines for returning a single error
func handleErr(w http.ResponseWriter, statusCode int, message, field string) {
	errs := []client.Error{{Message: message, Field: field}}
	handleErrs(w, statusCode, errs)
}
