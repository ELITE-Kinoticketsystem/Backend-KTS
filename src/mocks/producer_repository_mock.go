// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/producer_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/producer_repository.go -destination=./src/mocks/producer_repository_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockProducerRepositoryI is a mock of ProducerRepositoryI interface.
type MockProducerRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockProducerRepositoryIMockRecorder
}

// MockProducerRepositoryIMockRecorder is the mock recorder for MockProducerRepositoryI.
type MockProducerRepositoryIMockRecorder struct {
	mock *MockProducerRepositoryI
}

// NewMockProducerRepositoryI creates a new mock instance.
func NewMockProducerRepositoryI(ctrl *gomock.Controller) *MockProducerRepositoryI {
	mock := &MockProducerRepositoryI{ctrl: ctrl}
	mock.recorder = &MockProducerRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducerRepositoryI) EXPECT() *MockProducerRepositoryIMockRecorder {
	return m.recorder
}

// CreateProducer mocks base method.
func (m *MockProducerRepositoryI) CreateProducer(producer *model.Producers) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProducer", producer)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreateProducer indicates an expected call of CreateProducer.
func (mr *MockProducerRepositoryIMockRecorder) CreateProducer(producer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProducer", reflect.TypeOf((*MockProducerRepositoryI)(nil).CreateProducer), producer)
}

// CreateProducerPicture mocks base method.
func (m *MockProducerRepositoryI) CreateProducerPicture(producerPicture *model.ProducerPictures) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProducerPicture", producerPicture)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreateProducerPicture indicates an expected call of CreateProducerPicture.
func (mr *MockProducerRepositoryIMockRecorder) CreateProducerPicture(producerPicture any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProducerPicture", reflect.TypeOf((*MockProducerRepositoryI)(nil).CreateProducerPicture), producerPicture)
}

// DeleteProducer mocks base method.
func (m *MockProducerRepositoryI) DeleteProducer(id *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProducer", id)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// DeleteProducer indicates an expected call of DeleteProducer.
func (mr *MockProducerRepositoryIMockRecorder) DeleteProducer(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProducer", reflect.TypeOf((*MockProducerRepositoryI)(nil).DeleteProducer), id)
}

// GetProducerById mocks base method.
func (m *MockProducerRepositoryI) GetProducerById(id *uuid.UUID) (*models.ProducerDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducerById", id)
	ret0, _ := ret[0].(*models.ProducerDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetProducerById indicates an expected call of GetProducerById.
func (mr *MockProducerRepositoryIMockRecorder) GetProducerById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducerById", reflect.TypeOf((*MockProducerRepositoryI)(nil).GetProducerById), id)
}

// GetProducers mocks base method.
func (m *MockProducerRepositoryI) GetProducers() (*[]models.GetProducersDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducers")
	ret0, _ := ret[0].(*[]models.GetProducersDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetProducers indicates an expected call of GetProducers.
func (mr *MockProducerRepositoryIMockRecorder) GetProducers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducers", reflect.TypeOf((*MockProducerRepositoryI)(nil).GetProducers))
}

// UpdateProducer mocks base method.
func (m *MockProducerRepositoryI) UpdateProducer(producer *model.Producers) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProducer", producer)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// UpdateProducer indicates an expected call of UpdateProducer.
func (mr *MockProducerRepositoryIMockRecorder) UpdateProducer(producer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProducer", reflect.TypeOf((*MockProducerRepositoryI)(nil).UpdateProducer), producer)
}
