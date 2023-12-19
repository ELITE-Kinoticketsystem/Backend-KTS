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
		paramEventId         *uuid.UUID
		setExpectations      func(mockOrderController *mocks.MockOrderControllerI, eventId *uuid.UUID, userId *uuid.UUID)
		expectedResponseBody gin.H
		expectedStatus       int
		createOrderDTO       *models.CreateOrderDTO
	}{
		{
			name:         "Success",
			paramEventId: utils.NewUUID(),
			setExpectations: func(mockOrderController *mocks.MockOrderControllerI, eventId *uuid.UUID, userId *uuid.UUID) {
				mockOrderController.EXPECT().CreateOrder(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
					orderId,
					nil,
				)
			},
			expectedResponseBody: gin.H{
				"orderId": orderId,
			},
			expectedStatus: http.StatusOK,
			createOrderDTO: samples.GetOrderDTO(),
		},
		{
			name:         "Bad Request",
			paramEventId: utils.NewUUID(),
			setExpectations: func(mockOrderController *mocks.MockOrderControllerI, eventId *uuid.UUID, userId *uuid.UUID) {
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/orders/"+tc.paramEventId.String(), nil)
			gin.SetMode(gin.TestMode)

			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Params = []gin.Param{{Key: "eventId", Value: tc.paramEventId.String()}}

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
