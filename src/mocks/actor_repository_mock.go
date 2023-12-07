// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/actor_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/actor_repository.go -destination=./src/mocks/actor_repository_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockActorRepoI is a mock of ActorRepoI interface.
type MockActorRepoI struct {
	ctrl     *gomock.Controller
	recorder *MockActorRepoIMockRecorder
}

// MockActorRepoIMockRecorder is the mock recorder for MockActorRepoI.
type MockActorRepoIMockRecorder struct {
	mock *MockActorRepoI
}

// NewMockActorRepoI creates a new mock instance.
func NewMockActorRepoI(ctrl *gomock.Controller) *MockActorRepoI {
	mock := &MockActorRepoI{ctrl: ctrl}
	mock.recorder = &MockActorRepoIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActorRepoI) EXPECT() *MockActorRepoIMockRecorder {
	return m.recorder
}

// GetActorById mocks base method.
func (m *MockActorRepoI) GetActorById(actorId *uuid.UUID) (*models.ActorDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorById", actorId)
	ret0, _ := ret[0].(*models.ActorDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetActorById indicates an expected call of GetActorById.
func (mr *MockActorRepoIMockRecorder) GetActorById(actorId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorById", reflect.TypeOf((*MockActorRepoI)(nil).GetActorById), actorId)
}

// GetActors mocks base method.
func (m *MockActorRepoI) GetActors() (*[]models.GetActorsDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActors")
	ret0, _ := ret[0].(*[]models.GetActorsDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetActors indicates an expected call of GetActors.
func (mr *MockActorRepoIMockRecorder) GetActors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActors", reflect.TypeOf((*MockActorRepoI)(nil).GetActors))
}
