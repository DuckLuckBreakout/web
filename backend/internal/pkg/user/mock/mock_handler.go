// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DuckLuckBreakout/web/backend/internal/pkg/user (interfaces: Handler)

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// GetProfile mocks base method.
func (m *MockHandler) GetProfile(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetProfile", arg0, arg1)
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockHandlerMockRecorder) GetProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockHandler)(nil).GetProfile), arg0, arg1)
}

// GetProfileAvatar mocks base method.
func (m *MockHandler) GetProfileAvatar(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetProfileAvatar", arg0, arg1)
}

// GetProfileAvatar indicates an expected call of GetProfileAvatar.
func (mr *MockHandlerMockRecorder) GetProfileAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileAvatar", reflect.TypeOf((*MockHandler)(nil).GetProfileAvatar), arg0, arg1)
}

// Login mocks base method.
func (m *MockHandler) Login(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Login", arg0, arg1)
}

// Login indicates an expected call of Login.
func (mr *MockHandlerMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockHandler)(nil).Login), arg0, arg1)
}

// Logout mocks base method.
func (m *MockHandler) Logout(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Logout", arg0, arg1)
}

// Logout indicates an expected call of Logout.
func (mr *MockHandlerMockRecorder) Logout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockHandler)(nil).Logout), arg0, arg1)
}

// Signup mocks base method.
func (m *MockHandler) Signup(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Signup", arg0, arg1)
}

// Signup indicates an expected call of Signup.
func (mr *MockHandlerMockRecorder) Signup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockHandler)(nil).Signup), arg0, arg1)
}

// UpdateProfile mocks base method.
func (m *MockHandler) UpdateProfile(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockHandlerMockRecorder) UpdateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockHandler)(nil).UpdateProfile), arg0, arg1)
}

// UpdateProfileAvatar mocks base method.
func (m *MockHandler) UpdateProfileAvatar(arg0 http.ResponseWriter, arg1 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateProfileAvatar", arg0, arg1)
}

// UpdateProfileAvatar indicates an expected call of UpdateProfileAvatar.
func (mr *MockHandlerMockRecorder) UpdateProfileAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfileAvatar", reflect.TypeOf((*MockHandler)(nil).UpdateProfileAvatar), arg0, arg1)
}
