// Code generated by MockGen. DO NOT EDIT.
// Source: ./server.go

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidator is a mock of Validator interface.
type MockValidator struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorMockRecorder
}

// MockValidatorMockRecorder is the mock recorder for MockValidator.
type MockValidatorMockRecorder struct {
	mock *MockValidator
}

// NewMockValidator creates a new mock instance.
func NewMockValidator(ctrl *gomock.Controller) *MockValidator {
	mock := &MockValidator{ctrl: ctrl}
	mock.recorder = &MockValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidator) EXPECT() *MockValidatorMockRecorder {
	return m.recorder
}

// ValidateGateway mocks base method.
func (m *MockValidator) ValidateGateway(ctx context.Context, gateway *Gateway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGateway", ctx, gateway)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateGateway indicates an expected call of ValidateGateway.
func (mr *MockValidatorMockRecorder) ValidateGateway(ctx, gateway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGateway", reflect.TypeOf((*MockValidator)(nil).ValidateGateway), ctx, gateway)
}

// ValidateHTTPRoute mocks base method.
func (m *MockValidator) ValidateHTTPRoute(ctx context.Context, route *HTTPRoute) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateHTTPRoute", ctx, route)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateHTTPRoute indicates an expected call of ValidateHTTPRoute.
func (mr *MockValidatorMockRecorder) ValidateHTTPRoute(ctx, route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateHTTPRoute", reflect.TypeOf((*MockValidator)(nil).ValidateHTTPRoute), ctx, route)
}

// ValidateTCPRoute mocks base method.
func (m *MockValidator) ValidateTCPRoute(ctx context.Context, route *TCPRoute) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTCPRoute", ctx, route)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateTCPRoute indicates an expected call of ValidateTCPRoute.
func (mr *MockValidatorMockRecorder) ValidateTCPRoute(ctx, route interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTCPRoute", reflect.TypeOf((*MockValidator)(nil).ValidateTCPRoute), ctx, route)
}
