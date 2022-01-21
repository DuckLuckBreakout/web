// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DuckLuckBreakout/web/backend/internal/pkg/admin (interfaces: UseCase)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/DuckLuckBreakout/web/backend/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// ChangeOrderStatus mocks base method.
func (m *MockUseCase) ChangeOrderStatus(arg0 *models.UpdateOrder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeOrderStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeOrderStatus indicates an expected call of ChangeOrderStatus.
func (mr *MockUseCaseMockRecorder) ChangeOrderStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderStatus", reflect.TypeOf((*MockUseCase)(nil).ChangeOrderStatus), arg0)
}