package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTheatreHandler(t *testing.T) {
	sampleCreateTheatre := samples.GetSampleTheatreCreate()

	testCases := []struct {
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockTheatreControllerI, theatreData interface{})
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			body: sampleCreateTheatre,
			setExpectations: func(mockController *mocks.MockTheatreControllerI, theatreData interface{}) {
				mockController.EXPECT().CreateTheatre(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   nil,
		},
		{
			name: "Internal error",
			body: sampleCreateTheatre,
			setExpectations: func(mockController *mocks.MockTheatreControllerI, theatreData interface{}) {
				mockController.EXPECT().CreateTheatre(gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name: "Title string empty",
			body: models.CreateTheatreRequest{
				Name:    "",
				Address: sampleCreateTheatre.Address,
				LogoUrl: sampleCreateTheatre.LogoUrl,
			},
			setExpectations: func(mockController *mocks.MockTheatreControllerI, theatreData interface{}) {

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
			theatreController := mocks.NewMockTheatreControllerI(mockCtrl)

			jsonData, _ := json.Marshal(tc.body)
			req, _ := http.NewRequest("POST", "/theatre/", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(theatreController, tc.body)

			// WHEN
			handlers.CreateTheatre(theatreController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestCreateCinemaHall(t *testing.T) {
	sampleCreateCinemaHall := samples.GetSampleCreateCinemaHallRequest()

	testCases := []struct {
		name            string
		body            gin.H
		setExpectations func(mockController *mocks.MockTheatreControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			body: gin.H{
				"hallName":  sampleCreateCinemaHall.HallName,
				"seats":     sampleCreateCinemaHall.Seats,
				"theatreId": sampleCreateCinemaHall.TheatreId,
			},
			setExpectations: func(mockController *mocks.MockTheatreControllerI) {
				mockController.EXPECT().CreateCinemaHall(&sampleCreateCinemaHall).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   nil,
		},
		{
			name: "Internal error",
			body: gin.H{
				"hallName":  sampleCreateCinemaHall.HallName,
				"seats":     sampleCreateCinemaHall.Seats,
				"theatreId": sampleCreateCinemaHall.TheatreId,
			},
			setExpectations: func(mockController *mocks.MockTheatreControllerI) {
				mockController.EXPECT().CreateCinemaHall(&sampleCreateCinemaHall).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name: "Invalid id",
			body: gin.H{
				"HallName":  sampleCreateCinemaHall.HallName,
				"TheatreId": "invalid id",
				"Seats":     sampleCreateCinemaHall.Seats,
			},
			setExpectations: func(mockController *mocks.MockTheatreControllerI) {},
			expectedStatus:  http.StatusBadRequest,
			expectedBody:    gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name: "Empty field",
			body: gin.H{
				"HallName":  "",
				"TheatreId": sampleCreateCinemaHall.TheatreId,
				"Seats":     sampleCreateCinemaHall.Seats,
			},
			setExpectations: func(mockController *mocks.MockTheatreControllerI) {},
			expectedStatus:  http.StatusBadRequest,
			expectedBody:    gin.H{"errorMessage": "BAD_REQUEST"},
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
			theatreController := mocks.NewMockTheatreControllerI(mockCtrl)

			jsonData, _ := json.Marshal(tc.body)
			req, _ := http.NewRequest("POST", "/theatre/hall/", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(theatreController)

			// WHEN
			handlers.CreateCinemaHallHandler(theatreController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetCinemaHallsForTheatre(t *testing.T) {
	theatreId := myid.New()
	sampleCinemaHalls := samples.GetSampleCinemaHalls()

	testCases := []struct {
		name            string
		theatreId       string
		setExpectations func(mockController *mocks.MockTheatreControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:      "Success",
			theatreId: theatreId.String(),
			setExpectations: func(mockController *mocks.MockTheatreControllerI) {
				mockController.EXPECT().GetCinemaHallsForTheatre(&theatreId).Return(&sampleCinemaHalls, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleCinemaHalls,
		},
		{
			name:      "Internal error",
			theatreId: theatreId.String(),
			setExpectations: func(mockController *mocks.MockTheatreControllerI) {
				mockController.EXPECT().GetCinemaHallsForTheatre(&theatreId).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:            "Invalid id",
			theatreId:       "invalid id",
			setExpectations: func(mockController *mocks.MockTheatreControllerI) {},
			expectedStatus:  http.StatusBadRequest,
			expectedBody:    gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:      "Not found",
			theatreId: theatreId.String(),
			setExpectations: func(mockController *mocks.MockTheatreControllerI) {
				mockController.EXPECT().GetCinemaHallsForTheatre(&theatreId).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   gin.H{"errorMessage": "NOT_FOUND"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock context
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			// create mock controller
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			theatreController := mocks.NewMockTheatreControllerI(mockCtrl)

			// create mock request
			req, _ := http.NewRequest("GET", "/theatre/"+tc.theatreId+"/hall/", nil)
			c.Request = req
			c.AddParam("theatreId", tc.theatreId)

			// define expectations
			tc.setExpectations(theatreController)

			// WHEN
			// call GetCinemaHallsForTheatreHandler with mock context
			handlers.GetCinemaHallsForTheatreHandler(theatreController)(c)

			// THEN
			// check HTTP status code and response body
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
