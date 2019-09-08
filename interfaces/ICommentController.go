package interfaces

import "net/http"

// ICommentController interface
type ICommentController interface {
	Publish(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
}
