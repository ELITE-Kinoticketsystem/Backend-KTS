package controllers

import (
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestOrderController_CreateOrder(t *testing.T) {
	priceCategories := samples.GetPriceCategories()
	eventSeats := samples.GetGetSlectedSeatsDTO()
	order := samples.GetOrder(priceCategories, eventSeats, utils.NewUUID())

	tests := []struct {
		name          string
		expectedFuncs func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI)
		expectedErr   *models.KTSError
		expectOrderId bool
		orderRequest  models.CreateOrderDTO
	}{
		{
			name: "Create Order",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI) {
				mockEventSeatRepo.EXPECT().GetSelectedSeats(gomock.Any(), gomock.Any()).Return(eventSeats, nil)
				mockPriceCategoryRepo.EXPECT().GetPriceCategories().Return(priceCategories, nil)
				mockOrderRepo.EXPECT().CreateOrder(gomock.Any()).Return(utils.NewUUID(), nil)
				mockTicketRepo.EXPECT().CreateTicket(gomock.Any()).Return(utils.NewUUID(), nil).Times(2)
				mockEventSeatRepo.EXPECT().UpdateEventSeat(gomock.Any()).Return(nil).Times(2)
			},
			expectedErr:   nil,
			expectOrderId: true,
			orderRequest:  *order,
		},
		{
			name: "Create Order - GetSelectedSeats Error",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI) {
				mockEventSeatRepo.EXPECT().GetSelectedSeats(gomock.Any(), gomock.Any()).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedErr:   kts_errors.KTS_NOT_FOUND,
			expectOrderId: false,
			orderRequest:  *order,
		},
		{
			name: "Create Order - GetPriceCategories Error",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI) {
				mockEventSeatRepo.EXPECT().GetSelectedSeats(gomock.Any(), gomock.Any()).Return(eventSeats, nil)
				mockPriceCategoryRepo.EXPECT().GetPriceCategories().Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedErr:   kts_errors.KTS_NOT_FOUND,
			expectOrderId: false,
			orderRequest:  *order,
		},
		{
			name: "Create Order - CreateOrder Error",
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI) {
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
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI) {
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
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI) {
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
			expectedFuncs: func(mockOrderRepo *mocks.MockOrderRepoI, mockEventSeatRepo *mocks.MockEventSeatRepoI, mockPriceCategoryRepo *mocks.MockPriceCategoryRepositoryI, mockTicketRepo *mocks.MockTicketRepositoryI) {

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

			tc.expectedFuncs(mockOrderRepo, mockEventSeatRepo, mockPriceCategoryRepo, mockTicketRepo)

			oc := &OrderController{
				OrderRepo:         mockOrderRepo,
				EventSeatRepo:     mockEventSeatRepo,
				PriceCategoryRepo: mockPriceCategoryRepo,
				TicketRepo:        mockTicketRepo,
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
