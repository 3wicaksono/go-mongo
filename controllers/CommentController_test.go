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

func TestCommentController_Publish(t *testing.T) {
	var (
		commentService = new(mocks.ICommentService)
		controller     = new(CommentController)

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
			name:    "success",
			payload: models.CommentModel{Comment: "Looking to hire SE Asia's top dev talent!"},
			returnService: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
			expectedResponse: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
		},
	}

	controller.CommentService = commentService

	url = "/orgs/org-name/comments"
	router.HandleFunc(url, controller.Publish).Methods(http.MethodPost)
	rw := httptest.NewRecorder()

	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			commentService.On("Publish", testCase.payload).Return(testCase.returnService)

			req = httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(helpers.MakeJSON(testCase.payload))))
			router.ServeHTTP(rw, req)
			testCase.actualResponse.HTTPStatus = rw.Code
			assert.Equal(t, testCase.expectedResponse.BodyResponse, testCase.actualResponse.BodyResponse)
		})
	}
}

func TestCommentController_GetAll(t *testing.T) {
	var (
		commentService = new(mocks.ICommentService)
		controller     = new(CommentController)

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
			payload: models.CommentModel{OrgName: "org-name"},
			returnService: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
			expectedResponse: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
		},
	}

	controller.CommentService = commentService

	url = "/orgs/org-name/comments"
	router.HandleFunc(url, controller.GetAll).Methods(http.MethodGet)
	rw := httptest.NewRecorder()

	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			commentService.On("GetAll", mock.Anything).Return(testCase.returnService)

			req = httptest.NewRequest(http.MethodGet, url, nil)
			router.ServeHTTP(rw, req)
			testCase.actualResponse.HTTPStatus = rw.Code
			assert.Equal(t, testCase.expectedResponse.BodyResponse, testCase.actualResponse.BodyResponse)
		})
	}
}

func TestCommentController_Delete(t *testing.T) {
	var (
		commentService = new(mocks.ICommentService)
		controller     = new(CommentController)

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
			payload: models.CommentModel{OrgName: "org-name"},
			returnService: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
			expectedResponse: models.APIResponseModel{
				HTTPStatus: http.StatusOK,
			},
		},
	}

	controller.CommentService = commentService

	url = "/orgs/org-name/comments"
	router.HandleFunc(url, controller.Delete).Methods(http.MethodDelete)
	rw := httptest.NewRecorder()

	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			commentService.On("Delete", mock.Anything).Return(testCase.returnService)

			req = httptest.NewRequest(http.MethodDelete, url, nil)
			router.ServeHTTP(rw, req)
			testCase.actualResponse.HTTPStatus = rw.Code
			assert.Equal(t, testCase.expectedResponse.BodyResponse, testCase.actualResponse.BodyResponse)
		})
	}
}
