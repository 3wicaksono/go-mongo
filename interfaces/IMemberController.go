package interfaces

import "net/http"

// IMemberController interface
type IMemberController interface {
	Publish(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
}
