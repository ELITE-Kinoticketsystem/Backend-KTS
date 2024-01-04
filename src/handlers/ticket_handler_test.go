package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetTicketByIdHandler(t *testing.T) {
	sampleTicket := samples.GetSampleTicket()

	testCases := []struct {
		name            string
		paramTicketId   string
		setExpectations func(mockController *mocks.MockTicketControllerI, ticketId string)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:          "ID is empty",
			paramTicketId: "",
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:          "Internal error",
			paramTicketId: utils.NewUUID().String(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId string) {
				mockController.EXPECT().GetTicketById(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:          "Ticket not found",
			paramTicketId: utils.NewUUID().String(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId string) {
				mockController.EXPECT().GetTicketById(gomock.Any()).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   gin.H{"errorMessage": "NOT_FOUND"},
		},
		{
			name:          "Ticket found",
			paramTicketId: utils.NewUUID().String(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId string) {
				mockController.EXPECT().GetTicketById(gomock.Any()).Return(sampleTicket, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleTicket,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			ticketController := mocks.NewMockTicketControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/ticket/"+tc.paramTicketId, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "ticketId", Value: tc.paramTicketId}}

			tc.setExpectations(ticketController, tc.paramTicketId)

			// WHEN
			handlers.GetTicketByIdHandler(ticketController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestValidateTicketHandler(t *testing.T) {

	testCases := []struct {
		name            string
		paramTicketId   string
		setExpectations func(mockController *mocks.MockTicketControllerI, ticketId string)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:          "Success",
			paramTicketId: utils.NewUUID().String(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId string) {
				mockController.EXPECT().ValidateTicket(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   models.PatchValidateTicketResponse{Message: "Ticket successfully validated"},
		},
		{
			name:          "Internal error",
			paramTicketId: utils.NewUUID().String(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId string) {
				mockController.EXPECT().ValidateTicket(gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:          "Ticket not found",
			paramTicketId: utils.NewUUID().String(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId string) {
				mockController.EXPECT().ValidateTicket(gomock.Any()).Return(kts_errors.KTS_NOT_FOUND)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   gin.H{"errorMessage": "NOT_FOUND"},
		},
		{
			name:          "ID is empty",
			paramTicketId: "",
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			ticketController := mocks.NewMockTicketControllerI(mockCtrl)

			req, _ := http.NewRequest("PATCH", "/ticket/"+tc.paramTicketId, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "ticketId", Value: tc.paramTicketId}}

			tc.setExpectations(ticketController, tc.paramTicketId)

			// WHEN
			handlers.ValidateTicketHandler(ticketController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
