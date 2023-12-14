// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/theatre_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/theatre_repository.go -destination=./src/mocks/theatre_repository_mock.go -package=mocks
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

// MockTheaterRepoI is a mock of TheaterRepoI interface.
type MockTheaterRepoI struct {
	ctrl     *gomock.Controller
	recorder *MockTheaterRepoIMockRecorder
}

// MockTheaterRepoIMockRecorder is the mock recorder for MockTheaterRepoI.
type MockTheaterRepoIMockRecorder struct {
	mock *MockTheaterRepoI
}

// NewMockTheaterRepoI creates a new mock instance.
func NewMockTheaterRepoI(ctrl *gomock.Controller) *MockTheaterRepoI {
	mock := &MockTheaterRepoI{ctrl: ctrl}
	mock.recorder = &MockTheaterRepoIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTheaterRepoI) EXPECT() *MockTheaterRepoIMockRecorder {
	return m.recorder
}

// CreateAddress mocks base method.
func (m *MockTheaterRepoI) CreateAddress(address model.Addresses) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddress", address)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// CreateAddress indicates an expected call of CreateAddress.
func (mr *MockTheaterRepoIMockRecorder) CreateAddress(address any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddress", reflect.TypeOf((*MockTheaterRepoI)(nil).CreateAddress), address)
}

// CreateTheatre mocks base method.
func (m *MockTheaterRepoI) CreateTheatre(theatre model.Theatres) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTheatre", theatre)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// CreateTheatre indicates an expected call of CreateTheatre.
func (mr *MockTheaterRepoIMockRecorder) CreateTheatre(theatre any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTheatre", reflect.TypeOf((*MockTheaterRepoI)(nil).CreateTheatre), theatre)
}

// GetSeatsForCinemaHall mocks base method.
func (m *MockTheaterRepoI) GetSeatsForCinemaHall(cinemaHallId *uuid.UUID) ([]model.Seats, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeatsForCinemaHall", cinemaHallId)
	ret0, _ := ret[0].([]model.Seats)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetSeatsForCinemaHall indicates an expected call of GetSeatsForCinemaHall.
func (mr *MockTheaterRepoIMockRecorder) GetSeatsForCinemaHall(cinemaHallId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeatsForCinemaHall", reflect.TypeOf((*MockTheaterRepoI)(nil).GetSeatsForCinemaHall), cinemaHallId)
}
