// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/genre_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/genre_repository.go -destination=./src/mocks/genre_repository_mock.go -package=mocks
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

// MockGenreRepositoryI is a mock of GenreRepositoryI interface.
type MockGenreRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockGenreRepositoryIMockRecorder
}

// MockGenreRepositoryIMockRecorder is the mock recorder for MockGenreRepositoryI.
type MockGenreRepositoryIMockRecorder struct {
	mock *MockGenreRepositoryI
}

// NewMockGenreRepositoryI creates a new mock instance.
func NewMockGenreRepositoryI(ctrl *gomock.Controller) *MockGenreRepositoryI {
	mock := &MockGenreRepositoryI{ctrl: ctrl}
	mock.recorder = &MockGenreRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenreRepositoryI) EXPECT() *MockGenreRepositoryIMockRecorder {
	return m.recorder
}

// CreateGenre mocks base method.
func (m *MockGenreRepositoryI) CreateGenre(name *string) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGenre", name)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreateGenre indicates an expected call of CreateGenre.
func (mr *MockGenreRepositoryIMockRecorder) CreateGenre(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGenre", reflect.TypeOf((*MockGenreRepositoryI)(nil).CreateGenre), name)
}

// DeleteGenre mocks base method.
func (m *MockGenreRepositoryI) DeleteGenre(genreId *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGenre", genreId)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// DeleteGenre indicates an expected call of DeleteGenre.
func (mr *MockGenreRepositoryIMockRecorder) DeleteGenre(genreId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGenre", reflect.TypeOf((*MockGenreRepositoryI)(nil).DeleteGenre), genreId)
}

// GetGenreByName mocks base method.
func (m *MockGenreRepositoryI) GetGenreByName(name *string) (*model.Genres, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenreByName", name)
	ret0, _ := ret[0].(*model.Genres)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetGenreByName indicates an expected call of GetGenreByName.
func (mr *MockGenreRepositoryIMockRecorder) GetGenreByName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenreByName", reflect.TypeOf((*MockGenreRepositoryI)(nil).GetGenreByName), name)
}

// GetGenres mocks base method.
func (m *MockGenreRepositoryI) GetGenres() (*[]model.Genres, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenres")
	ret0, _ := ret[0].(*[]model.Genres)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetGenres indicates an expected call of GetGenres.
func (mr *MockGenreRepositoryIMockRecorder) GetGenres() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenres", reflect.TypeOf((*MockGenreRepositoryI)(nil).GetGenres))
}

// GetGenresWithMovies mocks base method.
func (m *MockGenreRepositoryI) GetGenresWithMovies() (*[]models.GenreWithMovies, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenresWithMovies")
	ret0, _ := ret[0].(*[]models.GenreWithMovies)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetGenresWithMovies indicates an expected call of GetGenresWithMovies.
func (mr *MockGenreRepositoryIMockRecorder) GetGenresWithMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenresWithMovies", reflect.TypeOf((*MockGenreRepositoryI)(nil).GetGenresWithMovies))
}

// UpdateGenre mocks base method.
func (m *MockGenreRepositoryI) UpdateGenre(genre *model.Genres) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGenre", genre)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// UpdateGenre indicates an expected call of UpdateGenre.
func (mr *MockGenreRepositoryIMockRecorder) UpdateGenre(genre any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGenre", reflect.TypeOf((*MockGenreRepositoryI)(nil).UpdateGenre), genre)
}
