package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetGenres(t *testing.T) {
	sampleGenres := utils.GetSampleGenres()

	testCases := []struct {
		name                 string
		body                 interface{}
		setExpectations      func(mockController *mocks.MockGenreControllerI)
		expectedResponseBody interface{}
		expectedStatus       int
	}{
		{
			name: "Empty result",
			body: nil,
			setExpectations: func(mockController *mocks.MockGenreControllerI) {
				mockController.EXPECT().GetGenres().Return(nil, kts_errors.KTS_MOVIE_NOT_FOUND)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "MOVIE_NOT_FOUND",
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "Multiple movies",
			body: nil,
			setExpectations: func(mockController *mocks.MockGenreControllerI) {
				mockController.EXPECT().GetGenres().Return(sampleGenres, nil)
			},
			expectedResponseBody: sampleGenres,
			expectedStatus:       http.StatusOK,
		},
		{
			name: "Error while querying movies",
			body: nil,
			setExpectations: func(mockController *mocks.MockGenreControllerI) {
				mockController.EXPECT().GetGenres().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "INTERNAL_ERROR",
			},
			expectedStatus: http.StatusInternalServerError,
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
			genreController := mocks.NewMockGenreControllerI(mockCtrl)

			// create mock request
			jsonData, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("GET", "/genres", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// define expectations
			tc.setExpectations(genreController)

			// WHEN
			// call CheckUsernameHandler with mock context
			GetGenres(genreController)(c)

			// THEN
			// check the HTTP status code

			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			if w.Body.Len() == 0 {
				assert.True(t, tc.expectedResponseBody == nil, "expected empty response body")
			} else {
				expectedResponseBody, _ := json.Marshal(tc.expectedResponseBody)
				assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
			}
		})
	}
}

func TestGetGenreByName(t *testing.T) {
	sampleGenre := utils.GetSampleGenreByName()

	genreName := sampleGenre.GenreName

	testCases := []struct {
		name                 string
		setExpectations      func(mockController *mocks.MockGenreControllerI)
		expectedResponseBody interface{}
		expectedStatus       int
	}{
		// {
		// 	name: "Movie not found",
		// 	body: "",
		// 	setExpectations: func(mockController *mocks.MockGenreControllerI, genreName *string) {
		// 		mockController.EXPECT().GetGenreByName(genreName).Return(nil, kts_errors.KTS_MOVIE_NOT_FOUND)
		// 	},
		// 	expectedResponseBody: gin.H{
		// 		"errorMessage": "MOVIE_NOT_FOUND",
		// 	},
		// 	expectedStatus: http.StatusNotFound,
		// },
		// {
		// 	name: "Multiple movies",
		// 	body: sampleGenre.GenreName,
		// 	setExpectations: func(mockController *mocks.MockGenreControllerI, genreName *string) {
		// 		mockController.EXPECT().GetGenreByName(genreName).Return(sampleGenre, nil)
		// 	},
		// 	expectedResponseBody: sampleGenre,
		// 	expectedStatus:       http.StatusOK,
		// },
		{
			name: "Error while querying movies",
			setExpectations: func(mockController *mocks.MockGenreControllerI) {
				mockController.EXPECT().GetGenreByName(genreName).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedResponseBody: gin.H{
				"errorMessage": "INTERNAL_ERROR",
			},
			expectedStatus: http.StatusInternalServerError,
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
			genreController := mocks.NewMockGenreControllerI(mockCtrl)

			// create mock request
			req := httptest.NewRequest("GET", "/genres/:name", nil)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// define expectations
			tc.setExpectations(genreController)

			// WHEN
			// call GetGenreByName with mock context
			GetGenreByName(genreController)(c)

			// THEN
			// check the HTTP status code
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedResponseBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
