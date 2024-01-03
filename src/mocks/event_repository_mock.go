// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/event_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/event_repository.go -destination=./src/mocks/event_repository_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockEventRepo is a mock of EventRepo interface.
type MockEventRepo struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepoMockRecorder
}

// MockEventRepoMockRecorder is the mock recorder for MockEventRepo.
type MockEventRepoMockRecorder struct {
	mock *MockEventRepo
}

// NewMockEventRepo creates a new mock instance.
func NewMockEventRepo(ctrl *gomock.Controller) *MockEventRepo {
	mock := &MockEventRepo{ctrl: ctrl}
	mock.recorder = &MockEventRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventRepo) EXPECT() *MockEventRepoMockRecorder {
	return m.recorder
}

// AddEventMovie mocks base method.
func (m *MockEventRepo) AddEventMovie(eventId, movieId *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEventMovie", eventId, movieId)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// AddEventMovie indicates an expected call of AddEventMovie.
func (mr *MockEventRepoMockRecorder) AddEventMovie(eventId, movieId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventMovie", reflect.TypeOf((*MockEventRepo)(nil).AddEventMovie), eventId, movieId)
}

// CreateEvent mocks base method.
func (m *MockEventRepo) CreateEvent(event *model.Events) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", event)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockEventRepoMockRecorder) CreateEvent(event any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockEventRepo)(nil).CreateEvent), event)
}

// CreateEventSeat mocks base method.
func (m *MockEventRepo) CreateEventSeat(eventSeat *model.EventSeats) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEventSeat", eventSeat)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// CreateEventSeat indicates an expected call of CreateEventSeat.
func (mr *MockEventRepoMockRecorder) CreateEventSeat(eventSeat any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEventSeat", reflect.TypeOf((*MockEventRepo)(nil).CreateEventSeat), eventSeat)
}

// CreateEventSeatCategory mocks base method.
func (m *MockEventRepo) CreateEventSeatCategory(eventSeatCategory *model.EventSeatCategories) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEventSeatCategory", eventSeatCategory)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// CreateEventSeatCategory indicates an expected call of CreateEventSeatCategory.
func (mr *MockEventRepoMockRecorder) CreateEventSeatCategory(eventSeatCategory any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEventSeatCategory", reflect.TypeOf((*MockEventRepo)(nil).CreateEventSeatCategory), eventSeatCategory)
}

// GetEventById mocks base method.
func (m *MockEventRepo) GetEventById(eventId *uuid.UUID) (*models.GetSpecialEventsDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventById", eventId)
	ret0, _ := ret[0].(*models.GetSpecialEventsDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetEventById indicates an expected call of GetEventById.
func (mr *MockEventRepoMockRecorder) GetEventById(eventId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventById", reflect.TypeOf((*MockEventRepo)(nil).GetEventById), eventId)
}

// GetEventsForMovie mocks base method.
func (m *MockEventRepo) GetEventsForMovie(movieId *uuid.UUID) ([]*model.Events, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsForMovie", movieId)
	ret0, _ := ret[0].([]*model.Events)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetEventsForMovie indicates an expected call of GetEventsForMovie.
func (mr *MockEventRepoMockRecorder) GetEventsForMovie(movieId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsForMovie", reflect.TypeOf((*MockEventRepo)(nil).GetEventsForMovie), movieId)
}

// GetSpecialEvents mocks base method.
func (m *MockEventRepo) GetSpecialEvents() (*[]models.GetSpecialEventsDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpecialEvents")
	ret0, _ := ret[0].(*[]models.GetSpecialEventsDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetSpecialEvents indicates an expected call of GetSpecialEvents.
func (mr *MockEventRepoMockRecorder) GetSpecialEvents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpecialEvents", reflect.TypeOf((*MockEventRepo)(nil).GetSpecialEvents))
}
