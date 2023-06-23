// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package tx is a generated GoMock package.
package tx

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockContextBeginner is a mock of ContextBeginner interface.
type MockContextBeginner struct {
	ctrl     *gomock.Controller
	recorder *MockContextBeginnerMockRecorder
}

// MockContextBeginnerMockRecorder is the mock recorder for MockContextBeginner.
type MockContextBeginnerMockRecorder struct {
	mock *MockContextBeginner
}

// NewMockContextBeginner creates a new mock instance.
func NewMockContextBeginner(ctrl *gomock.Controller) *MockContextBeginner {
	mock := &MockContextBeginner{ctrl: ctrl}
	mock.recorder = &MockContextBeginnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContextBeginner) EXPECT() *MockContextBeginnerMockRecorder {
	return m.recorder
}

// BeginTx mocks base method.
func (m *MockContextBeginner) BeginTx(ctx context.Context, fn func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockContextBeginnerMockRecorder) BeginTx(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockContextBeginner)(nil).BeginTx), ctx, fn)
}
