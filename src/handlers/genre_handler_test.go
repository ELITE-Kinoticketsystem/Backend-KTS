package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/handlers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdateGenreHandler(t *testing.T) {
	sampleUpdateGenre := samples.GetSampleGenre()

	testCases := []struct {
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockGenreControllerI, genreData interface{})
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			body: sampleUpdateGenre,
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreData interface{}) {
				mockController.EXPECT().UpdateGenre(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleUpdateGenre,
		},
		{
			name: "Internal error",
			body: sampleUpdateGenre,
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreData interface{}) {
				mockController.EXPECT().UpdateGenre(gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name: "Genre Name contains empty string",
			body: model.Genres{
				ID:        sampleUpdateGenre.ID,
				GenreName: "",
			},
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreData interface{}) {

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
			genreController := mocks.NewMockGenreControllerI(mockCtrl)

			jsonData, _ := json.Marshal(tc.body)
			req, _ := http.NewRequest("PUT", "/genres/", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(genreController, tc.body)

			// WHEN
			handlers.UpdateGenre(genreController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestDeleteGenreHandler(t *testing.T) {
	sampleUpdateGenreId := utils.NewUUID()

	testCases := []struct {
		name            string
		genreId         string
		setExpectations func(mockController *mocks.MockGenreControllerI, genreId string)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:    "Success",
			genreId: sampleUpdateGenreId.String(),
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreId string) {
				mockController.EXPECT().DeleteGenre(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   nil,
		},
		{
			name:    "Internal error",
			genreId: sampleUpdateGenreId.String(),
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreId string) {
				mockController.EXPECT().DeleteGenre(gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:    "GenreNmae string empty",
			genreId: "",
			setExpectations: func(mockController *mocks.MockGenreControllerI, movieId string) {

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
			genreController := mocks.NewMockGenreControllerI(mockCtrl)

			req, _ := http.NewRequest("Delete", "/genres/"+tc.genreId, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "id", Value: tc.genreId}}

			tc.setExpectations(genreController, tc.genreId)

			// WHEN
			handlers.DeleteGenre(genreController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetGenresHandler(t *testing.T) {
	sampleGenres := samples.GetSampleGenres()

	testCases := []struct {
		name            string
		setExpectations func(mockController *mocks.MockGenreControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			setExpectations: func(mockController *mocks.MockGenreControllerI) {
				mockController.EXPECT().GetGenres().Return(sampleGenres, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleGenres,
		},
		{
			name: "Internal error",
			setExpectations: func(mockController *mocks.MockGenreControllerI) {
				mockController.EXPECT().GetGenres().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
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
			genreController := mocks.NewMockGenreControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/genres/", nil)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(genreController)

			// WHEN
			handlers.GetGenres(genreController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetGenreByNameHandler(t *testing.T) {
	sampleGenre := samples.GetSampleGenre()

	testCases := []struct {
		name            string
		genreName       string
		setExpectations func(mockController *mocks.MockGenreControllerI, genreName string)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:      "Success",
			genreName: sampleGenre.GenreName,
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreName string) {
				mockController.EXPECT().GetGenreByName(gomock.Any()).Return(sampleGenre, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleGenre,
		},
		{
			name:      "Internal error",
			genreName: sampleGenre.GenreName,
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreName string) {
				mockController.EXPECT().GetGenreByName(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:      "Genre Name contains empty string",
			genreName: "",
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreName string) {

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
			genreController := mocks.NewMockGenreControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/genre/"+tc.genreName, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "name", Value: tc.genreName}}

			tc.setExpectations(genreController, tc.genreName)

			// WHEN
			handlers.GetGenreByName(genreController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestCreateGenreHandler(t *testing.T) {
	genreId := samples.GetSampleGenre().ID
	genreName := samples.GetSampleGenre().GenreName

	testCases := []struct {
		name            string
		genreName       string
		setExpectations func(mockController *mocks.MockGenreControllerI, genreName string)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:      "Success",
			genreName: genreName,
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreName string) {
				mockController.EXPECT().CreateGenre(gomock.Any()).Return(genreId, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   genreId,
		},
		{
			name:      "Internal error",
			genreName: genreName,
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreName string) {
				mockController.EXPECT().CreateGenre(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:      "GenreName string empty",
			genreName: "",
			setExpectations: func(mockController *mocks.MockGenreControllerI, genreName string) {

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
			genreController := mocks.NewMockGenreControllerI(mockCtrl)

			req, _ := http.NewRequest("POST", "/genres/"+tc.genreName, nil)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req
			c.Params = []gin.Param{{Key: "name", Value: tc.genreName}}

			tc.setExpectations(genreController, tc.genreName)

			// WHEN
			handlers.CreateGenre(genreController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}


func TestGetGenresWithMoviesHandler(t *testing.T) {
	sampleGenreWithMovie := samples.GetSampleGenresWithMovies()

	testCases := []struct {
		name            string
		setExpectations func(mockController *mocks.MockGenreControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			setExpectations: func(mockController *mocks.MockGenreControllerI) {
				mockController.EXPECT().GetGenresWithMovies().Return(sampleGenreWithMovie, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleGenreWithMovie,
		},
		{
			name: "Internal error",
			setExpectations: func(mockController *mocks.MockGenreControllerI) {
				mockController.EXPECT().GetGenresWithMovies().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
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
			genreController := mocks.NewMockGenreControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/genres/movies/", nil)
			c.Request = req

			tc.setExpectations(genreController)

			// WHEN
			handlers.GetGenresWithMovies(genreController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
