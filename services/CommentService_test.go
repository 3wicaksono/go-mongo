package services

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-mongo/interfaces/mocks"
	"go-mongo/models"
)

func TestCommentService_Publish(t *testing.T) {
	var (
		repo    = new(mocks.ICommentRepository)
		service = new(CommentService)
	)

	type testObject struct {
		name             string
		payload          models.CommentModel
		resultRepoObject models.CommentModel
		resultRepoErr    error
		expectedHTTPCode int
	}

	testScenario := []testObject{
		testObject{
			name:             "invalid validation",
			payload:          models.CommentModel{},
			resultRepoObject: models.CommentModel{},
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusBadRequest,
		},
		testObject{
			name:             "success",
			payload:          models.CommentModel{Comment: "abc", OrgName: "org"},
			resultRepoObject: models.CommentModel{},
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusOK,
		},
		testObject{
			name:             "error",
			payload:          models.CommentModel{Comment: "abc2", OrgName: "org"},
			resultRepoObject: models.CommentModel{},
			resultRepoErr:    errors.New("error"),
			expectedHTTPCode: http.StatusInternalServerError,
		},
	}

	service.CommentRepository = repo
	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("Publish", testCase.payload).Return(testCase.resultRepoObject, testCase.resultRepoErr)
			rsp := service.Publish(testCase.payload)
			assert.Equal(t, testCase.expectedHTTPCode, rsp.HTTPStatus)
		})
	}
}

func TestCommentService_GetAll(t *testing.T) {
	var (
		repo    = new(mocks.ICommentRepository)
		service = new(CommentService)
	)

	type testObject struct {
		name             string
		payload          models.CommentModel
		resultRepoObject interface{}
		resultRepoErr    error
		expectedHTTPCode int
	}

	testScenario := []testObject{
		testObject{
			name:             "invalid validation",
			payload:          models.CommentModel{OrgName: "&^*&^*"},
			resultRepoObject: []models.CommentModel{},
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusBadRequest,
		},
		testObject{
			name:             "error",
			payload:          models.CommentModel{Comment: "abc", OrgName: "org"},
			resultRepoObject: []models.CommentModel{},
			resultRepoErr:    errors.New("error"),
			expectedHTTPCode: http.StatusInternalServerError,
		},
		testObject{
			name:             "error",
			payload:          models.CommentModel{Comment: "abc1", OrgName: "org"},
			resultRepoObject: nil,
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusNotFound,
		},
		testObject{
			name:             "success",
			payload:          models.CommentModel{Comment: "abc2", OrgName: "org"},
			resultRepoObject: []models.CommentModel{},
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusOK,
		},
	}

	service.CommentRepository = repo
	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("GetAll", testCase.payload).Return(testCase.resultRepoObject, testCase.resultRepoErr)
			rsp := service.GetAll(testCase.payload)
			assert.Equal(t, testCase.expectedHTTPCode, rsp.HTTPStatus)
		})
	}
}

func TestCommentService_Delete(t *testing.T) {
	var (
		repo    = new(mocks.ICommentRepository)
		service = new(CommentService)
	)

	type testObject struct {
		name             string
		payload          models.CommentModel
		resultRepoObject int
		resultRepoErr    error
		expectedHTTPCode int
	}

	testScenario := []testObject{
		testObject{
			name:             "invalid validation",
			payload:          models.CommentModel{OrgName: "&^*&^*"},
			resultRepoObject: 0,
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusBadRequest,
		},
		testObject{
			name:             "error",
			payload:          models.CommentModel{Comment: "abc", OrgName: "org"},
			resultRepoObject: 0,
			resultRepoErr:    errors.New("error"),
			expectedHTTPCode: http.StatusInternalServerError,
		},
		testObject{
			name:             "error",
			payload:          models.CommentModel{Comment: "abc1", OrgName: "org"},
			resultRepoObject: 0,
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusNotFound,
		},
		testObject{
			name:             "success",
			payload:          models.CommentModel{Comment: "abc2", OrgName: "org"},
			resultRepoObject: 1,
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusOK,
		},
	}

	service.CommentRepository = repo
	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("Delete", testCase.payload).Return(testCase.resultRepoObject, testCase.resultRepoErr)
			rsp := service.Delete(testCase.payload)
			assert.Equal(t, testCase.expectedHTTPCode, rsp.HTTPStatus)
		})
	}
}
