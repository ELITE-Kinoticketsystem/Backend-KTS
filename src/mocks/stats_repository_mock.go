// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/stats_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/stats_repository.go -destination=./src/mocks/stats_repository_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	gomock "go.uber.org/mock/gomock"
)

// MockStatsRepositoryI is a mock of StatsRepositoryI interface.
type MockStatsRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockStatsRepositoryIMockRecorder
}

// MockStatsRepositoryIMockRecorder is the mock recorder for MockStatsRepositoryI.
type MockStatsRepositoryIMockRecorder struct {
	mock *MockStatsRepositoryI
}

// NewMockStatsRepositoryI creates a new mock instance.
func NewMockStatsRepositoryI(ctrl *gomock.Controller) *MockStatsRepositoryI {
	mock := &MockStatsRepositoryI{ctrl: ctrl}
	mock.recorder = &MockStatsRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatsRepositoryI) EXPECT() *MockStatsRepositoryIMockRecorder {
	return m.recorder
}

// GetOrdersForStats mocks base method.
func (m *MockStatsRepositoryI) GetOrdersForStats() (*[]models.GetOrderDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersForStats")
	ret0, _ := ret[0].(*[]models.GetOrderDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetOrdersForStats indicates an expected call of GetOrdersForStats.
func (mr *MockStatsRepositoryIMockRecorder) GetOrdersForStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersForStats", reflect.TypeOf((*MockStatsRepositoryI)(nil).GetOrdersForStats))
}

// GetTotalVisits mocks base method.
func (m *MockStatsRepositoryI) GetTotalVisits(startTime, endTime time.Time, in string) (*[]models.StatsVisits, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalVisits", startTime, endTime, in)
	ret0, _ := ret[0].(*[]models.StatsVisits)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetTotalVisits indicates an expected call of GetTotalVisits.
func (mr *MockStatsRepositoryIMockRecorder) GetTotalVisits(startTime, endTime, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalVisits", reflect.TypeOf((*MockStatsRepositoryI)(nil).GetTotalVisits), startTime, endTime, in)
}

// GetTotalVisitsForTheatre mocks base method.
func (m *MockStatsRepositoryI) GetTotalVisitsForTheatre(startTime, endTime time.Time, in, theatreName string) (*[]models.StatsVisits, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalVisitsForTheatre", startTime, endTime, in, theatreName)
	ret0, _ := ret[0].(*[]models.StatsVisits)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetTotalVisitsForTheatre indicates an expected call of GetTotalVisitsForTheatre.
func (mr *MockStatsRepositoryIMockRecorder) GetTotalVisitsForTheatre(startTime, endTime, in, theatreName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalVisitsForTheatre", reflect.TypeOf((*MockStatsRepositoryI)(nil).GetTotalVisitsForTheatre), startTime, endTime, in, theatreName)
}
