package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateMovieHandler(t *testing.T) {
	sampleUpdateMovie := samples.GetSampleMovieById()

	testCases := []struct {
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockMovieControllerI, movieData interface{})
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			body: sampleUpdateMovie,
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieData interface{}) {
				mockController.EXPECT().UpdateMovie(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleUpdateMovie,
		},
		{
			name: "Internal error",
			body: sampleUpdateMovie,
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieData interface{}) {
				mockController.EXPECT().UpdateMovie(gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name: "Movie title contains empty string",
			body: model.Movies{
				ID:           sampleUpdateMovie.ID,
				Title:        "",
				Description:  "Test Description 1",
				BannerPicURL: sampleUpdateMovie.BannerPicURL,
				CoverPicURL:  sampleUpdateMovie.CoverPicURL,
				TrailerURL:   sampleUpdateMovie.TrailerURL,
				Rating:       sampleUpdateMovie.Rating,
				ReleaseDate:  time.Now(),
				TimeInMin:    120,
				Fsk:          18,
			},
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieData interface{}) {

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
			movieController := mocks.NewMockMovieControllerI(mockCtrl)

			jsonData, _ := json.Marshal(tc.body)
			req, _ := http.NewRequest("PUT", "/movies/", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(movieController, tc.body)

			// WHEN
			handlers.UpdateMovie(movieController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestDeleteMovieHandler(t *testing.T) {
	sampleUpdateMovieId := utils.NewUUID()

	testCases := []struct {
		name            string
		setExpectations func(mockController *mocks.MockMovieControllerI, movieId *uuid.UUID)
		expectedStatus  int
		expectedBody    string
	}{
		{
			name: "Success",
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieId *uuid.UUID) {
				mockController.EXPECT().DeleteMovie(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Movie deleted"}`,
		},
		// {
		// 	name:          "Internal error",
		// 	paramTicketId: utils.NewUUID(),
		// 	setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId *uuid.UUID) {
		// 		mockController.EXPECT().ValidateTicket(gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
		// 	},
		// 	expectedStatus: http.StatusInternalServerError,
		// 	expectedBody:   `{"errorMessage":"INTERNAL_ERROR"}`,
		// },
		// {
		// 	name:          "Ticket not found",
		// 	paramTicketId: utils.NewUUID(),
		// 	setExpectations: func(mockController *mocks.MockTicketControllerI, ticketId *uuid.UUID) {
		// 		mockController.EXPECT().ValidateTicket(gomock.Any()).Return(kts_errors.KTS_NOT_FOUND)
		// 	},
		// 	expectedStatus: http.StatusNotFound,
		// 	expectedBody:   `{"errorMessage":"NOT_FOUND"}`,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			w := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(w)

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			movieController := mocks.NewMockMovieControllerI(mockCtrl)

			req, _ := http.NewRequest("Delete", "/movies/"+sampleUpdateMovieId.String(), nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "movieId", Value: sampleUpdateMovieId.String()}}

			tc.setExpectations(movieController, sampleUpdateMovieId)

			// WHEN
			handlers.DeleteMovie(movieController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			assert.Regexp(t, tc.expectedBody, w.Body.String(), "wrong HTTP response body")
		})
	}
}
