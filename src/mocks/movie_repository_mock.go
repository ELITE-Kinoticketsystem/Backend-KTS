// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/movie_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/movie_repository.go -destination=./src/mocks/movie_repository_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	reflect "reflect"

	model "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockMovieRepositoryI is a mock of MovieRepositoryI interface.
type MockMovieRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockMovieRepositoryIMockRecorder
}

// MockMovieRepositoryIMockRecorder is the mock recorder for MockMovieRepositoryI.
type MockMovieRepositoryIMockRecorder struct {
	mock *MockMovieRepositoryI
}

// NewMockMovieRepositoryI creates a new mock instance.
func NewMockMovieRepositoryI(ctrl *gomock.Controller) *MockMovieRepositoryI {
	mock := &MockMovieRepositoryI{ctrl: ctrl}
	mock.recorder = &MockMovieRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieRepositoryI) EXPECT() *MockMovieRepositoryIMockRecorder {
	return m.recorder
}

// CreateMovie mocks base method.
func (m *MockMovieRepositoryI) CreateMovie(tx *sql.Tx, movie *model.Movies) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMovie", tx, movie)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreateMovie indicates an expected call of CreateMovie.
func (mr *MockMovieRepositoryIMockRecorder) CreateMovie(tx, movie any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMovie", reflect.TypeOf((*MockMovieRepositoryI)(nil).CreateMovie), tx, movie)
}

// DeleteMovie mocks base method.
func (m *MockMovieRepositoryI) DeleteMovie(movieId *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMovie", movieId)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// DeleteMovie indicates an expected call of DeleteMovie.
func (mr *MockMovieRepositoryIMockRecorder) DeleteMovie(movieId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockMovieRepositoryI)(nil).DeleteMovie), movieId)
}

// GetDatabaseConnection mocks base method.
func (m *MockMovieRepositoryI) GetDatabaseConnection() *sql.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabaseConnection")
	ret0, _ := ret[0].(*sql.DB)
	return ret0
}

// GetDatabaseConnection indicates an expected call of GetDatabaseConnection.
func (mr *MockMovieRepositoryIMockRecorder) GetDatabaseConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabaseConnection", reflect.TypeOf((*MockMovieRepositoryI)(nil).GetDatabaseConnection))
}

// GetMovieById mocks base method.
func (m *MockMovieRepositoryI) GetMovieById(movieId *uuid.UUID) (*models.MovieWithEverything, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieById", movieId)
	ret0, _ := ret[0].(*models.MovieWithEverything)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetMovieById indicates an expected call of GetMovieById.
func (mr *MockMovieRepositoryIMockRecorder) GetMovieById(movieId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieById", reflect.TypeOf((*MockMovieRepositoryI)(nil).GetMovieById), movieId)
}

// GetMovieByName mocks base method.
func (m *MockMovieRepositoryI) GetMovieByName(movieName *string) (*model.Movies, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieByName", movieName)
	ret0, _ := ret[0].(*model.Movies)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetMovieByName indicates an expected call of GetMovieByName.
func (mr *MockMovieRepositoryIMockRecorder) GetMovieByName(movieName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieByName", reflect.TypeOf((*MockMovieRepositoryI)(nil).GetMovieByName), movieName)
}

// GetMovies mocks base method.
func (m *MockMovieRepositoryI) GetMovies() (*[]model.Movies, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovies")
	ret0, _ := ret[0].(*[]model.Movies)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetMovies indicates an expected call of GetMovies.
func (mr *MockMovieRepositoryIMockRecorder) GetMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovies", reflect.TypeOf((*MockMovieRepositoryI)(nil).GetMovies))
}

// GetMoviesWithGenres mocks base method.
func (m *MockMovieRepositoryI) GetMoviesWithGenres() (*[]models.MovieWithGenres, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMoviesWithGenres")
	ret0, _ := ret[0].(*[]models.MovieWithGenres)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetMoviesWithGenres indicates an expected call of GetMoviesWithGenres.
func (mr *MockMovieRepositoryIMockRecorder) GetMoviesWithGenres() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMoviesWithGenres", reflect.TypeOf((*MockMovieRepositoryI)(nil).GetMoviesWithGenres))
}

// NewTransaction mocks base method.
func (m *MockMovieRepositoryI) NewTransaction() (*sql.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTransaction")
	ret0, _ := ret[0].(*sql.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTransaction indicates an expected call of NewTransaction.
func (mr *MockMovieRepositoryIMockRecorder) NewTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTransaction", reflect.TypeOf((*MockMovieRepositoryI)(nil).NewTransaction))
}

// UpdateMovie mocks base method.
func (m *MockMovieRepositoryI) UpdateMovie(movie *model.Movies) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovie", movie)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// UpdateMovie indicates an expected call of UpdateMovie.
func (mr *MockMovieRepositoryIMockRecorder) UpdateMovie(movie any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovie", reflect.TypeOf((*MockMovieRepositoryI)(nil).UpdateMovie), movie)
}

// UpdateRating mocks base method.
func (m *MockMovieRepositoryI) UpdateRating(movieId *uuid.UUID, rating float64) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRating", movieId, rating)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// UpdateRating indicates an expected call of UpdateRating.
func (mr *MockMovieRepositoryIMockRecorder) UpdateRating(movieId, rating any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRating", reflect.TypeOf((*MockMovieRepositoryI)(nil).UpdateRating), movieId, rating)
}
