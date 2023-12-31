package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetTicketByIdHandler(t *testing.T) {
	// sampleTicket := utils.GetSampleTicket()

	testCases := []struct {
		name            string
		paramTicketId   *uuid.UUID
		setExpectations func(mockController *mocks.MockTicketControllerI, ticketId *uuid.UUID)
		expectedStatus  int
		expectedBody    string
	}{
		// {
		// 	name:          "Success",
		// 	paramTicketId: utils.NewUUID(),
		// 	setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId *uuid.UUID) {
		// 		mockController.EXPECT().GetTicketById(gomock.Any()).Return(&sampleTicket, nil)
		// 	},
		// 	expectedStatus: http.StatusOK,
		// 	// expectedBody: gin.H{
		// 	// 	"ID":        sampleTicket.ID.String(),
		// 	// 	"Validated": sampleTicket.Validated,
		// 	// 	"Price":     strconv.Itoa(sampleTicket.Price),
		// 	// 	"Seats": {
		// 	// 		"ID":             sampleTicket.Seats.ID,
		// 	// 		"RowNr":          strconv.Itoa(sampleTicket.Seats.RowNr),
		// 	// 		"ColumnNr":       strconv.Itoa(sampleTicket.Seats.ColumnNr),
		// 	// 		"SeatCategoryID": sampleTicket.Seats.SeatCategoryID.String(),
		// 	// 		"CinemaHallID":   sampleTicket.Seats.CinemaHallID.String(),
		// 	// 		"Type":           sampleTicket.Seats.Type,
		// 	// 	},
		// 	// 	"Order": {
		// 	// 		"ID":              sampleTicket.Order.ID.String(),
		// 	// 		"Totalprice":      strconv.Itoa(sampleTicket.Order.Totalprice),
		// 	// 		"IsPaid":          sampleTicket.Order.IsPaid,
		// 	// 		"PaymentMethodID": sampleTicket.Order.PaymentMethodID.String(),
		// 	// 		"UserID":          sampleTicket.Order.UserID.String(),
		// 	// 	},
		// 	// 	"Event": {
		// 	// 		"ID":           sampleTicket.Event.ID.String(),
		// 	// 		"Title":        sampleTicket.Event.Title,
		// 	// 		"Start":        sampleTicket.Event.Start.String(),
		// 	// 		"End":          sampleTicket.Event.End.String(),
		// 	// 		"Description":  sampleTicket.Event.Description,
		// 	// 		"EventType":    sampleTicket.Event.EventType,
		// 	// 		"CinemaHallID": sampleTicket.Event.CinemaHallID.String(),
		// 	// 	},
		// 	// },
		// 	// expectedBody: gin.H{
		// 	// 	"Ticket": models.TicketDTO{},
		// 	//  },
		// },
		{
			name:          "Internal error",
			paramTicketId: utils.NewUUID(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId *uuid.UUID) {
				mockController.EXPECT().GetTicketById(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"errorMessage":"INTERNAL_ERROR"}`,
		},
		{
			name:          "Ticket not found",
			paramTicketId: utils.NewUUID(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId *uuid.UUID) {
				mockController.EXPECT().GetTicketById(gomock.Any()).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"errorMessage":"NOT_FOUND"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/ticket/"+tc.paramTicketId.String(), nil)
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Params = []gin.Param{{Key: "ticketId", Value: tc.paramTicketId.String()}}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			ticketController := mocks.NewMockTicketControllerI(mockCtrl)

			ticketId := tc.paramTicketId

			tc.setExpectations(ticketController, ticketId)

			// WHEN
			handlers.GetTicketByIdHandler(ticketController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			assert.Regexp(t, tc.expectedBody, w.Body.String(), "wrong HTTP response body")
		})
	}
}

func TestValidateTicketHandler(t *testing.T) {

	testCases := []struct {
		name            string
		paramTicketId   *uuid.UUID
		setExpectations func(mockController *mocks.MockTicketControllerI, ticketId *uuid.UUID)
		expectedStatus  int
		expectedBody    string
	}{
		{
			name:          "Success",
			paramTicketId: utils.NewUUID(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId *uuid.UUID) {
				mockController.EXPECT().ValidateTicket(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Ticket successfully validated"}`,
		},
		{
			name:          "Internal error",
			paramTicketId: utils.NewUUID(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId *uuid.UUID) {
				mockController.EXPECT().ValidateTicket(gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"errorMessage":"INTERNAL_ERROR"}`,
		},
		{
			name:          "Ticket not found",
			paramTicketId: utils.NewUUID(),
			setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId *uuid.UUID) {
				mockController.EXPECT().ValidateTicket(gomock.Any()).Return(kts_errors.KTS_NOT_FOUND)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"errorMessage":"NOT_FOUND"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/ticket/"+tc.paramTicketId.String(), nil)
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Params = []gin.Param{{Key: "ticketId", Value: tc.paramTicketId.String()}}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			ticketController := mocks.NewMockTicketControllerI(mockCtrl)

			ticketId := tc.paramTicketId

			tc.setExpectations(ticketController, ticketId)

			// WHEN
			handlers.ValidateTicketHandler(ticketController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			assert.Regexp(t, tc.expectedBody, w.Body.String(), "wrong HTTP response body")
		})
	}
}
