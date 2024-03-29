package controllers

import (
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestOrderController_CreateOrder(t *testing.T) {
	priceCategories := samples.GetPriceCategories()
	eventSeats := samples.GetGetSlectedSeatsDTO()
	order := samples.GetOrder(priceCategories, eventSeats, utils.NewUUID())
	user := samples.GetSampleUser()
	getOrderByIdOrder := samples.GetOrderSample()

	tests := []struct {
		name          string
		expectedFuncs func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI, mockUserRepo *mocks.MockUserRepositoryI, mockMailMgr *mocks.MockMailMgr)
		expectedErr   *models.KTSError
		expectOrderId bool
		orderRequest  models.CreateOrderDTO
	}{
		{
			name: "Create Order",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI, mockUserRepo *mocks.MockUserRepositoryI, mockMailMgr *mocks.MockMailMgr) {
				mockEventSeatRepo.EXPECT().GetSelectedSeats(gomock.Any(), gomock.Any()).Return(eventSeats, nil)
				mockPriceCategoryRepo.EXPECT().GetPriceCategories().Return(priceCategories, nil)
				mockOrderRepo.EXPECT().CreateOrder(gomock.Any()).Return(utils.NewUUID(), nil)
				mockTicketRepo.EXPECT().CreateTicket(gomock.Any()).Return(utils.NewUUID(), nil).Times(2)
				mockEventSeatRepo.EXPECT().UpdateEventSeat(gomock.Any()).Return(nil).Times(2)

				mockOrderRepo.EXPECT().GetOrderById(gomock.Any(), gomock.Any()).Return(&getOrderByIdOrder, nil)

				mockUserRepo.EXPECT().GetUserById(gomock.Any()).Return(&user, nil)

				mockMailMgr.EXPECT().SendOrderConfirmationMail(user.Email, gomock.Any()).Return(nil)

			},
			expectedErr:   nil,
			expectOrderId: true,
			orderRequest:  *order,
		},
		{
			name: "Create Order - GetSelectedSeats Error",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI, mockUserRepo *mocks.MockUserRepositoryI, mockMailMgr *mocks.MockMailMgr) {
				mockEventSeatRepo.EXPECT().GetSelectedSeats(gomock.Any(), gomock.Any()).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedErr:   kts_errors.KTS_NOT_FOUND,
			expectOrderId: false,
			orderRequest:  *order,
		},
		{
			name: "Create Order - GetPriceCategories Error",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI, mockUserRepo *mocks.MockUserRepositoryI, mockMailMgr *mocks.MockMailMgr) {
				mockEventSeatRepo.EXPECT().GetSelectedSeats(gomock.Any(), gomock.Any()).Return(eventSeats, nil)
				mockPriceCategoryRepo.EXPECT().GetPriceCategories().Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedErr:   kts_errors.KTS_NOT_FOUND,
			expectOrderId: false,
			orderRequest:  *order,
		},
		{
			name: "Create Order - CreateOrder Error",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI, mockUserRepo *mocks.MockUserRepositoryI, mockMailMgr *mocks.MockMailMgr) {
				mockEventSeatRepo.EXPECT().GetSelectedSeats(gomock.Any(), gomock.Any()).Return(eventSeats, nil)
				mockPriceCategoryRepo.EXPECT().GetPriceCategories().Return(priceCategories, nil)
				mockOrderRepo.EXPECT().CreateOrder(gomock.Any()).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedErr:   kts_errors.KTS_NOT_FOUND,
			expectOrderId: false,
			orderRequest:  *order,
		},
		{
			name: "Create Order - CreateTicket Error",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI, mockUserRepo *mocks.MockUserRepositoryI, mockMailMgr *mocks.MockMailMgr) {
				mockEventSeatRepo.EXPECT().GetSelectedSeats(gomock.Any(), gomock.Any()).Return(eventSeats, nil)
				mockPriceCategoryRepo.EXPECT().GetPriceCategories().Return(priceCategories, nil)
				mockOrderRepo.EXPECT().CreateOrder(gomock.Any()).Return(utils.NewUUID(), nil)
				mockTicketRepo.EXPECT().CreateTicket(gomock.Any()).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedErr:   kts_errors.KTS_NOT_FOUND,
			expectOrderId: false,
			orderRequest:  *order,
		},
		{
			name: "Create Order - UpdateEventSeat Error",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI, mockUserRepo *mocks.MockUserRepositoryI, mockMailMgr *mocks.MockMailMgr) {
				mockEventSeatRepo.EXPECT().GetSelectedSeats(gomock.Any(), gomock.Any()).Return(eventSeats, nil)
				mockPriceCategoryRepo.EXPECT().GetPriceCategories().Return(priceCategories, nil)
				mockOrderRepo.EXPECT().CreateOrder(gomock.Any()).Return(utils.NewUUID(), nil)
				mockTicketRepo.EXPECT().CreateTicket(gomock.Any()).Return(utils.NewUUID(), nil).Times(2)
				mockEventSeatRepo.EXPECT().UpdateEventSeat(gomock.Any()).Return(kts_errors.KTS_NOT_FOUND)
			},
			expectedErr:   kts_errors.KTS_NOT_FOUND,
			expectOrderId: false,
			orderRequest:  *order,
		},
		{
			name: "Create Order - Empty OrderRequest",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI, mockUserRepo *mocks.MockUserRepositoryI, mockMailMgr *mocks.MockMailMgr) {

			},
			expectedErr:   kts_errors.KTS_BAD_REQUEST,
			expectOrderId: false,
			orderRequest:  models.CreateOrderDTO{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			mockOrderRepo := mocks.NewMockOrderRepoI(mockCtrl)
			mockEventSeatRepo := mocks.NewMockEventSeatRepoI(mockCtrl)
			mockPriceCategoryRepo := mocks.NewMockPriceCategoryRepositoryI(mockCtrl)
			mockTicketRepo := mocks.NewMockTicketRepositoryI(mockCtrl)
			mockUserRepo := mocks.NewMockUserRepositoryI(mockCtrl)

			mockMailMgr := mocks.NewMockMailMgr(mockCtrl)

			tc.expectedFuncs(mockOrderRepo, mockEventSeatRepo, mockPriceCategoryRepo, mockTicketRepo, mockUserRepo, mockMailMgr)

			oc := &OrderController{
				OrderRepo:         mockOrderRepo,
				EventSeatRepo:     mockEventSeatRepo,
				PriceCategoryRepo: mockPriceCategoryRepo,
				TicketRepo:        mockTicketRepo,
				UserRepo:          mockUserRepo,
				MailMgr:           mockMailMgr,
			}

			orderId, err := oc.CreateOrder(tc.orderRequest, utils.NewUUID(), utils.NewUUID(), false)

			if tc.expectOrderId {
				assert.NotNil(t, orderId)
			} else {
				assert.Nil(t, orderId)
			}
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestGetOrderById(t *testing.T) {
	testCases := []struct {
		name            string
		orderId         *uuid.UUID
		userId          *uuid.UUID
		setExpectations func(mockRepo mocks.MockOrderRepoI, orderId *uuid.UUID, userId *uuid.UUID)
		expectedOrder   *models.GetOrderDTO
		expectedError   *models.KTSError
	}{
		{
			name:    "Failed",
			orderId: utils.NewUUID(),
			userId:  utils.NewUUID(),
			setExpectations: func(mockRepo mocks.MockOrderRepoI, orderId *uuid.UUID, userId *uuid.UUID) {
				mockRepo.EXPECT().GetOrderById(orderId, userId).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedOrder: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			orderRepo := mocks.NewMockOrderRepoI(mockCtrl)
			orderController := OrderController{
				OrderRepo: orderRepo,
			}

			// define expectations
			tc.setExpectations(*orderRepo, tc.orderId, tc.userId)

			// WHEN
			order, kts_err := orderController.GetOrderById(tc.orderId, tc.userId)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectedOrder, order)
		})
	}
}

func TestGetOrders(t *testing.T) {
	testCases := []struct {
		name            string
		userId          *uuid.UUID
		setExpectations func(mockRepo mocks.MockOrderRepoI, userId *uuid.UUID)
		expectedOrders  *[]models.GetOrderDTO
		expectedError   *models.KTSError
	}{
		{
			name:   "Failed",
			userId: utils.NewUUID(),
			setExpectations: func(mockRepo mocks.MockOrderRepoI, userId *uuid.UUID) {
				mockRepo.EXPECT().GetOrders(userId).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedOrders: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			orderRepo := mocks.NewMockOrderRepoI(mockCtrl)
			orderController := OrderController{
				OrderRepo: orderRepo,
			}

			// define expectations
			tc.setExpectations(*orderRepo, tc.userId)

			// WHEN
			order, kts_err := orderController.GetOrders(tc.userId)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectedOrders, order)
		})
	}
}
