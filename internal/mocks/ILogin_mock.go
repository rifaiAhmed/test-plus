// Code generated by MockGen. DO NOT EDIT.
// Source: ILogin.go
//
// Generated by this command:
//
//	mockgen -source=ILogin.go -destination=../mocks/ILogin_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	models "test-plus/internal/model"

	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockILoginService is a mock of ILoginService interface.
type MockILoginService struct {
	ctrl     *gomock.Controller
	recorder *MockILoginServiceMockRecorder
	isgomock struct{}
}

// MockILoginServiceMockRecorder is the mock recorder for MockILoginService.
type MockILoginServiceMockRecorder struct {
	mock *MockILoginService
}

// NewMockILoginService creates a new mock instance.
func NewMockILoginService(ctrl *gomock.Controller) *MockILoginService {
	mock := &MockILoginService{ctrl: ctrl}
	mock.recorder = &MockILoginServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILoginService) EXPECT() *MockILoginServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockILoginService) Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, req)
	ret0, _ := ret[0].(models.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockILoginServiceMockRecorder) Login(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockILoginService)(nil).Login), ctx, req)
}

// MockILoginHandler is a mock of ILoginHandler interface.
type MockILoginHandler struct {
	ctrl     *gomock.Controller
	recorder *MockILoginHandlerMockRecorder
	isgomock struct{}
}

// MockILoginHandlerMockRecorder is the mock recorder for MockILoginHandler.
type MockILoginHandlerMockRecorder struct {
	mock *MockILoginHandler
}

// NewMockILoginHandler creates a new mock instance.
func NewMockILoginHandler(ctrl *gomock.Controller) *MockILoginHandler {
	mock := &MockILoginHandler{ctrl: ctrl}
	mock.recorder = &MockILoginHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILoginHandler) EXPECT() *MockILoginHandlerMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockILoginHandler) Login(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Login", c)
}

// Login indicates an expected call of Login.
func (mr *MockILoginHandlerMockRecorder) Login(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockILoginHandler)(nil).Login), c)
}
