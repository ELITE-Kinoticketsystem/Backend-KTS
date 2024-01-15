// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/review_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/review_repository.go -destination=./src/mocks/review_repository_mock.go -package=mocks
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

// MockReviewRepositoryI is a mock of ReviewRepositoryI interface.
type MockReviewRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockReviewRepositoryIMockRecorder
}

// MockReviewRepositoryIMockRecorder is the mock recorder for MockReviewRepositoryI.
type MockReviewRepositoryIMockRecorder struct {
	mock *MockReviewRepositoryI
}

// NewMockReviewRepositoryI creates a new mock instance.
func NewMockReviewRepositoryI(ctrl *gomock.Controller) *MockReviewRepositoryI {
	mock := &MockReviewRepositoryI{ctrl: ctrl}
	mock.recorder = &MockReviewRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewRepositoryI) EXPECT() *MockReviewRepositoryIMockRecorder {
	return m.recorder
}

// CreateReview mocks base method.
func (m *MockReviewRepositoryI) CreateReview(review model.Reviews) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", review)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockReviewRepositoryIMockRecorder) CreateReview(review any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockReviewRepositoryI)(nil).CreateReview), review)
}

// DeleteReview mocks base method.
func (m *MockReviewRepositoryI) DeleteReview(id *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReview", id)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// DeleteReview indicates an expected call of DeleteReview.
func (mr *MockReviewRepositoryIMockRecorder) DeleteReview(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReview", reflect.TypeOf((*MockReviewRepositoryI)(nil).DeleteReview), id)
}

// DeleteReviewForMovie mocks base method.
func (m *MockReviewRepositoryI) DeleteReviewForMovie(movieId *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReviewForMovie", movieId)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// DeleteReviewForMovie indicates an expected call of DeleteReviewForMovie.
func (mr *MockReviewRepositoryIMockRecorder) DeleteReviewForMovie(movieId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReviewForMovie", reflect.TypeOf((*MockReviewRepositoryI)(nil).DeleteReviewForMovie), movieId)
}

// GetRatingForMovie mocks base method.
func (m *MockReviewRepositoryI) GetRatingForMovie(movieId *uuid.UUID) (*models.NewRating, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRatingForMovie", movieId)
	ret0, _ := ret[0].(*models.NewRating)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetRatingForMovie indicates an expected call of GetRatingForMovie.
func (mr *MockReviewRepositoryIMockRecorder) GetRatingForMovie(movieId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRatingForMovie", reflect.TypeOf((*MockReviewRepositoryI)(nil).GetRatingForMovie), movieId)
}

// GetReviewById mocks base method.
func (m *MockReviewRepositoryI) GetReviewById(id *uuid.UUID) (*model.Reviews, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewById", id)
	ret0, _ := ret[0].(*model.Reviews)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetReviewById indicates an expected call of GetReviewById.
func (mr *MockReviewRepositoryIMockRecorder) GetReviewById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewById", reflect.TypeOf((*MockReviewRepositoryI)(nil).GetReviewById), id)
}
