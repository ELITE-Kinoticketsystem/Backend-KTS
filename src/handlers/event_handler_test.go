package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetEventByIdHandler(t *testing.T) {
	eventId := utils.NewUUID()
	event := samples.GetGetSpecialEventsDTO(eventId)

	eventJson, _ := json.Marshal(event)
	tests := []struct {
		name               string
		paramEventId       string
		setExpectations    func(mockEventController *mocks.MockEventControllerI)
		expectedStatus     int
		expectedBodyString string
	}{
		{
			name:         "Success",
			paramEventId: eventId.String(),
			setExpectations: func(mockEventController *mocks.MockEventControllerI) {
				mockEventController.EXPECT().GetEventById(gomock.Any()).Return(
					event,
					nil,
				)
			},
			expectedStatus:     http.StatusOK,
			expectedBodyString: string(eventJson),
		},
		{
			name:         "Error",
			paramEventId: eventId.String(),
			setExpectations: func(mockEventController *mocks.MockEventControllerI) {
				mockEventController.EXPECT().GetEventById(gomock.Any()).Return(
					nil,
					kts_errors.KTS_INTERNAL_ERROR,
				)
			},
			expectedStatus:     http.StatusInternalServerError,
			expectedBodyString: "{\"errorMessage\":\"INTERNAL_ERROR\"}",
		},
		{
			name:         "Error - Bad request",
			paramEventId: "",
			setExpectations: func(mockEventController *mocks.MockEventControllerI) {

			},
			expectedStatus:     http.StatusBadRequest,
			expectedBodyString: "{\"errorMessage\":\"BAD_REQUEST\"}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/events/"+tc.paramEventId, nil)
			gin.SetMode(gin.TestMode)

			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Params = []gin.Param{{Key: "eventId", Value: tc.paramEventId}}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			eventController := mocks.NewMockEventControllerI(mockCtrl)

			tc.setExpectations(eventController)

			// WHEN
			handlers.GetEventByIdHandler(eventController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			assert.Equal(t, tc.expectedBodyString, w.Body.String(), "wrong HTTP response body")
		})
	}
}
func TestGetEventsForMovieHandler(t *testing.T) {
	movieId := utils.NewUUID()
	theatreId := utils.NewUUID()

	events := samples.GetModelEvents()

	tests := []struct {
		name            string
		paramMovieId    string
		paramTheatreId  string
		setExpectations func(mockEventController *mocks.MockEventControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:           "Success",
			paramMovieId:   movieId.String(),
			paramTheatreId: theatreId.String(),
			setExpectations: func(mockEventController *mocks.MockEventControllerI) {
				mockEventController.EXPECT().GetEventsForMovie(gomock.Any(), gomock.Any()).Return(
					events,
					nil,
				)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   events,
		},
		{
			name:           "Error",
			paramMovieId:   movieId.String(),
			paramTheatreId: theatreId.String(),
			setExpectations: func(mockEventController *mocks.MockEventControllerI) {
				mockEventController.EXPECT().GetEventsForMovie(gomock.Any(), gomock.Any()).Return(
					nil,
					kts_errors.KTS_INTERNAL_ERROR,
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:           "Error - Bad request",
			paramMovieId:   movieId.String(),
			paramTheatreId: "",
			setExpectations: func(mockEventController *mocks.MockEventControllerI) {
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name:           "Error - Bad request",
			paramMovieId:   "",
			paramTheatreId: "",
			setExpectations: func(mockEventController *mocks.MockEventControllerI) {
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/movies/"+tc.paramMovieId+"/events/"+tc.paramTheatreId, nil)
			gin.SetMode(gin.TestMode)

			c, _ := gin.CreateTestContext(w)
			c.Request = req
			c.Params = []gin.Param{{Key: "id", Value: tc.paramMovieId}, {Key: "theatreId", Value: tc.paramTheatreId}}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			eventController := mocks.NewMockEventControllerI(mockCtrl)

			tc.setExpectations(eventController)

			// WHEN
			handlers.GetEventsForMovieHandler(eventController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestCreateEventHandler(t *testing.T) {
	sampleCreateEvent := &models.CreateEvtDTO{
		Events: model.Events{
			Title:        "Test Event",
			Start:        time.Now(),
			End:          time.Now().Add(time.Hour),
			Description:  nil,
			EventType:    "Test event type",
			CinemaHallID: utils.NewUUID(),
		},
		Movies: []*uuid.UUID{
			utils.NewUUID(),
		},
		EventSeatCategories: []model.EventSeatCategories{
			{
				SeatCategoryID: utils.NewUUID(),
				Price:          100,
			},
		},
	}

	testCases := []struct {
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockEventControllerI, event interface{})
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			body: sampleCreateEvent,
			setExpectations: func(mockController *mocks.MockEventControllerI, event interface{}) {
				mockController.EXPECT().CreateEvent(gomock.Any()).Return(sampleCreateEvent.ID, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   sampleCreateEvent.ID,
		},
		{
			name: "Internal error",
			body: sampleCreateEvent,
			setExpectations: func(mockController *mocks.MockEventControllerI, event interface{}) {
				mockController.EXPECT().CreateEvent(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name: "Bad request",
			body: models.CreateEvtDTO{
				Events: model.Events{
					Title: "",
				},
			},
			setExpectations: func(mockController *mocks.MockEventControllerI, event interface{}) {

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
			eventController := mocks.NewMockEventControllerI(mockCtrl)

			jsonData, _ := json.Marshal(tc.body)
			req, _ := http.NewRequest("POST", "/events/", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(eventController, tc.body)

			// WHEN
			handlers.CreateEventHandler(eventController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetSpecialEventsHandler(t *testing.T) {
	sampleSpecialEvent := []models.GetSpecialEventsDTO{}

	movieRating := 4.9
	sampleSpecialEvent = append(sampleSpecialEvent, models.GetSpecialEventsDTO{
		Events: model.Events{
			Title:        "Test Event",
			Start:        time.Now(),
			End:          time.Now().Add(time.Hour),
			Description:  nil,
			EventType:    "Test event type",
			CinemaHallID: utils.NewUUID(),
		},
		Movies: []*model.Movies{
			{
				ID:           utils.NewUUID(),
				Title:        "Test movie",
				Description:  "Test description",
				BannerPicURL: utils.GetStringPointer("Test banner pic url"),
				TrailerURL:   utils.GetStringPointer("Test trailer url"),
				CoverPicURL:  utils.GetStringPointer("Test cover pic url"),
				Rating:       &movieRating,
				ReleaseDate:  time.Now(),
				TimeInMin:    120,
				Fsk:          18,
			},
		},
	})

	testCases := []struct {
		name            string
		setExpectations func(mockController *mocks.MockEventControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			setExpectations: func(mockController *mocks.MockEventControllerI) {
				mockController.EXPECT().GetSpecialEvents().Return(&sampleSpecialEvent, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   &sampleSpecialEvent,
		},
		{
			name: "Internal error",
			setExpectations: func(mockController *mocks.MockEventControllerI) {
				mockController.EXPECT().GetSpecialEvents().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
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
			eventController := mocks.NewMockEventControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/events/special", nil)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(eventController)

			// WHEN
			handlers.GetSpecialEventsHandler(eventController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
