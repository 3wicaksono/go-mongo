package routes

import (
	"net/http"

	"go-mongo/interfaces"

	"github.com/gorilla/mux"
)

// Route struct
type Route struct {
	CommentController interfaces.ICommentController
	MemberController  interfaces.IMemberController
}

// GetRoute init route list
func (r *Route) GetRoute() http.Handler {
	var (
		router         = mux.NewRouter()
		pathOrgComment = "/orgs/{org-name}/comments"
		pathOrgMember  = "/orgs/{org-name}/members"
	)

	router.HandleFunc(pathOrgComment, r.CommentController.Publish).Methods(http.MethodPost)
	router.HandleFunc(pathOrgComment, r.CommentController.GetAll).Methods(http.MethodGet)
	router.HandleFunc(pathOrgComment, r.CommentController.Delete).Methods(http.MethodDelete)

	router.HandleFunc(pathOrgMember, r.MemberController.Publish).Methods(http.MethodPost)
	router.HandleFunc(pathOrgMember, r.MemberController.GetAll).Methods(http.MethodGet)
	router.HandleFunc(pathOrgMember, r.MemberController.Delete).Methods(http.MethodDelete)

	return router
}
