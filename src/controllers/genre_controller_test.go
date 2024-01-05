package controllers

import (
	"testing"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetGenres(t *testing.T) {
	sampleGenre := samples.GetSampleGenres()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI)
		expectedGenre   *[]model.Genres
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI) {
				mockRepo.EXPECT().GetGenres().Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple genres",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI) {
				mockRepo.EXPECT().GetGenres().Return(sampleGenre, nil)
			},
			expectedGenre: sampleGenre,
			expectedError: nil,
		},
		{
			name: "Error while querying genre",
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
			// create mock genre repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock)

			// WHEN
			genres, kts_err := genreController.GetGenres()

			// THEN
			// check expected error and genres
			assert.Equal(t, tc.expectedGenre, genres)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetGenreByName(t *testing.T) {
	sampleGenre := samples.GetSampleGenre()

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
				mockRepo.EXPECT().GetGenreByName(genreName).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "One genre",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().GetGenreByName(genreName).Return(sampleGenre, nil)
			},
			expectedGenre: sampleGenre,
			expectedError: nil,
		},
		{
			name: "Error while querying genre",
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
			// create mock genre repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock, &genreName)

			// WHEN
			genre, kts_err := genreController.GetGenreByName(&genreName)

			// THEN
			assert.Equal(t, tc.expectedGenre, genre)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestGetGenresWithMovies(t *testing.T) {
	sampleGenre := samples.GetSampleGenresWithMovies()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI)
		expectedGenre   *[]models.GenreWithMovies
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI) {
				mockRepo.EXPECT().GetGenresWithMovies().Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple GenresWithMovies",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI) {
				mockRepo.EXPECT().GetGenresWithMovies().Return(sampleGenre, nil)
			},
			expectedGenre: sampleGenre,
			expectedError: nil,
		},
		{
			name: "Error while querying GenresWithMovies",
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
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock)

			// WHEN
			genresWithMovies, kts_err := genreController.GetGenresWithMovies()

			// THEN
			assert.Equal(t, tc.expectedGenre, genresWithMovies)
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestCreateGenre(t *testing.T) {
	sampleGenre := samples.GetSampleGenre()

	testCases := []struct {
		name            string
		genreName       *string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI, genreName *string)
		expectedGenreId bool
		expectedError   *models.KTSError
	}{
		{
			name:      "Empty result",
			genreName: &sampleGenre.GenreName,
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().CreateGenre(genreName).Return(nil, kts_errors.KTS_NOT_FOUND)
			},
			expectedGenreId: false,
			expectedError:   kts_errors.KTS_NOT_FOUND,
		},
		{
			name:      "Create genre",
			genreName: &sampleGenre.GenreName,
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().CreateGenre(genreName).Return(sampleGenre.ID, nil)
			},
			expectedGenreId: true,
			expectedError:   nil,
		},
		{
			name:      "Error while creating genre",
			genreName: &sampleGenre.GenreName,
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {
				mockRepo.EXPECT().CreateGenre(genreName).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedGenreId: false,
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name:      "GenreName == nil",
			genreName: nil,
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genreName *string) {

			},
			expectedGenreId: false,
			expectedError:   kts_errors.KTS_BAD_REQUEST,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock, tc.genreName)

			// WHEN
			genreId, kts_err := genreController.CreateGenre(tc.genreName)

			// THEN
			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			if tc.expectedGenreId && genreId == nil {
				t.Error("Expected genre ID, got nil")
			}

		})
	}
}

func TestUpdateGenre(t *testing.T) {
	sampleGenre := samples.GetSampleGenre()

	testCases := []struct {
		name            string
		setExpectations func(mockRepo mocks.MockGenreRepositoryI, genre *model.Genres)
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mockRepo mocks.MockGenreRepositoryI, genre *model.Genres) {
				mockRepo.EXPECT().UpdateGenre(genre).Return(kts_errors.KTS_NOT_FOUND)
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
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
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo: genreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock, sampleGenre)

			// WHEN
			kts_err := genreController.UpdateGenre(sampleGenre)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestDeleteGenre(t *testing.T) {
	sampleGenre := samples.GetSampleGenre()

	genreID := sampleGenre.ID

	testCases := []struct {
		name            string
		setExpectations func(mockGenreRepo mocks.MockGenreRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, genreID *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "MovieGenre deletion failed",
			setExpectations: func(mockGenreRepo mocks.MockGenreRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, genreID *uuid.UUID) {
				mockMovieGenreRepo.EXPECT().RemoveAllMovieCombinationWithGenre(genreID).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Empty result",
			setExpectations: func(mockGenreRepo mocks.MockGenreRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, genreID *uuid.UUID) {
				mockMovieGenreRepo.EXPECT().RemoveAllMovieCombinationWithGenre(genreID).Return(nil)
				mockGenreRepo.EXPECT().DeleteGenre(genreID).Return(kts_errors.KTS_NOT_FOUND)
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Delete genre",
			setExpectations: func(mockGenreRepo mocks.MockGenreRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, genreID *uuid.UUID) {
				mockMovieGenreRepo.EXPECT().RemoveAllMovieCombinationWithGenre(genreID).Return(nil)
				mockGenreRepo.EXPECT().DeleteGenre(genreID).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error while deleting genre",
			setExpectations: func(mockGenreRepo mocks.MockGenreRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, genreID *uuid.UUID) {
				mockMovieGenreRepo.EXPECT().RemoveAllMovieCombinationWithGenre(genreID).Return(nil)
				mockGenreRepo.EXPECT().DeleteGenre(genreID).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			genreRepoMock := mocks.NewMockGenreRepositoryI(mockCtrl)
			movieGenreRepoMock := mocks.NewMockMovieGenreRepositoryI(mockCtrl)
			genreController := GenreController{
				GenreRepo:      genreRepoMock,
				MovieGenreRepo: movieGenreRepoMock,
			}

			// define expectations
			tc.setExpectations(*genreRepoMock, *movieGenreRepoMock, genreID)

			// WHEN
			kts_err := genreController.DeleteGenre(genreID)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}
