// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/DuckLuckBreakout/web/backend/internal/pkg/review (interfaces: Repository)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/DuckLuckBreakout/web/backend/internal/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddReview mocks base method.
func (m *MockRepository) AddReview(arg0 uint64, arg1 *models.Review) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddReview", arg0, arg1)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddReview indicates an expected call of AddReview.
func (mr *MockRepositoryMockRecorder) AddReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReview", reflect.TypeOf((*MockRepository)(nil).AddReview), arg0, arg1)
}

// CheckReview mocks base method.
func (m *MockRepository) CheckReview(arg0, arg1 uint64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckReview", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckReview indicates an expected call of CheckReview.
func (mr *MockRepositoryMockRecorder) CheckReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckReview", reflect.TypeOf((*MockRepository)(nil).CheckReview), arg0, arg1)
}

// CreateSortString mocks base method.
func (m *MockRepository) CreateSortString(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSortString", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSortString indicates an expected call of CreateSortString.
func (mr *MockRepositoryMockRecorder) CreateSortString(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSortString", reflect.TypeOf((*MockRepository)(nil).CreateSortString), arg0, arg1)
}

// GetCountPages mocks base method.
func (m *MockRepository) GetCountPages(arg0 uint64, arg1 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCountPages", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCountPages indicates an expected call of GetCountPages.
func (mr *MockRepositoryMockRecorder) GetCountPages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCountPages", reflect.TypeOf((*MockRepository)(nil).GetCountPages), arg0, arg1)
}

// SelectRangeReviews mocks base method.
func (m *MockRepository) SelectRangeReviews(arg0 uint64, arg1 string, arg2 *models.PaginatorReviews) ([]*models.ViewReview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectRangeReviews", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.ViewReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectRangeReviews indicates an expected call of SelectRangeReviews.
func (mr *MockRepositoryMockRecorder) SelectRangeReviews(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectRangeReviews", reflect.TypeOf((*MockRepository)(nil).SelectRangeReviews), arg0, arg1, arg2)
}

// SelectStatisticsByProductId mocks base method.
func (m *MockRepository) SelectStatisticsByProductId(arg0 uint64) (*models.ReviewStatistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectStatisticsByProductId", arg0)
	ret0, _ := ret[0].(*models.ReviewStatistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectStatisticsByProductId indicates an expected call of SelectStatisticsByProductId.
func (mr *MockRepositoryMockRecorder) SelectStatisticsByProductId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectStatisticsByProductId", reflect.TypeOf((*MockRepository)(nil).SelectStatisticsByProductId), arg0)
}
