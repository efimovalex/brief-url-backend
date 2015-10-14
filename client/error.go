package client

/*
   Client code to help facilite go services with interacting with your service.
   Below are errors that other services may receive from this service
*/

const (
	ErrorMessageInternalService  = "an internal service had an error"
	ErrorMessageBadPayload       = "bad json payload"
	ErrorMessageResourceNotFound = "resource not found"
)

// ErrorResult is the struct clients should consume
type ErrorResult struct {
	Errors []Error `json:"errors"`
}

// Error represents individual errors the client may get
type Error struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}
