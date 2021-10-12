// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/domain/repository/other.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOtherService is a mock of OtherService interface.
type MockOtherService struct {
	ctrl     *gomock.Controller
	recorder *MockOtherServiceMockRecorder
}

// MockOtherServiceMockRecorder is the mock recorder for MockOtherService.
type MockOtherServiceMockRecorder struct {
	mock *MockOtherService
}

// NewMockOtherService creates a new mock instance.
func NewMockOtherService(ctrl *gomock.Controller) *MockOtherService {
	mock := &MockOtherService{ctrl: ctrl}
	mock.recorder = &MockOtherServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOtherService) EXPECT() *MockOtherServiceMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *MockOtherService) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockOtherServiceMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockOtherService)(nil).Ping), ctx)
}
