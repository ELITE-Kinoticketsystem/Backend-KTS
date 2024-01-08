// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/repositories/order_repository.go
//
// Generated by this command:
//
//	mockgen -source=./src/repositories/order_repository.go -destination=./src/mocks/order_repository_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	myid "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderRepoI is a mock of OrderRepoI interface.
type MockOrderRepoI struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepoIMockRecorder
}

// MockOrderRepoIMockRecorder is the mock recorder for MockOrderRepoI.
type MockOrderRepoIMockRecorder struct {
	mock *MockOrderRepoI
}

// NewMockOrderRepoI creates a new mock instance.
func NewMockOrderRepoI(ctrl *gomock.Controller) *MockOrderRepoI {
	mock := &MockOrderRepoI{ctrl: ctrl}
	mock.recorder = &MockOrderRepoIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepoI) EXPECT() *MockOrderRepoIMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderRepoI) CreateOrder(order *model.Orders) (*myid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", order)
	ret0, _ := ret[0].(*myid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepoIMockRecorder) CreateOrder(order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepoI)(nil).CreateOrder), order)
}

// GetOrderById mocks base method.
func (m *MockOrderRepoI) GetOrderById(orderId, userId *myid.UUID) (*models.GetOrderDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", orderId, userId)
	ret0, _ := ret[0].(*models.GetOrderDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockOrderRepoIMockRecorder) GetOrderById(orderId, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockOrderRepoI)(nil).GetOrderById), orderId, userId)
}

// GetOrders mocks base method.
func (m *MockOrderRepoI) GetOrders(userId *myid.UUID) (*[]models.GetOrderDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", userId)
	ret0, _ := ret[0].(*[]models.GetOrderDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderRepoIMockRecorder) GetOrders(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderRepoI)(nil).GetOrders), userId)
}
