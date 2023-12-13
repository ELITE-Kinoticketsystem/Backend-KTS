// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/controllers/eventseats_controller.go
//
// Generated by this command:
//
//	mockgen -source=./src/controllers/eventseats_controller.go -destination=./src/mocks/eventseats_controller_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockEventSeatControllerI is a mock of EventSeatControllerI interface.
type MockEventSeatControllerI struct {
	ctrl     *gomock.Controller
	recorder *MockEventSeatControllerIMockRecorder
}

// MockEventSeatControllerIMockRecorder is the mock recorder for MockEventSeatControllerI.
type MockEventSeatControllerIMockRecorder struct {
	mock *MockEventSeatControllerI
}

// NewMockEventSeatControllerI creates a new mock instance.
func NewMockEventSeatControllerI(ctrl *gomock.Controller) *MockEventSeatControllerI {
	mock := &MockEventSeatControllerI{ctrl: ctrl}
	mock.recorder = &MockEventSeatControllerIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventSeatControllerI) EXPECT() *MockEventSeatControllerIMockRecorder {
	return m.recorder
}

// BlockEventSeat mocks base method.
func (m *MockEventSeatControllerI) BlockEventSeat(eventSeatId, userId *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockEventSeat", eventSeatId, userId)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// BlockEventSeat indicates an expected call of BlockEventSeat.
func (mr *MockEventSeatControllerIMockRecorder) BlockEventSeat(eventSeatId, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockEventSeat", reflect.TypeOf((*MockEventSeatControllerI)(nil).BlockEventSeat), eventSeatId, userId)
}

// GetEventSeats mocks base method.
func (m *MockEventSeatControllerI) GetEventSeats(eventId, userId *uuid.UUID) (*[][]models.GetSeatsForSeatSelectorDTO, *[]models.GetSeatsForSeatSelectorDTO, *time.Time, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventSeats", eventId, userId)
	ret0, _ := ret[0].(*[][]models.GetSeatsForSeatSelectorDTO)
	ret1, _ := ret[1].(*[]models.GetSeatsForSeatSelectorDTO)
	ret2, _ := ret[2].(*time.Time)
	ret3, _ := ret[3].(*models.KTSError)
	return ret0, ret1, ret2, ret3
}

// GetEventSeats indicates an expected call of GetEventSeats.
func (mr *MockEventSeatControllerIMockRecorder) GetEventSeats(eventId, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventSeats", reflect.TypeOf((*MockEventSeatControllerI)(nil).GetEventSeats), eventId, userId)
}
