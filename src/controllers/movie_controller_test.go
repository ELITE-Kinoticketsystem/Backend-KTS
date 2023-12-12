package controllers

import (
	"testing"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetMovies(t *testing.T) {
	sampleMovies := utils.GetSampleMovies()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockMovieRepositoryI)
		expectedMovies  *[]model.Movies
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI) {
				mockRepo.EXPECT().GetMovies().Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedMovies: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI) {
				mockRepo.EXPECT().GetMovies().Return(sampleMovies, nil)
			},
			expectedMovies: sampleMovies,
			expectedError:  nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI) {
				mockRepo.EXPECT().GetMovies().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedMovies: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			movieRepoMock := mocks.NewMockMovieRepositoryI(mockCtrl)
			movieController := MovieController{
				MovieRepo: movieRepoMock,
			}

			// define expectations
			tc.setExpectations(*movieRepoMock)

			// WHEN
			// call RegisterUser with registrationData
			movies, kts_err := movieController.GetMovies()

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedMovies, movies)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetMovieByName(t *testing.T) {
	sampleMovie := utils.GetSampleMovieById()

	name := sampleMovie.Title

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockMovieRepositoryI, name *string)
		expectedMovies  *model.Movies
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI, name *string) {
				mockRepo.EXPECT().GetMovieByName(name).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedMovies: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI, name *string) {
				mockRepo.EXPECT().GetMovieByName(name).Return(sampleMovie, nil)
			},
			expectedMovies: sampleMovie,
			expectedError:  nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI, name *string) {
				mockRepo.EXPECT().GetMovieByName(name).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedMovies: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			movieRepoMock := mocks.NewMockMovieRepositoryI(mockCtrl)

			movieController := MovieController{
				MovieRepo: movieRepoMock,
			}

			// define expectations
			tc.setExpectations(*movieRepoMock, &name)

			// WHEN
			// call RegisterUser with registrationData
			movies, kts_err := movieController.GetMovieByName(&name)

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedMovies, movies)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetMoviesWithGenres(t *testing.T) {
	sampleMovies := utils.GetSampleMoviesWithGenres()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockMovieRepositoryI)
		expectedMovies  *[]models.MovieWithGenres
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI) {
				mockRepo.EXPECT().GetMoviesWithGenres().Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedMovies: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI) {
				mockRepo.EXPECT().GetMoviesWithGenres().Return(sampleMovies, nil)
			},
			expectedMovies: sampleMovies,
			expectedError:  nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI) {
				mockRepo.EXPECT().GetMoviesWithGenres().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedMovies: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			movieRepoMock := mocks.NewMockMovieRepositoryI(mockCtrl)
			movieController := MovieController{
				MovieRepo: movieRepoMock,
			}

			// define expectations
			tc.setExpectations(*movieRepoMock)

			// WHEN
			// call RegisterUser with registrationData
			movies, kts_err := movieController.GetMoviesWithGenres()

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedMovies, movies)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetMovieById(t *testing.T) {
	sampleMovie := utils.GetSampleMovieByIdWithEverything()

	id := sampleMovie.ID

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockMovieRepositoryI, movieId *uuid.UUID)
		expectedMovies  *models.MovieWithEverything
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI, movieId *uuid.UUID) {
				mockRepo.EXPECT().GetMovieById(id).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedMovies: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI, movieId *uuid.UUID) {
				mockRepo.EXPECT().GetMovieById(id).Return(sampleMovie, nil)
			},
			expectedMovies: sampleMovie,
			expectedError:  nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mockRepo mocks.MockMovieRepositoryI, movieId *uuid.UUID) {
				mockRepo.EXPECT().GetMovieById(id).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedMovies: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			movieRepoMock := mocks.NewMockMovieRepositoryI(mockCtrl)

			movieController := MovieController{
				MovieRepo: movieRepoMock,
			}

			// define expectations
			tc.setExpectations(*movieRepoMock, id)

			// WHEN
			// call RegisterUser with registrationData
			movies, kts_err := movieController.GetMovieById(id)

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedMovies, movies)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}
