package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetEventSeatsHandler(t *testing.T) {
	testCases := []struct {
		name                 string
		paramId              string
		setExpectations      func(mockController *mocks.MockEventSeatControllerI, eventSeatId *uuid.UUID, userId *uuid.UUID)
		expectedResponseBody gin.H
		expectedStatus       int
	}{
		{
			name:    "Success",
			paramId: "123",
			setExpectations: func(mockController *mocks.MockEventSeatControllerI, eventSeatId *uuid.UUID, userId *uuid.UUID) {
				mockController.EXPECT().GetEventSeats(gomock.Any(), gomock.Any()).Return(

					&[][]models.GetSeatsForSeatSelectorDTO{},
					&[]models.GetSeatsForSeatSelectorDTO{},
					nil,
					nil)
			},
			expectedResponseBody: gin.H{
				"seat_rows":        []models.GetEventSeatsDTO{},
				"currentUserSeats": []models.GetEventSeatsDTO{},
				"blockedUntil":     nil,
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:    "Internal error",
			paramId: "123",
			setExpectations: func(mockController *mocks.MockEventSeatControllerI, eventSeatId *uuid.UUID, userId *uuid.UUID) {
				mockController.EXPECT().GetEventSeats(gomock.Any(), gomock.Any()).Return(
					nil,
					nil,
					nil,
					kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "INTERNAL_ERROR",
			},
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:    "Event seat not found",
			paramId: "123",
			setExpectations: func(mockController *mocks.MockEventSeatControllerI, eventSeatId *uuid.UUID, userId *uuid.UUID) {
				mockController.EXPECT().GetEventSeats(gomock.Any(), gomock.Any()).Return(
					nil,
					nil,
					nil,
					kts_errors.KTS_NOT_FOUND)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "NOT_FOUND",
			},
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/eventseats/"+tc.paramId, nil)
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)
			c.Request = req

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			eventSeatController := mocks.NewMockEventSeatControllerI(mockCtrl)

			userId := utils.NewUUID()
			id := uuid.New()
			c.Params = []gin.Param{{Key: "id", Value: id.String()}}

			ctx := context.WithValue(c.Request.Context(), models.ContextKeyUserID, userId)
			c.Request = c.Request.WithContext(ctx)

			tc.setExpectations(eventSeatController, &id, userId)

			// WHEN
			handlers.GetEventSeatsHandler(eventSeatController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedResponseBody)

			assert.Equal(t, string(expectedResponseBody), w.Body.String(), "wrong HTTP response body")
		})
	}
}
