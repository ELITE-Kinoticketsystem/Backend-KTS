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
		movieId         string
		setExpectations func(mockController *mocks.MockMovieControllerI, movieId string)
		expectedStatus  int
		expectedBody    string
	}{
		{
			name:    "Success",
			movieId: sampleUpdateMovieId.String(),
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieId string) {
				mockController.EXPECT().DeleteMovie(gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Movie deleted"}`,
		},
		{
			name:    "Internal error",
			movieId: sampleUpdateMovieId.String(),
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieId string) {
				mockController.EXPECT().DeleteMovie(gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"errorMessage":"INTERNAL_ERROR"}`,
		},
		{
			name:    "Movie not found",
			movieId: "",
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieId string) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"errorMessage":"BAD_REQUEST"}`,
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

			req, _ := http.NewRequest("Delete", "/movies/"+tc.movieId, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "movieId", Value: tc.movieId}}

			tc.setExpectations(movieController, tc.movieId)

			// WHEN
			handlers.DeleteMovie(movieController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			assert.Regexp(t, tc.expectedBody, w.Body.String(), "wrong HTTP response body")
		})
	}
}

func TestGetMoviesHandler(t *testing.T) {
	sampleMovies := samples.GetSampleMovies()

	testCases := []struct {
		name            string
		setExpectations func(mockController *mocks.MockMovieControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			setExpectations: func(mockController *mocks.MockMovieControllerI) {
				mockController.EXPECT().GetMovies().Return(sampleMovies, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleMovies,
		},
		{
			name: "Internal error",
			setExpectations: func(mockController *mocks.MockMovieControllerI) {
				mockController.EXPECT().GetMovies().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
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
			movieController := mocks.NewMockMovieControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/movies/", nil)
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(movieController)

			// WHEN
			handlers.GetMovies(movieController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetMovieByIdHandler(t *testing.T) {
	sampleMovie := samples.GetSampleMovieByIdWithEverything()

	testCases := []struct {
		name            string
		movieId         string
		setExpectations func(mockController *mocks.MockMovieControllerI, movieId string)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:    "Success",
			movieId: sampleMovie.ID.String(),
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieId string) {
				mockController.EXPECT().GetMovieById(gomock.Any()).Return(sampleMovie, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleMovie,
		},
		{
			name:    "Internal error",
			movieId: sampleMovie.ID.String(),
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieId string) {
				mockController.EXPECT().GetMovieById(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name:    "Movie title contains empty string",
			movieId: "",
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieId string) {

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

			req, _ := http.NewRequest("GET", "/movies/"+tc.movieId, nil)
			c.Request = req
			c.Params = []gin.Param{{Key: "id", Value: tc.movieId}}

			tc.setExpectations(movieController, tc.movieId)

			// WHEN
			handlers.GetMovieById(movieController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestGetMoviesWithGenreHandler(t *testing.T) {
	sampleMovieWithGenre := samples.GetSampleMoviesWithGenres()

	testCases := []struct {
		name            string
		setExpectations func(mockController *mocks.MockMovieControllerI)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			setExpectations: func(mockController *mocks.MockMovieControllerI) {
				mockController.EXPECT().GetMoviesWithGenres().Return(sampleMovieWithGenre, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   sampleMovieWithGenre,
		},
		{
			name: "Internal error",
			setExpectations: func(mockController *mocks.MockMovieControllerI) {
				mockController.EXPECT().GetMoviesWithGenres().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
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
			movieController := mocks.NewMockMovieControllerI(mockCtrl)

			req, _ := http.NewRequest("GET", "/movies/genres/", nil)
			c.Request = req

			tc.setExpectations(movieController)

			// WHEN
			handlers.GetMoviesWithGenres(movieController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}

func TestCreateMovieHandler(t *testing.T) {
	sampleCreateMovie := samples.GetSampleMovieDTOCreate()

	testCases := []struct {
		name            string
		body            interface{}
		setExpectations func(mockController *mocks.MockMovieControllerI, movieData interface{})
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name: "Success",
			body: sampleCreateMovie,
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieData interface{}) {
				mockController.EXPECT().CreateMovie(gomock.Any()).Return(sampleCreateMovie.ID, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   sampleCreateMovie.ID,
		},
		{
			name: "Internal error",
			body: sampleCreateMovie,
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieData interface{}) {
				mockController.EXPECT().CreateMovie(gomock.Any()).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   gin.H{"errorMessage": "INTERNAL_ERROR"},
		},
		{
			name: "Title string empty",
			body: models.MovieDTOCreate{
				Movies: model.Movies{
					ID:           sampleCreateMovie.ID,
					Title:        "",
					Description:  sampleCreateMovie.Description,
					BannerPicURL: sampleCreateMovie.BannerPicURL,
					CoverPicURL:  sampleCreateMovie.CoverPicURL,
					TrailerURL:   sampleCreateMovie.TrailerURL,
					Rating:       sampleCreateMovie.Rating,
					ReleaseDate:  sampleCreateMovie.ReleaseDate,
					TimeInMin:    sampleCreateMovie.TimeInMin,
					Fsk:          sampleCreateMovie.Fsk,
				},
				GenresID:    sampleCreateMovie.GenresID,
				ActorsID:    sampleCreateMovie.ActorsID,
			},
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieData interface{}) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name: "Genre Nil",
			body: models.MovieDTOCreate{
				Movies:      sampleCreateMovie.Movies,
				GenresID:    []struct{ ID *uuid.UUID }{{ID: nil}},
				ActorsID:    sampleCreateMovie.ActorsID,
			},
			setExpectations: func(mockController *mocks.MockMovieControllerI, movieData interface{}) {

			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   gin.H{"errorMessage": "BAD_REQUEST"},
		},
		{
			name: "Actor Nil",
			body: models.MovieDTOCreate{
				Movies:      sampleCreateMovie.Movies,
				GenresID:    sampleCreateMovie.GenresID,
				ActorsID:    []struct{ ID *uuid.UUID }{{ID: nil}},
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
			req, _ := http.NewRequest("POST", "/movies/", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			tc.setExpectations(movieController, tc.body)

			// WHEN
			handlers.CreateMovie(movieController)(c)

			// THEN
			assert.Equal(t, tc.expectedStatus, w.Code, "wrong HTTP status code")
			expectedResponseBody, _ := json.Marshal(tc.expectedBody)
			assert.Equal(t, bytes.NewBuffer(expectedResponseBody).String(), w.Body.String(), "wrong response body")
		})
	}
}
