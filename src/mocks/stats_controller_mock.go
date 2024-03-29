// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/controllers/stats_controller.go
//
// Generated by this command:
//
//	mockgen -source=./src/controllers/stats_controller.go -destination=./src/mocks/stats_controller_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	gomock "go.uber.org/mock/gomock"
)

// MockStatsControllerI is a mock of StatsControllerI interface.
type MockStatsControllerI struct {
	ctrl     *gomock.Controller
	recorder *MockStatsControllerIMockRecorder
}

// MockStatsControllerIMockRecorder is the mock recorder for MockStatsControllerI.
type MockStatsControllerIMockRecorder struct {
	mock *MockStatsControllerI
}

// NewMockStatsControllerI creates a new mock instance.
func NewMockStatsControllerI(ctrl *gomock.Controller) *MockStatsControllerI {
	mock := &MockStatsControllerI{ctrl: ctrl}
	mock.recorder = &MockStatsControllerIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatsControllerI) EXPECT() *MockStatsControllerIMockRecorder {
	return m.recorder
}

// GetMoviesSortedByTicketAmount mocks base method.
func (m *MockStatsControllerI) GetMoviesSortedByTicketAmount() (*[]models.GetEventWithTicketCount, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMoviesSortedByTicketAmount")
	ret0, _ := ret[0].(*[]models.GetEventWithTicketCount)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetMoviesSortedByTicketAmount indicates an expected call of GetMoviesSortedByTicketAmount.
func (mr *MockStatsControllerIMockRecorder) GetMoviesSortedByTicketAmount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMoviesSortedByTicketAmount", reflect.TypeOf((*MockStatsControllerI)(nil).GetMoviesSortedByTicketAmount))
}

// GetOrdersForStats mocks base method.
func (m *MockStatsControllerI) GetOrdersForStats() (*[]models.GetOrderDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersForStats")
	ret0, _ := ret[0].(*[]models.GetOrderDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetOrdersForStats indicates an expected call of GetOrdersForStats.
func (mr *MockStatsControllerIMockRecorder) GetOrdersForStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersForStats", reflect.TypeOf((*MockStatsControllerI)(nil).GetOrdersForStats))
}

// GetTotalVisits mocks base method.
func (m *MockStatsControllerI) GetTotalVisits(startTime, endTime time.Time, in string) (*models.StatsVisitsTwoArrays, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalVisits", startTime, endTime, in)
	ret0, _ := ret[0].(*models.StatsVisitsTwoArrays)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetTotalVisits indicates an expected call of GetTotalVisits.
func (mr *MockStatsControllerIMockRecorder) GetTotalVisits(startTime, endTime, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalVisits", reflect.TypeOf((*MockStatsControllerI)(nil).GetTotalVisits), startTime, endTime, in)
}

// GetTotalVisitsForTheatre mocks base method.
func (m *MockStatsControllerI) GetTotalVisitsForTheatre(startTime, endTime time.Time, in, theatreName string) (*models.StatsVisitsTwoArrays, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalVisitsForTheatre", startTime, endTime, in, theatreName)
	ret0, _ := ret[0].(*models.StatsVisitsTwoArrays)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetTotalVisitsForTheatre indicates an expected call of GetTotalVisitsForTheatre.
func (mr *MockStatsControllerIMockRecorder) GetTotalVisitsForTheatre(startTime, endTime, in, theatreName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalVisitsForTheatre", reflect.TypeOf((*MockStatsControllerI)(nil).GetTotalVisitsForTheatre), startTime, endTime, in, theatreName)
}
