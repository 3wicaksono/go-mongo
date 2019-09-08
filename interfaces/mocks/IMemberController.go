// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import http "net/http"

import mock "github.com/stretchr/testify/mock"

// IMemberController is an autogenerated mock type for the IMemberController type
type IMemberController struct {
	mock.Mock
}

// Delete provides a mock function with given fields: response, request
func (_m *IMemberController) Delete(response http.ResponseWriter, request *http.Request) {
	_m.Called(response, request)
}

// GetAll provides a mock function with given fields: response, request
func (_m *IMemberController) GetAll(response http.ResponseWriter, request *http.Request) {
	_m.Called(response, request)
}

// Publish provides a mock function with given fields: response, request
func (_m *IMemberController) Publish(response http.ResponseWriter, request *http.Request) {
	_m.Called(response, request)
}
