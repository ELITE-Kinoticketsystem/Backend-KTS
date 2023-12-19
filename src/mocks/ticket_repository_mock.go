// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/ticket_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/ticket_repository.go -destination=./src/mocks/ticket_repository_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockTicketRepoI is a mock of TicketRepoI interface.
type MockTicketRepoI struct {
	ctrl     *gomock.Controller
	recorder *MockTicketRepoIMockRecorder
}

// MockTicketRepoIMockRecorder is the mock recorder for MockTicketRepoI.
type MockTicketRepoIMockRecorder struct {
	mock *MockTicketRepoI
}

// NewMockTicketRepoI creates a new mock instance.
func NewMockTicketRepoI(ctrl *gomock.Controller) *MockTicketRepoI {
	mock := &MockTicketRepoI{ctrl: ctrl}
	mock.recorder = &MockTicketRepoIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketRepoI) EXPECT() *MockTicketRepoIMockRecorder {
	return m.recorder
}

// CreateTicket mocks base method.
func (m *MockTicketRepoI) CreateTicket(ticket *model.Tickets) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicket", ticket)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreateTicket indicates an expected call of CreateTicket.
func (mr *MockTicketRepoIMockRecorder) CreateTicket(ticket any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicket", reflect.TypeOf((*MockTicketRepoI)(nil).CreateTicket), ticket)
}
