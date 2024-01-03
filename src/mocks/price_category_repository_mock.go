// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/price_category_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/price_category_repository.go -destination=./src/mocks/price_category_repository_mock.go -package=mocks
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

// MockPriceCategoryRepositoryI is a mock of PriceCategoryRepositoryI interface.
type MockPriceCategoryRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockPriceCategoryRepositoryIMockRecorder
}

// MockPriceCategoryRepositoryIMockRecorder is the mock recorder for MockPriceCategoryRepositoryI.
type MockPriceCategoryRepositoryIMockRecorder struct {
	mock *MockPriceCategoryRepositoryI
}

// NewMockPriceCategoryRepositoryI creates a new mock instance.
func NewMockPriceCategoryRepositoryI(ctrl *gomock.Controller) *MockPriceCategoryRepositoryI {
	mock := &MockPriceCategoryRepositoryI{ctrl: ctrl}
	mock.recorder = &MockPriceCategoryRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPriceCategoryRepositoryI) EXPECT() *MockPriceCategoryRepositoryIMockRecorder {
	return m.recorder
}

// CreatePriceCategory mocks base method.
func (m *MockPriceCategoryRepositoryI) CreatePriceCategory(priceCategory *model.PriceCategories) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePriceCategory", priceCategory)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreatePriceCategory indicates an expected call of CreatePriceCategory.
func (mr *MockPriceCategoryRepositoryIMockRecorder) CreatePriceCategory(priceCategory any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePriceCategory", reflect.TypeOf((*MockPriceCategoryRepositoryI)(nil).CreatePriceCategory), priceCategory)
}

// DeletePriceCategory mocks base method.
func (m *MockPriceCategoryRepositoryI) DeletePriceCategory(id *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePriceCategory", id)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// DeletePriceCategory indicates an expected call of DeletePriceCategory.
func (mr *MockPriceCategoryRepositoryIMockRecorder) DeletePriceCategory(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePriceCategory", reflect.TypeOf((*MockPriceCategoryRepositoryI)(nil).DeletePriceCategory), id)
}

// GetPriceCategories mocks base method.
func (m *MockPriceCategoryRepositoryI) GetPriceCategories() (*[]model.PriceCategories, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPriceCategories")
	ret0, _ := ret[0].(*[]model.PriceCategories)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetPriceCategories indicates an expected call of GetPriceCategories.
func (mr *MockPriceCategoryRepositoryIMockRecorder) GetPriceCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPriceCategories", reflect.TypeOf((*MockPriceCategoryRepositoryI)(nil).GetPriceCategories))
}

// GetPriceCategoryById mocks base method.
func (m *MockPriceCategoryRepositoryI) GetPriceCategoryById(id *uuid.UUID) (*model.PriceCategories, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPriceCategoryById", id)
	ret0, _ := ret[0].(*model.PriceCategories)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetPriceCategoryById indicates an expected call of GetPriceCategoryById.
func (mr *MockPriceCategoryRepositoryIMockRecorder) GetPriceCategoryById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPriceCategoryById", reflect.TypeOf((*MockPriceCategoryRepositoryI)(nil).GetPriceCategoryById), id)
}

// UpdatePriceCategory mocks base method.
func (m *MockPriceCategoryRepositoryI) UpdatePriceCategory(priceCategory *model.PriceCategories) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePriceCategory", priceCategory)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// UpdatePriceCategory indicates an expected call of UpdatePriceCategory.
func (mr *MockPriceCategoryRepositoryIMockRecorder) UpdatePriceCategory(priceCategory any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePriceCategory", reflect.TypeOf((*MockPriceCategoryRepositoryI)(nil).UpdatePriceCategory), priceCategory)
}
