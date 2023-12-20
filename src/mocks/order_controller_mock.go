// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/controllers/order_controller.go
//
// Generated by this command:
//
//	mockgen -source=./src/controllers/order_controller.go -destination=./src/mocks/order_controller_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderControllerI is a mock of OrderControllerI interface.
type MockOrderControllerI struct {
	ctrl     *gomock.Controller
	recorder *MockOrderControllerIMockRecorder
}

// MockOrderControllerIMockRecorder is the mock recorder for MockOrderControllerI.
type MockOrderControllerIMockRecorder struct {
	mock *MockOrderControllerI
}

// NewMockOrderControllerI creates a new mock instance.
func NewMockOrderControllerI(ctrl *gomock.Controller) *MockOrderControllerI {
	mock := &MockOrderControllerI{ctrl: ctrl}
	mock.recorder = &MockOrderControllerIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderControllerI) EXPECT() *MockOrderControllerIMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderControllerI) CreateOrder(CreateOrderDTO models.CreateOrderDTO, eventId, userId *uuid.UUID, isReservation bool) (*uuid.UUID, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", CreateOrderDTO, eventId, userId, isReservation)
	ret0, _ := ret[0].(*uuid.UUID)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderControllerIMockRecorder) CreateOrder(CreateOrderDTO, eventId, userId, isReservation any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderControllerI)(nil).CreateOrder), CreateOrderDTO, eventId, userId, isReservation)
}

// GetOrderById mocks base method.
func (m *MockOrderControllerI) GetOrderById(orderId, userId *uuid.UUID) (*models.GetOrderDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", orderId, userId)
	ret0, _ := ret[0].(*models.GetOrderDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockOrderControllerIMockRecorder) GetOrderById(orderId, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockOrderControllerI)(nil).GetOrderById), orderId, userId)
}

// GetOrders mocks base method.
func (m *MockOrderControllerI) GetOrders(userId *uuid.UUID) (*[]models.GetOrderDTO, *models.KTSError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", userId)
	ret0, _ := ret[0].(*[]models.GetOrderDTO)
	ret1, _ := ret[1].(*models.KTSError)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderControllerIMockRecorder) GetOrders(userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderControllerI)(nil).GetOrders), userId)
}
