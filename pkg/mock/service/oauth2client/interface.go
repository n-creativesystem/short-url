// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package oauth2client is a generated GoMock package.
package oauth2client

import (
	context "context"
	reflect "reflect"

	oauth2 "github.com/go-oauth2/oauth2/v4"
	gomock "github.com/golang/mock/gomock"
	oauth2client "github.com/n-creativesystem/short-url/pkg/domain/oauth2_client"
	oauth2client0 "github.com/n-creativesystem/short-url/pkg/service/oauth2_client"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// DeleteClient mocks base method.
func (m *MockService) DeleteClient(ctx context.Context, user, clientId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClient", ctx, user, clientId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClient indicates an expected call of DeleteClient.
func (mr *MockServiceMockRecorder) DeleteClient(ctx, user, clientId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClient", reflect.TypeOf((*MockService)(nil).DeleteClient), ctx, user, clientId)
}

// FindAll mocks base method.
func (m *MockService) FindAll(ctx context.Context, user string) ([]oauth2client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, user)
	ret0, _ := ret[0].([]oauth2client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockServiceMockRecorder) FindAll(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockService)(nil).FindAll), ctx, user)
}

// FindByID mocks base method.
func (m *MockService) FindByID(ctx context.Context, id, user string) (*oauth2client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id, user)
	ret0, _ := ret[0].(*oauth2client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockServiceMockRecorder) FindByID(ctx, id, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockService)(nil).FindByID), ctx, id, user)
}

// GetByID mocks base method.
func (m *MockService) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(oauth2.ClientInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockServiceMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockService)(nil).GetByID), ctx, id)
}

// RegisterClient mocks base method.
func (m *MockService) RegisterClient(ctx context.Context, user, appName string) (oauth2client0.RegisterResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterClient", ctx, user, appName)
	ret0, _ := ret[0].(oauth2client0.RegisterResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterClient indicates an expected call of RegisterClient.
func (mr *MockServiceMockRecorder) RegisterClient(ctx, user, appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterClient", reflect.TypeOf((*MockService)(nil).RegisterClient), ctx, user, appName)
}

// UpdateClient mocks base method.
func (m *MockService) UpdateClient(ctx context.Context, id, user, appName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateClient", ctx, id, user, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateClient indicates an expected call of UpdateClient.
func (mr *MockServiceMockRecorder) UpdateClient(ctx, id, user, appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateClient", reflect.TypeOf((*MockService)(nil).UpdateClient), ctx, id, user, appName)
}
