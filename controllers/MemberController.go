package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-mongo/helpers"
	"go-mongo/interfaces"
	"go-mongo/models"
)

// MemberController struct
type MemberController struct {
	MemberService interfaces.IMemberService
}

// Publish create member
func (c *MemberController) Publish(response http.ResponseWriter, request *http.Request) {
	var (
		params = mux.Vars(request)
		member models.MemberModel
		result models.APIResponseModel
	)

	err := json.NewDecoder(request.Body).Decode(&member)
	if err != nil {
		helpers.Response(response, http.StatusInternalServerError, models.ResponseModel{Message: err.Error()})
		return
	}

	member.OrgName = params["org-name"]
	result = c.MemberService.Publish(member)
	helpers.Response(response, result.HTTPStatus, result.BodyResponse)
	return
}

// GetAll get list all member by org name
func (c *MemberController) GetAll(response http.ResponseWriter, request *http.Request) {
	var (
		params = mux.Vars(request)
		result = c.MemberService.GetAll(models.MemberModel{
			OrgName: params["org-name"],
		})
	)

	helpers.Response(response, result.HTTPStatus, result.BodyResponse)
	return
}

// Delete soft delete all member by org name
func (c *MemberController) Delete(response http.ResponseWriter, request *http.Request) {
	var (
		params = mux.Vars(request)
		result = c.MemberService.Delete(models.MemberModel{
			OrgName: params["org-name"],
		})
	)

	helpers.Response(response, result.HTTPStatus, result.BodyResponse)
	return
}
