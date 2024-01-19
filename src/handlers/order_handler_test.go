package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateOrderHandler(t *testing.T) {
	orderId := utils.NewUUID()
	tests := []struct {
		name                 string
		paramEventId         string
		setExpectations      func(mockOrderController *mocks.MockOrderControllerI, eventId string, userId *uuid.UUID)
		expectedResponseBody gin.H
		expectedStatus       int
		createOrderDTO       *models.CreateOrderDTO
	}{
		{
			name:         "Success",
			paramEventId: utils.NewUUID().String(),
			setExpectations: func(mockOrderController *mocks.MockOrderControllerI, eventId string, userId *uuid.UUID) {
				mockOrderController.EXPECT().CreateOrder(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
					orderId,
					nil,
				)
			},
			expectedResponseBody: gin.H{
				"id": orderId,
			},
			expectedStatus: http.StatusOK,
			createOrderDTO: samples.GetOrderDTO(),
		},
		{
			name:         "Bad Request",
			paramEventId: utils.NewUUID().String(),
			setExpectations: func(mockOrderController *mocks.MockOrderControllerI, eventId string, userId *uuid.UUID) {
				mockOrderController.EXPECT().CreateOrder(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
					nil,
					kts_errors.KTS_INTERNAL_ERROR,
				)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "INTERNAL_ERROR",
			},
			expectedStatus: http.StatusInternalServerError,
			createOrderDTO: samples.GetOrderDTO(),
		},
		{
			name:         "Bad Request",
			paramEventId: "",
			setExpectations: func(mockOrderController *mocks.MockOrderControllerI, eventId string, userId *uuid.UUID) {

			},
			expectedResponseBody: gin.H{
				"errorMessage": "BAD_REQUEST",
			},
			expectedStatus: http.StatusBadRequest,
			createOrderDTO: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/orders/"+tc.paramEventId, nil)
			gin.SetMode(gin.TestMode)

			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Params = []gin.Param{{Key: "eventId", Value: tc.paramEventId}}

			userId := utils.NewUUID()

			ctx := context.WithValue(c.Request.Context(), models.ContextKeyUserID, userId)
			c.Request = c.Request.WithContext(ctx)
			jsonData, err := json.Marshal(tc.createOrderDTO)
			if err != nil {
				log.Fatal(err)
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonData))

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			orderController := mocks.NewMockOrderControllerI(mockCtrl)

			tc.setExpectations(orderController, tc.paramEventId, userId)

			// WHEN
			CreateOrderHandler(orderController, false)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedResponseBody)
			assert.Equal(t, string(expectedResponseBody), w.Body.String(), "wrong HTTP response body")
		})
	}
}

func TestGetOrderByIdHandler(t *testing.T) {
	order := &(*samples.GetGetOrderDto())[0]

	tests := []struct {
		name            string
		paramOrderId    string
		setExpectations func(mockOrderController *mocks.MockOrderControllerI, orderId string)
		expectedStatus  int
		ExpectedBody    interface{}
	}{
		{
			name:         "Success",
			paramOrderId: utils.NewUUID().String(),
			setExpectations: func(mockOrderController *mocks.MockOrderControllerI, orderId string) {
				mockOrderController.EXPECT().GetOrderById(gomock.Any(), gomock.Any()).Return(
					order,
					nil,
				)
			},
			expectedStatus: http.StatusOK,
			ExpectedBody:   order,
		},
		{
			name:         "Bad Request",
			paramOrderId: utils.NewUUID().String(),
			setExpectations: func(mockOrderController *mocks.MockOrderControllerI, orderId string) {
				mockOrderController.EXPECT().GetOrderById(gomock.Any(), gomock.Any()).Return(
					nil,
					kts_errors.KTS_INTERNAL_ERROR,
				)
			},
			expectedStatus: http.StatusInternalServerError,
			ExpectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:         "Bad Request",
			paramOrderId: "",
			setExpectations: func(mockOrderController *mocks.MockOrderControllerI, orderId string) {
				
			},
			expectedStatus: http.StatusBadRequest,
			ExpectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/orders/"+tc.paramOrderId, nil)
			gin.SetMode(gin.TestMode)

			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Params = []gin.Param{{Key: "orderId", Value: tc.paramOrderId}}

			userId := utils.NewUUID()

			ctx := context.WithValue(c.Request.Context(), models.ContextKeyUserID, userId)
			c.Request = c.Request.WithContext(ctx)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			orderController := mocks.NewMockOrderControllerI(mockCtrl)

			tc.setExpectations(orderController, tc.paramOrderId)

			// WHEN
			GetOrderByIdHandler(orderController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.ExpectedBody)
			assert.Equal(t, string(expectedResponseBody), w.Body.String(), "wrong HTTP response body")

		})
	}
}

func TestGetOrders(t *testing.T) {
	orders := samples.GetGetOrderDto()
	ordersJson, _ := json.Marshal(orders)
	tests := []struct {
		name               string
		setExpectations    func(mockOrderController *mocks.MockOrderControllerI)
		expectedStatus     int
		ExpectedBodyString string
	}{
		{
			name: "Success",
			setExpectations: func(mockOrderController *mocks.MockOrderControllerI) {
				mockOrderController.EXPECT().GetOrders(gomock.Any()).Return(
					orders,
					nil,
				)
			},
			expectedStatus:     http.StatusOK,
			ExpectedBodyString: string(ordersJson),
		},
		{
			name: "Bad Request",
			setExpectations: func(mockOrderController *mocks.MockOrderControllerI) {
				mockOrderController.EXPECT().GetOrders(gomock.Any()).Return(
					nil,
					kts_errors.KTS_INTERNAL_ERROR,
				)
			},
			expectedStatus:     http.StatusInternalServerError,
			ExpectedBodyString: "{\"errorMessage\":\"INTERNAL_ERROR\"}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/orders", nil)
			gin.SetMode(gin.TestMode)

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			userId := utils.NewUUID()

			ctx := context.WithValue(c.Request.Context(), models.ContextKeyUserID, userId)
			c.Request = c.Request.WithContext(ctx)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			orderController := mocks.NewMockOrderControllerI(mockCtrl)

			tc.setExpectations(orderController)

			// WHEN
			GetOrdersHandler(orderController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			assert.Equal(t, tc.ExpectedBodyString, w.Body.String(), "wrong HTTP response body")

		})
	}
}
