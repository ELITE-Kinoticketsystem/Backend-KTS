package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetEventByIdHandler(t *testing.T) {
	eventId := myid.NewUUID()
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
			GetEventByIdHandler(eventController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			assert.Equal(t, tc.expectedBodyString, w.Body.String(), "wrong HTTP response body")
		})
	}
}
func TestGetEventsForMovieHandler(t *testing.T) {
	movieId := myid.NewUUID()
	theatreId := myid.NewUUID()

	events := samples.GetModelEvents()

	eventJsons, _ := json.Marshal(events)

	tests := []struct {
		name               string
		paramMovieId       string
		paramTheatreId     string
		setExpectations    func(mockEventController *mocks.MockEventControllerI)
		expectedStatus     int
		expectedBodyString string
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
			expectedStatus:     http.StatusOK,
			expectedBodyString: string(eventJsons),
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
			expectedStatus:     http.StatusInternalServerError,
			expectedBodyString: "{\"errorMessage\":\"INTERNAL_ERROR\"}",
		},
		{
			name:           "Error - Bad request",
			paramMovieId:   "",
			paramTheatreId: "",
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
			GetEventsForMovieHandler(eventController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			assert.Equal(t, tc.expectedBodyString, w.Body.String(), "wrong HTTP response body")
		})
	}
}
