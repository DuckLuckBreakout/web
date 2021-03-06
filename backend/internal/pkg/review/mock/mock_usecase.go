// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DuckLuckBreakout/web/backend/internal/pkg/review (interfaces: UseCase)

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

// AddNewReviewForProduct mocks base method.
func (m *MockUseCase) AddNewReviewForProduct(arg0 uint64, arg1 *models.Review) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewReviewForProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewReviewForProduct indicates an expected call of AddNewReviewForProduct.
func (mr *MockUseCaseMockRecorder) AddNewReviewForProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewReviewForProduct", reflect.TypeOf((*MockUseCase)(nil).AddNewReviewForProduct), arg0, arg1)
}

// CheckReviewUserRights mocks base method.
func (m *MockUseCase) CheckReviewUserRights(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckReviewUserRights", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckReviewUserRights indicates an expected call of CheckReviewUserRights.
func (mr *MockUseCaseMockRecorder) CheckReviewUserRights(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReviewUserRights", reflect.TypeOf((*MockUseCase)(nil).CheckReviewUserRights), arg0, arg1)
}

// GetReviewsByProductId mocks base method.
func (m *MockUseCase) GetReviewsByProductId(arg0 uint64, arg1 *models.PaginatorReviews) (*models.RangeReviews, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewsByProductId", arg0, arg1)
	ret0, _ := ret[0].(*models.RangeReviews)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviewsByProductId indicates an expected call of GetReviewsByProductId.
func (mr *MockUseCaseMockRecorder) GetReviewsByProductId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewsByProductId", reflect.TypeOf((*MockUseCase)(nil).GetReviewsByProductId), arg0, arg1)
}

// GetStatisticsByProductId mocks base method.
func (m *MockUseCase) GetStatisticsByProductId(arg0 uint64) (*models.ReviewStatistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatisticsByProductId", arg0)
	ret0, _ := ret[0].(*models.ReviewStatistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatisticsByProductId indicates an expected call of GetStatisticsByProductId.
func (mr *MockUseCaseMockRecorder) GetStatisticsByProductId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatisticsByProductId", reflect.TypeOf((*MockUseCase)(nil).GetStatisticsByProductId), arg0)
}
