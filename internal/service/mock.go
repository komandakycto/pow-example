// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package service is a generated GoMock package.
package service

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockQuotesService is a mock of QuotesService interface
type MockQuotesService struct {
	ctrl     *gomock.Controller
	recorder *MockQuotesServiceMockRecorder
}

// MockQuotesServiceMockRecorder is the mock recorder for MockQuotesService
type MockQuotesServiceMockRecorder struct {
	mock *MockQuotesService
}

// NewMockQuotesService creates a new mock instance
func NewMockQuotesService(ctrl *gomock.Controller) *MockQuotesService {
	mock := &MockQuotesService{ctrl: ctrl}
	mock.recorder = &MockQuotesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuotesService) EXPECT() *MockQuotesServiceMockRecorder {
	return m.recorder
}

// GetRandomQuote mocks base method
func (m *MockQuotesService) GetRandomQuote() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomQuote")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomQuote indicates an expected call of GetRandomQuote
func (mr *MockQuotesServiceMockRecorder) GetRandomQuote() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomQuote", reflect.TypeOf((*MockQuotesService)(nil).GetRandomQuote))
}

// MockHashcashService is a mock of HashcashService interface
type MockHashcashService struct {
	ctrl     *gomock.Controller
	recorder *MockHashcashServiceMockRecorder
}

// MockHashcashServiceMockRecorder is the mock recorder for MockHashcashService
type MockHashcashServiceMockRecorder struct {
	mock *MockHashcashService
}

// NewMockHashcashService creates a new mock instance
func NewMockHashcashService(ctrl *gomock.Controller) *MockHashcashService {
	mock := &MockHashcashService{ctrl: ctrl}
	mock.recorder = &MockHashcashServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHashcashService) EXPECT() *MockHashcashServiceMockRecorder {
	return m.recorder
}

// GenerateChallenge mocks base method
func (m *MockHashcashService) GenerateChallenge() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateChallenge")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateChallenge indicates an expected call of GenerateChallenge
func (mr *MockHashcashServiceMockRecorder) GenerateChallenge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateChallenge", reflect.TypeOf((*MockHashcashService)(nil).GenerateChallenge))
}

// VerifyPoW mocks base method
func (m *MockHashcashService) VerifyPoW(challenge, nonce string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPoW", challenge, nonce)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifyPoW indicates an expected call of VerifyPoW
func (mr *MockHashcashServiceMockRecorder) VerifyPoW(challenge, nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPoW", reflect.TypeOf((*MockHashcashService)(nil).VerifyPoW), challenge, nonce)
}
