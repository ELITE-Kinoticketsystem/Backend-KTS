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

	model "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	myid "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
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

// CreateCinemaHall mocks base method.
func (m *MockTheaterRepoI) CreateCinemaHall(cinemaHall model.CinemaHalls) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCinemaHall", cinemaHall)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// CreateCinemaHall indicates an expected call of CreateCinemaHall.
func (mr *MockTheaterRepoIMockRecorder) CreateCinemaHall(cinemaHall any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCinemaHall", reflect.TypeOf((*MockTheaterRepoI)(nil).CreateCinemaHall), cinemaHall)
}

// CreateSeat mocks base method.
func (m *MockTheaterRepoI) CreateSeat(seat model.Seats) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSeat", seat)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// CreateSeat indicates an expected call of CreateSeat.
func (mr *MockTheaterRepoIMockRecorder) CreateSeat(seat any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSeat", reflect.TypeOf((*MockTheaterRepoI)(nil).CreateSeat), seat)
}

// CreateSeats mocks base method.
func (m *MockTheaterRepoI) CreateSeats(seat []model.Seats) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSeats", seat)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// CreateSeats indicates an expected call of CreateSeats.
func (mr *MockTheaterRepoIMockRecorder) CreateSeats(seat any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSeats", reflect.TypeOf((*MockTheaterRepoI)(nil).CreateSeats), seat)
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

// GetCinemaHallsForTheatre mocks base method.
func (m *MockTheaterRepoI) GetCinemaHallsForTheatre(theatreId *myid.UUID) (*[]model.CinemaHalls, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCinemaHallsForTheatre", theatreId)
	ret0, _ := ret[0].(*[]model.CinemaHalls)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetCinemaHallsForTheatre indicates an expected call of GetCinemaHallsForTheatre.
func (mr *MockTheaterRepoIMockRecorder) GetCinemaHallsForTheatre(theatreId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCinemaHallsForTheatre", reflect.TypeOf((*MockTheaterRepoI)(nil).GetCinemaHallsForTheatre), theatreId)
}

// GetSeatCategories mocks base method.
func (m *MockTheaterRepoI) GetSeatCategories() ([]model.SeatCategories, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeatCategories")
	ret0, _ := ret[0].([]model.SeatCategories)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetSeatCategories indicates an expected call of GetSeatCategories.
func (mr *MockTheaterRepoIMockRecorder) GetSeatCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeatCategories", reflect.TypeOf((*MockTheaterRepoI)(nil).GetSeatCategories))
}

// GetSeatsForCinemaHall mocks base method.
func (m *MockTheaterRepoI) GetSeatsForCinemaHall(cinemaHallId *myid.UUID) ([]model.Seats, *models.KTSError) {
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

// GetTheatres mocks base method.
func (m *MockTheaterRepoI) GetTheatres() (*[]model.Theatres, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTheatres")
	ret0, _ := ret[0].(*[]model.Theatres)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetTheatres indicates an expected call of GetTheatres.
func (mr *MockTheaterRepoIMockRecorder) GetTheatres() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTheatres", reflect.TypeOf((*MockTheaterRepoI)(nil).GetTheatres))
}
