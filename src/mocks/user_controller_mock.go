// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/controllers/user_controller.go
//
// Generated by this command:
//
//	mockgen -source=./src/controllers/user_controller.go -destination=./src/mocks/user_controller_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUserControllerI is a mock of UserControllerI interface.
type MockUserControllerI struct {
	ctrl     *gomock.Controller
	recorder *MockUserControllerIMockRecorder
}

// MockUserControllerIMockRecorder is the mock recorder for MockUserControllerI.
type MockUserControllerIMockRecorder struct {
	mock *MockUserControllerI
}

// NewMockUserControllerI creates a new mock instance.
func NewMockUserControllerI(ctrl *gomock.Controller) *MockUserControllerI {
	mock := &MockUserControllerI{ctrl: ctrl}
	mock.recorder = &MockUserControllerIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserControllerI) EXPECT() *MockUserControllerIMockRecorder {
	return m.recorder
}

// RegisterUser mocks base method.
func (m *MockUserControllerI) RegisterUser(registrationData models.RegistrationRequest) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", registrationData)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockUserControllerIMockRecorder) RegisterUser(registrationData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockUserControllerI)(nil).RegisterUser), registrationData)
}
