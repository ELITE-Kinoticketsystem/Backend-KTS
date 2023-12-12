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

func TestGetGenres(t *testing.T) {
	sampleGenre := utils.GetSampleGenres()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI)
		expectedGenre   *[]model.Genres
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI) {
				mockRepo.EXPECT().GetGenres().Return(nil, kts_errors.KTS_MOVIE_NOT_FOUND)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_MOVIE_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI) {
				mockRepo.EXPECT().GetGenres().Return(sampleGenre, nil)
			},
			expectedGenre: sampleGenre,
			expectedError: nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI) {
				mockRepo.EXPECT().GetGenres().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock)

			// WHEN
			// call RegisterUser with registrationData
			movies, kts_err := genreController.GetGenres()

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedGenre, movies)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetGenreByName(t *testing.T) {
	sampleGenre := utils.GetSampleGenre()

	genreName := sampleGenre.GenreName

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI, genreName *string)
		expectedGenre   *model.Genres
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().GetGenreByName(genreName).Return(nil, kts_errors.KTS_MOVIE_NOT_FOUND)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_MOVIE_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().GetGenreByName(genreName).Return(sampleGenre, nil)
			},
			expectedGenre: sampleGenre,
			expectedError: nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().GetGenreByName(genreName).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock, &genreName)

			// WHEN
			// call RegisterUser with registrationData
			movies, kts_err := genreController.GetGenreByName(&genreName)

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedGenre, movies)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetGenreByNameWithMovies(t *testing.T) {
	sampleGenre := utils.GetSampleGenreByNameWithMovies()

	genreName := sampleGenre.GenreName

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI, genreName *string)
		expectedGenre   *models.GenreWithMovies
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().GetGenreByNameWithMovies(genreName).Return(nil, kts_errors.KTS_MOVIE_NOT_FOUND)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_MOVIE_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().GetGenreByNameWithMovies(genreName).Return(sampleGenre, nil)
			},
			expectedGenre: sampleGenre,
			expectedError: nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().GetGenreByNameWithMovies(genreName).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock, &genreName)

			// WHEN
			// call RegisterUser with registrationData
			movies, kts_err := genreController.GetGenreByNameWithMovies(&genreName)

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedGenre, movies)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetGenresWithMovies(t *testing.T) {
	sampleGenre := utils.GetSampleGenresWithMovies()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI)
		expectedGenre   *[]models.GenreWithMovies
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI) {
				mockRepo.EXPECT().GetGenresWithMovies().Return(nil, kts_errors.KTS_MOVIE_NOT_FOUND)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_MOVIE_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI) {
				mockRepo.EXPECT().GetGenresWithMovies().Return(sampleGenre, nil)
			},
			expectedGenre: sampleGenre,
			expectedError: nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI) {
				mockRepo.EXPECT().GetGenresWithMovies().Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock)

			// WHEN
			// call RegisterUser with registrationData
			movies, kts_err := genreController.GetGenresWithMovies()

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedGenre, movies)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestCreateGenre(t *testing.T) {
	sampleGenre := utils.GetSampleGenre()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI, genreName *string)
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().CreateGenre(genreName).Return(kts_errors.KTS_MOVIE_NOT_FOUND)
			},
			expectedError: kts_errors.KTS_MOVIE_NOT_FOUND,
		},
		{
			name: "Create genre",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().CreateGenre(genreName).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error while creating genre",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().CreateGenre(genreName).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock, &sampleGenre.GenreName)

			// WHEN
			// call RegisterUser with registrationData
			kts_err := genreController.CreateGenre(&sampleGenre.GenreName)

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestUpdateGenre(t *testing.T) {
	sampleGenre := utils.GetSampleGenre()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI, genre *model.Genres)
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genre *model.Genres) {
				mockRepo.EXPECT().UpdateGenre(genre).Return(kts_errors.KTS_MOVIE_NOT_FOUND)
			},
			expectedError: kts_errors.KTS_MOVIE_NOT_FOUND,
		},
		{
			name: "Update genre",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genre *model.Genres) {
				mockRepo.EXPECT().UpdateGenre(genre).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error while updating genre",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genre *model.Genres) {
				mockRepo.EXPECT().UpdateGenre(genre).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock, sampleGenre)

			// WHEN
			// call RegisterUser with registrationData
			kts_err := genreController.UpdateGenre(sampleGenre)

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestDeleteGenre(t *testing.T) {
	sampleGenre := utils.GetSampleGenre()

	genreID := sampleGenre.ID

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI, genreID *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreID *uuid.UUID) {
				mockRepo.EXPECT().DeleteGenre(genreID).Return(kts_errors.KTS_MOVIE_NOT_FOUND)
			},
			expectedError: kts_errors.KTS_MOVIE_NOT_FOUND,
		},
		{
			name: "Delete genre",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreID *uuid.UUID) {
				mockRepo.EXPECT().DeleteGenre(genreID).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error while deleting genre",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreID *uuid.UUID) {
				mockRepo.EXPECT().DeleteGenre(genreID).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock, genreID)

			// WHEN
			// call RegisterUser with registrationData
			kts_err := genreController.DeleteGenre(genreID)

			// THEN
			// check expected error and user
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}
