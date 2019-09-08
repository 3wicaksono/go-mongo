package services

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-mongo/interfaces/mocks"
	"go-mongo/models"
)

func TestMemberService_Publish(t *testing.T) {
	var (
		repo    = new(mocks.IMemberRepository)
		service = new(MemberService)
	)

	type testObject struct {
		name             string
		payload          models.MemberModel
		resultRepoObject models.MemberModel
		resultRepoErr    error
		expectedHTTPCode int
	}

	testScenario := []testObject{
		testObject{
			name:             "invalid validation",
			payload:          models.MemberModel{},
			resultRepoObject: models.MemberModel{},
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusBadRequest,
		},
		testObject{
			name: "success",
			payload: models.MemberModel{
				OrgName:        "org",
				Username:       "abc",
				AvatarURL:      "abc",
				TotalFollower:  1,
				TotalFollowing: 1,
			},
			resultRepoObject: models.MemberModel{},
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusOK,
		},
		testObject{
			name: "error",
			payload: models.MemberModel{
				OrgName:        "org",
				Username:       "abc1",
				AvatarURL:      "abc",
				TotalFollower:  1,
				TotalFollowing: 1,
			},
			resultRepoObject: models.MemberModel{},
			resultRepoErr:    errors.New("error"),
			expectedHTTPCode: http.StatusInternalServerError,
		},
	}

	service.MemberRepository = repo
	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("Publish", testCase.payload).Return(testCase.resultRepoObject, testCase.resultRepoErr)
			rsp := service.Publish(testCase.payload)
			assert.Equal(t, testCase.expectedHTTPCode, rsp.HTTPStatus)
		})
	}
}

func TestMemberService_GetAll(t *testing.T) {
	var (
		repo    = new(mocks.IMemberRepository)
		service = new(MemberService)
	)

	type testObject struct {
		name             string
		payload          models.MemberModel
		resultRepoObject interface{}
		resultRepoErr    error
		expectedHTTPCode int
	}

	testScenario := []testObject{
		testObject{
			name:             "invalid validation",
			payload:          models.MemberModel{OrgName: "&^*&^*"},
			resultRepoObject: []models.MemberModel{},
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusBadRequest,
		},
		testObject{
			name: "error",
			payload: models.MemberModel{
				OrgName:        "org",
				Username:       "abc",
				AvatarURL:      "abc",
				TotalFollower:  1,
				TotalFollowing: 1,
			},
			resultRepoObject: []models.MemberModel{},
			resultRepoErr:    errors.New("error"),
			expectedHTTPCode: http.StatusInternalServerError,
		},
		testObject{
			name: "error",
			payload: models.MemberModel{
				OrgName:        "org",
				Username:       "abc1",
				AvatarURL:      "abc",
				TotalFollower:  1,
				TotalFollowing: 1,
			},
			resultRepoObject: nil,
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusNotFound,
		},
		testObject{
			name: "success",
			payload: models.MemberModel{
				OrgName:        "org",
				Username:       "abc3",
				AvatarURL:      "abc",
				TotalFollower:  1,
				TotalFollowing: 1,
			},
			resultRepoObject: []models.MemberModel{},
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusOK,
		},
	}

	service.MemberRepository = repo
	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("GetAll", testCase.payload).Return(testCase.resultRepoObject, testCase.resultRepoErr)
			rsp := service.GetAll(testCase.payload)
			assert.Equal(t, testCase.expectedHTTPCode, rsp.HTTPStatus)
		})
	}
}

func TestMemberService_Delete(t *testing.T) {
	var (
		repo    = new(mocks.IMemberRepository)
		service = new(MemberService)
	)

	type testObject struct {
		name             string
		payload          models.MemberModel
		resultRepoObject int
		resultRepoErr    error
		expectedHTTPCode int
	}

	testScenario := []testObject{
		testObject{
			name:             "invalid validation",
			payload:          models.MemberModel{OrgName: "&^*&^*"},
			resultRepoObject: 0,
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusBadRequest,
		},
		testObject{
			name: "error",
			payload: models.MemberModel{
				OrgName:        "org",
				Username:       "abc",
				AvatarURL:      "abc",
				TotalFollower:  1,
				TotalFollowing: 1,
			},
			resultRepoObject: 0,
			resultRepoErr:    errors.New("error"),
			expectedHTTPCode: http.StatusInternalServerError,
		},
		testObject{
			name: "error",
			payload: models.MemberModel{
				OrgName:        "org",
				Username:       "abc1",
				AvatarURL:      "abc",
				TotalFollower:  1,
				TotalFollowing: 1,
			},
			resultRepoObject: 0,
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusNotFound,
		},
		testObject{
			name: "success",
			payload: models.MemberModel{
				OrgName:        "org",
				Username:       "abc2",
				AvatarURL:      "abc",
				TotalFollower:  1,
				TotalFollowing: 1,
			},
			resultRepoObject: 1,
			resultRepoErr:    nil,
			expectedHTTPCode: http.StatusOK,
		},
	}

	service.MemberRepository = repo
	for _, testCase := range testScenario {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("Delete", testCase.payload).Return(testCase.resultRepoObject, testCase.resultRepoErr)
			rsp := service.Delete(testCase.payload)
			assert.Equal(t, testCase.expectedHTTPCode, rsp.HTTPStatus)
		})
	}
}
