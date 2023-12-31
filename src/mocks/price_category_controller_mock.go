// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/controllers/price_category_controller.go
//
// Generated by this command:
//
//	mockgen -source=./src/controllers/price_category_controller.go -destination=./src/mocks/price_category_controller_mock.go -package=mocks
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

// MockPriceCategoryControllerI is a mock of PriceCategoryControllerI interface.
type MockPriceCategoryControllerI struct {
	ctrl     *gomock.Controller
	recorder *MockPriceCategoryControllerIMockRecorder
}

// MockPriceCategoryControllerIMockRecorder is the mock recorder for MockPriceCategoryControllerI.
type MockPriceCategoryControllerIMockRecorder struct {
	mock *MockPriceCategoryControllerI
}

// NewMockPriceCategoryControllerI creates a new mock instance.
func NewMockPriceCategoryControllerI(ctrl *gomock.Controller) *MockPriceCategoryControllerI {
	mock := &MockPriceCategoryControllerI{ctrl: ctrl}
	mock.recorder = &MockPriceCategoryControllerIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPriceCategoryControllerI) EXPECT() *MockPriceCategoryControllerIMockRecorder {
	return m.recorder
}

// CreatePriceCategory mocks base method.
func (m *MockPriceCategoryControllerI) CreatePriceCategory(priceCategory *model.PriceCategories) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePriceCategory", priceCategory)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreatePriceCategory indicates an expected call of CreatePriceCategory.
func (mr *MockPriceCategoryControllerIMockRecorder) CreatePriceCategory(priceCategory any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePriceCategory", reflect.TypeOf((*MockPriceCategoryControllerI)(nil).CreatePriceCategory), priceCategory)
}

// DeletePriceCategory mocks base method.
func (m *MockPriceCategoryControllerI) DeletePriceCategory(id *uuid.UUID) *models.KTSError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePriceCategory", id)
	ret0, _ := ret[0].(*models.KTSError)
	return ret0
}

// DeletePriceCategory indicates an expected call of DeletePriceCategory.
func (mr *MockPriceCategoryControllerIMockRecorder) DeletePriceCategory(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePriceCategory", reflect.TypeOf((*MockPriceCategoryControllerI)(nil).DeletePriceCategory), id)
}

// GetPriceCategories mocks base method.
func (m *MockPriceCategoryControllerI) GetPriceCategories() (*[]model.PriceCategories, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPriceCategories")
	ret0, _ := ret[0].(*[]model.PriceCategories)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetPriceCategories indicates an expected call of GetPriceCategories.
func (mr *MockPriceCategoryControllerIMockRecorder) GetPriceCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPriceCategories", reflect.TypeOf((*MockPriceCategoryControllerI)(nil).GetPriceCategories))
}

// GetPriceCategoryById mocks base method.
func (m *MockPriceCategoryControllerI) GetPriceCategoryById(id *uuid.UUID) (*model.PriceCategories, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPriceCategoryById", id)
	ret0, _ := ret[0].(*model.PriceCategories)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetPriceCategoryById indicates an expected call of GetPriceCategoryById.
func (mr *MockPriceCategoryControllerIMockRecorder) GetPriceCategoryById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPriceCategoryById", reflect.TypeOf((*MockPriceCategoryControllerI)(nil).GetPriceCategoryById), id)
}

// UpdatePriceCategory mocks base method.
func (m *MockPriceCategoryControllerI) UpdatePriceCategory(priceCategory *model.PriceCategories) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePriceCategory", priceCategory)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// UpdatePriceCategory indicates an expected call of UpdatePriceCategory.
func (mr *MockPriceCategoryControllerIMockRecorder) UpdatePriceCategory(priceCategory any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePriceCategory", reflect.TypeOf((*MockPriceCategoryControllerI)(nil).UpdatePriceCategory), priceCategory)
}
