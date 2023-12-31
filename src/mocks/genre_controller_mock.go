// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/controllers/genre_controller.go
//
// Generated by this command:
//
//	mockgen -source=./src/controllers/genre_controller.go -destination=./src/mocks/genre_controller_mock.go -package=mocks
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

// MockGenreControllerI is a mock of GenreControllerI interface.
type MockGenreControllerI struct {
	ctrl     *gomock.Controller
	recorder *MockGenreControllerIMockRecorder
}

// MockGenreControllerIMockRecorder is the mock recorder for MockGenreControllerI.
type MockGenreControllerIMockRecorder struct {
	mock *MockGenreControllerI
}

// NewMockGenreControllerI creates a new mock instance.
func NewMockGenreControllerI(ctrl *gomock.Controller) *MockGenreControllerI {
	mock := &MockGenreControllerI{ctrl: ctrl}
	mock.recorder = &MockGenreControllerIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenreControllerI) EXPECT() *MockGenreControllerIMockRecorder {
	return m.recorder
}

// CreateGenre mocks base method.
func (m *MockGenreControllerI) CreateGenre(name *string) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGenre", name)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreateGenre indicates an expected call of CreateGenre.
func (mr *MockGenreControllerIMockRecorder) CreateGenre(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGenre", reflect.TypeOf((*MockGenreControllerI)(nil).CreateGenre), name)
}

// DeleteGenre mocks base method.
func (m *MockGenreControllerI) DeleteGenre(genre_id *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGenre", genre_id)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// DeleteGenre indicates an expected call of DeleteGenre.
func (mr *MockGenreControllerIMockRecorder) DeleteGenre(genre_id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGenre", reflect.TypeOf((*MockGenreControllerI)(nil).DeleteGenre), genre_id)
}

// GetGenreByName mocks base method.
func (m *MockGenreControllerI) GetGenreByName(name *string) (*model.Genres, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenreByName", name)
	ret0, _ := ret[0].(*model.Genres)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetGenreByName indicates an expected call of GetGenreByName.
func (mr *MockGenreControllerIMockRecorder) GetGenreByName(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenreByName", reflect.TypeOf((*MockGenreControllerI)(nil).GetGenreByName), name)
}

// GetGenres mocks base method.
func (m *MockGenreControllerI) GetGenres() (*[]model.Genres, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenres")
	ret0, _ := ret[0].(*[]model.Genres)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetGenres indicates an expected call of GetGenres.
func (mr *MockGenreControllerIMockRecorder) GetGenres() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenres", reflect.TypeOf((*MockGenreControllerI)(nil).GetGenres))
}

// GetGenresWithMovies mocks base method.
func (m *MockGenreControllerI) GetGenresWithMovies() (*[]models.GenreWithMovies, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenresWithMovies")
	ret0, _ := ret[0].(*[]models.GenreWithMovies)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetGenresWithMovies indicates an expected call of GetGenresWithMovies.
func (mr *MockGenreControllerIMockRecorder) GetGenresWithMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenresWithMovies", reflect.TypeOf((*MockGenreControllerI)(nil).GetGenresWithMovies))
}

// UpdateGenre mocks base method.
func (m *MockGenreControllerI) UpdateGenre(genre *model.Genres) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGenre", genre)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// UpdateGenre indicates an expected call of UpdateGenre.
func (mr *MockGenreControllerIMockRecorder) UpdateGenre(genre any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGenre", reflect.TypeOf((*MockGenreControllerI)(nil).UpdateGenre), genre)
}
