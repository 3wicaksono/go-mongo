package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-mongo/helpers"
	"go-mongo/interfaces"
	"go-mongo/models"
)

// CommentController struct
type CommentController struct {
	CommentService interfaces.ICommentService
}

// Publish create comment
func (c *CommentController) Publish(response http.ResponseWriter, request *http.Request) {
	var (
		params  = mux.Vars(request)
		comment models.CommentModel
		result  models.APIResponseModel
	)

	err := json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		helpers.Response(response, http.StatusInternalServerError, models.ResponseModel{Message: err.Error()})
		return
	}

	comment.OrgName = params["org-name"]
	result = c.CommentService.Publish(comment)
	helpers.Response(response, result.HTTPStatus, result.BodyResponse)
	return
}

// GetAll get list all comment by org name
func (c *CommentController) GetAll(response http.ResponseWriter, request *http.Request) {
	var (
		params = mux.Vars(request)
		result = c.CommentService.GetAll(models.CommentModel{
			OrgName: params["org-name"],
		})
	)

	helpers.Response(response, result.HTTPStatus, result.BodyResponse)
	return
}

// Delete soft delete all comment by org name
func (c *CommentController) Delete(response http.ResponseWriter, request *http.Request) {
	var (
		params = mux.Vars(request)
		result = c.CommentService.Delete(models.CommentModel{
			OrgName: params["org-name"],
		})
	)

	helpers.Response(response, result.HTTPStatus, result.BodyResponse)
	return
}
