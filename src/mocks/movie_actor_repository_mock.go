// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/movie_actor_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/movie_actor_repository.go -destination=./src/mocks/movie_actor_repository_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	reflect "reflect"

	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockMovieActorRepositoryI is a mock of MovieActorRepositoryI interface.
type MockMovieActorRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockMovieActorRepositoryIMockRecorder
}

// MockMovieActorRepositoryIMockRecorder is the mock recorder for MockMovieActorRepositoryI.
type MockMovieActorRepositoryIMockRecorder struct {
	mock *MockMovieActorRepositoryI
}

// NewMockMovieActorRepositoryI creates a new mock instance.
func NewMockMovieActorRepositoryI(ctrl *gomock.Controller) *MockMovieActorRepositoryI {
	mock := &MockMovieActorRepositoryI{ctrl: ctrl}
	mock.recorder = &MockMovieActorRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieActorRepositoryI) EXPECT() *MockMovieActorRepositoryIMockRecorder {
	return m.recorder
}

// AddMovieActor mocks base method.
func (m *MockMovieActorRepositoryI) AddMovieActor(tx *sql.Tx, movieId, actorId *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMovieActor", tx, movieId, actorId)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// AddMovieActor indicates an expected call of AddMovieActor.
func (mr *MockMovieActorRepositoryIMockRecorder) AddMovieActor(tx, movieId, actorId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMovieActor", reflect.TypeOf((*MockMovieActorRepositoryI)(nil).AddMovieActor), tx, movieId, actorId)
}

// RemoveAllActorCombinationWithMovie mocks base method.
func (m *MockMovieActorRepositoryI) RemoveAllActorCombinationWithMovie(movieId *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllActorCombinationWithMovie", movieId)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// RemoveAllActorCombinationWithMovie indicates an expected call of RemoveAllActorCombinationWithMovie.
func (mr *MockMovieActorRepositoryIMockRecorder) RemoveAllActorCombinationWithMovie(movieId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllActorCombinationWithMovie", reflect.TypeOf((*MockMovieActorRepositoryI)(nil).RemoveAllActorCombinationWithMovie), movieId)
}

// RemoveMovieActor mocks base method.
func (m *MockMovieActorRepositoryI) RemoveMovieActor(movieId, actorId *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMovieActor", movieId, actorId)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// RemoveMovieActor indicates an expected call of RemoveMovieActor.
func (mr *MockMovieActorRepositoryIMockRecorder) RemoveMovieActor(movieId, actorId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMovieActor", reflect.TypeOf((*MockMovieActorRepositoryI)(nil).RemoveMovieActor), movieId, actorId)
}
