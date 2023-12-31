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

// MockTicketRepositoryI is a mock of TicketRepositoryI interface.
type MockTicketRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockTicketRepositoryIMockRecorder
}

// MockTicketRepositoryIMockRecorder is the mock recorder for MockTicketRepositoryI.
type MockTicketRepositoryIMockRecorder struct {
	mock *MockTicketRepositoryI
}

// NewMockTicketRepositoryI creates a new mock instance.
func NewMockTicketRepositoryI(ctrl *gomock.Controller) *MockTicketRepositoryI {
	mock := &MockTicketRepositoryI{ctrl: ctrl}
	mock.recorder = &MockTicketRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketRepositoryI) EXPECT() *MockTicketRepositoryIMockRecorder {
	return m.recorder
}

// CreateTicket mocks base method.
func (m *MockTicketRepositoryI) CreateTicket(ticket *model.Tickets) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicket", ticket)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreateTicket indicates an expected call of CreateTicket.
func (mr *MockTicketRepositoryIMockRecorder) CreateTicket(ticket any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicket", reflect.TypeOf((*MockTicketRepositoryI)(nil).CreateTicket), ticket)
}

// GetTicketById mocks base method.
func (m *MockTicketRepositoryI) GetTicketById(id *uuid.UUID) (*models.TicketDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketById", id)
	ret0, _ := ret[0].(*models.TicketDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetTicketById indicates an expected call of GetTicketById.
func (mr *MockTicketRepositoryIMockRecorder) GetTicketById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketById", reflect.TypeOf((*MockTicketRepositoryI)(nil).GetTicketById), id)
}

// ValidateTicket mocks base method.
func (m *MockTicketRepositoryI) ValidateTicket(id *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTicket", id)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// ValidateTicket indicates an expected call of ValidateTicket.
func (mr *MockTicketRepositoryIMockRecorder) ValidateTicket(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTicket", reflect.TypeOf((*MockTicketRepositoryI)(nil).ValidateTicket), id)
}
