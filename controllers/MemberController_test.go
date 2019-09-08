package controllers

import (
	"bytes"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"go-mongo/helpers"
	"go-mongo/interfaces/mocks"
	"go-mongo/models"
)

func TestMemberController_Publish(t *testing.T) {
	var (
		memberService = new(mocks.IMemberService)
		controller    = new(MemberController)

		router = mux.NewRouter()
		req    *http.Request
		url    string
	)

	type testObject struct {
		name                                            string
		payload                                         interface{}
		returnService, expectedResponse, actualResponse models.APIResponseModel
	}

	testScenario := []testObject{
		testObject{
			name:    "error",
			payload: "abc",
			returnService: models.APIResponseModel{
				HTTPStatus: http.StatusInternalServerError,
			},
			expectedResponse: models.APIResponseModel{
				HTTPStatus: http.StatusInternalServerError,
			},
		},
		testObject{
			name: "success",
			payload: models.MemberModel{
				Username:       "abc",
				AvatarURL:      "abc",
				TotalFollower:  1,
				TotalFollowing: 1,
			},
			returnService: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
			expectedResponse: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
		},
	}

	controller.MemberService = memberService

	url = "/orgs/org-name/members"
	router.HandleFunc(url, controller.Publish).Methods(http.MethodPost)
	rw := httptest.NewRecorder()

	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			memberService.On("Publish", testCase.payload).Return(testCase.returnService)

			req = httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(helpers.MakeJSON(testCase.payload))))
			router.ServeHTTP(rw, req)
			testCase.actualResponse.HTTPStatus = rw.Code
			assert.Equal(t, testCase.expectedResponse.BodyResponse, testCase.actualResponse.BodyResponse)
		})
	}
}

func TestMemberController_GetAll(t *testing.T) {
	var (
		memberService = new(mocks.IMemberService)
		controller    = new(MemberController)

		router = mux.NewRouter()
		req    *http.Request
		url    string
	)

	type testObject struct {
		name                                            string
		payload                                         interface{}
		returnService, expectedResponse, actualResponse models.APIResponseModel
	}

	testScenario := []testObject{
		testObject{
			name:    "success",
			payload: models.MemberModel{OrgName: "org-name"},
			returnService: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
			expectedResponse: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
		},
	}

	controller.MemberService = memberService

	url = "/orgs/org-name/members"
	router.HandleFunc(url, controller.GetAll).Methods(http.MethodGet)
	rw := httptest.NewRecorder()

	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			memberService.On("GetAll", mock.Anything).Return(testCase.returnService)

			req = httptest.NewRequest(http.MethodGet, url, nil)
			router.ServeHTTP(rw, req)
			testCase.actualResponse.HTTPStatus = rw.Code
			assert.Equal(t, testCase.expectedResponse.BodyResponse, testCase.actualResponse.BodyResponse)
		})
	}
}

func TestMemberController_Delete(t *testing.T) {
	var (
		memberService = new(mocks.IMemberService)
		controller    = new(MemberController)

		router = mux.NewRouter()
		req    *http.Request
		url    string
	)

	type testObject struct {
		name                                            string
		payload                                         interface{}
		returnService, expectedResponse, actualResponse models.APIResponseModel
	}

	testScenario := []testObject{
		testObject{
			name:    "success",
			payload: models.MemberModel{OrgName: "org-name"},
			returnService: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
			expectedResponse: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
		},
	}

	controller.MemberService = memberService

	url = "/orgs/org-name/members"
	router.HandleFunc(url, controller.Delete).Methods(http.MethodDelete)
	rw := httptest.NewRecorder()

	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			memberService.On("Delete", mock.Anything).Return(testCase.returnService)

			req = httptest.NewRequest(http.MethodDelete, url, nil)
			router.ServeHTTP(rw, req)
			testCase.actualResponse.HTTPStatus = rw.Code
			assert.Equal(t, testCase.expectedResponse.BodyResponse, testCase.actualResponse.BodyResponse)
		})
	}
}
