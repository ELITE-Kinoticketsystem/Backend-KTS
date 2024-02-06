package controllers

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/mocks"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetMovies(t *testing.T) {
	sampleMovies := samples.GetSampleMovies()

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
	sampleMovie := samples.GetSampleMovieById()

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
	sampleMovies := samples.GetSampleMoviesWithGenres()

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
	sampleMovie := samples.GetSampleMovieByIdWithEverything()

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

func TestCreateMovie(t *testing.T) {
	sampleMovie := samples.GetSampleMovieDTOCreate()

	testCases := []struct {
		name             string
		movieDTOModel    *models.MovieDTOCreate
		setExpectations  func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, movie *models.MovieDTOCreate, db *sql.DB, dbMock sqlmock.Sqlmock)
		expectedMoviesId bool
		expectedError    *models.KTSError
	}{
		{
			name:          "CreateTransaction internal error",
			movieDTOModel: sampleMovie,
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, movie *models.MovieDTOCreate, db *sql.DB, dbMock sqlmock.Sqlmock) {
				mockMovieRepo.EXPECT().NewTransaction().Return(nil, sql.ErrTxDone)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{ // Done
			name:          "Bad Request",
			movieDTOModel: &models.MovieDTOCreate{},
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, movie *models.MovieDTOCreate, db *sql.DB, dbMock sqlmock.Sqlmock) {

			},
			expectedMoviesId: false,
			expectedError:    kts_errors.KTS_BAD_REQUEST,
		},
		{
			name: "Movie failed",
			movieDTOModel: &models.MovieDTOCreate{
				Movies: sampleMovie.Movies,
			},
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, movie *models.MovieDTOCreate, db *sql.DB, dbMock sqlmock.Sqlmock) {
				dbMock.ExpectBegin()
				tx, _ := db.Begin()
				mockMovieRepo.EXPECT().NewTransaction().Return(tx, nil)
				mockMovieRepo.EXPECT().CreateMovie(tx, &movie.Movies).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedMoviesId: false,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Genre failed",
			movieDTOModel: &models.MovieDTOCreate{
				Movies: sampleMovie.Movies,

				GenresID: sampleMovie.GenresID,
			},
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, movie *models.MovieDTOCreate, db *sql.DB, dbMock sqlmock.Sqlmock) {
				dbMock.ExpectBegin()
				tx, _ := db.Begin()
				mockMovieRepo.EXPECT().NewTransaction().Return(tx, nil)
				mockMovieRepo.EXPECT().CreateMovie(tx, &movie.Movies).Return(sampleMovie.ID, nil)
				mockMovieGenreRepo.EXPECT().AddMovieGenre(tx, sampleMovie.ID, gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedMoviesId: false,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Actor failed",
			movieDTOModel: &models.MovieDTOCreate{
				Movies: sampleMovie.Movies,

				GenresID: sampleMovie.GenresID,
				ActorsID: sampleMovie.ActorsID,
			},
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, movie *models.MovieDTOCreate, db *sql.DB, dbMock sqlmock.Sqlmock) {
				dbMock.ExpectBegin()
				tx, _ := db.Begin()
				mockMovieRepo.EXPECT().NewTransaction().Return(tx, nil)
				mockMovieRepo.EXPECT().CreateMovie(tx, &movie.Movies).Return(sampleMovie.ID, nil)
				mockMovieGenreRepo.EXPECT().AddMovieGenre(tx, sampleMovie.ID, gomock.Any()).Return(nil)
				mockMovieActorRepo.EXPECT().AddMovieActor(tx, sampleMovie.ID, gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedMoviesId: false,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Producer failed",
			movieDTOModel: &models.MovieDTOCreate{
				Movies: sampleMovie.Movies,

				GenresID:    sampleMovie.GenresID,
				ActorsID:    sampleMovie.ActorsID,
				ProducersID: sampleMovie.ProducersID,
			},
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, movie *models.MovieDTOCreate, db *sql.DB, dbMock sqlmock.Sqlmock) {
				dbMock.ExpectBegin()
				tx, _ := db.Begin()
				mockMovieRepo.EXPECT().NewTransaction().Return(tx, nil)
				mockMovieRepo.EXPECT().CreateMovie(tx, &movie.Movies).Return(sampleMovie.ID, nil)
				mockMovieGenreRepo.EXPECT().AddMovieGenre(tx, sampleMovie.ID, gomock.Any()).Return(nil)
				mockMovieActorRepo.EXPECT().AddMovieActor(tx, sampleMovie.ID, gomock.Any()).Return(nil)
			},
			expectedMoviesId: false,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie successfuly created",
			movieDTOModel: &models.MovieDTOCreate{
				Movies: sampleMovie.Movies,

				GenresID:    sampleMovie.GenresID,
				ActorsID:    sampleMovie.ActorsID,
				ProducersID: sampleMovie.ProducersID,
			},
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, movie *models.MovieDTOCreate, db *sql.DB, dbMock sqlmock.Sqlmock) {
				dbMock.ExpectBegin()
				tx, _ := db.Begin()
				mockMovieRepo.EXPECT().NewTransaction().Return(tx, nil)
				mockMovieRepo.EXPECT().CreateMovie(tx, &movie.Movies).Return(sampleMovie.ID, nil)
				mockMovieGenreRepo.EXPECT().AddMovieGenre(tx, sampleMovie.ID, gomock.Any()).Return(nil)
				mockMovieActorRepo.EXPECT().AddMovieActor(tx, sampleMovie.ID, gomock.Any()).Return(nil)
				dbMock.ExpectCommit().WillReturnError(sql.ErrTxDone)
			},
			expectedMoviesId: false,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie successfuly created",
			movieDTOModel: &models.MovieDTOCreate{
				Movies: sampleMovie.Movies,

				GenresID:    sampleMovie.GenresID,
				ActorsID:    sampleMovie.ActorsID,
				ProducersID: sampleMovie.ProducersID,
			},
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, movie *models.MovieDTOCreate, db *sql.DB, dbMock sqlmock.Sqlmock) {
				dbMock.ExpectBegin()
				tx, _ := db.Begin()
				mockMovieRepo.EXPECT().NewTransaction().Return(tx, nil)
				mockMovieRepo.EXPECT().CreateMovie(tx, &movie.Movies).Return(sampleMovie.ID, nil)
				mockMovieGenreRepo.EXPECT().AddMovieGenre(tx, sampleMovie.ID, gomock.Any()).Return(nil)
				mockMovieActorRepo.EXPECT().AddMovieActor(tx, sampleMovie.ID, gomock.Any()).Return(nil)
				dbMock.ExpectCommit()
			},
			expectedMoviesId: true,
			expectedError:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// create mock user repo
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			movieRepoMock := mocks.NewMockMovieRepositoryI(mockCtrl)
			genreRepoMock := mocks.NewMockMovieGenreRepositoryI(mockCtrl)
			actorRepoMock := mocks.NewMockMovieActorRepositoryI(mockCtrl)

			movieController := MovieController{
				MovieRepo:         movieRepoMock,
				MovieGenreRepo:    genreRepoMock,
				MovieActorRepo:    actorRepoMock,
			}

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error while setting up mock database: %s", err)
			}

			// define expectations
			tc.setExpectations(*movieRepoMock, *genreRepoMock, *actorRepoMock, tc.movieDTOModel, db, mock)

			// WHEN
			// call RegisterUser with registrationData
			movieId, kts_err := movieController.CreateMovie(tc.movieDTOModel)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)

			if tc.expectedMoviesId && movieId == nil {
				t.Error("Expected genre ID, got nil")
			}

			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUpdateMovie(t *testing.T) {
	sampleMovie := samples.GetSampleMovieById()

	testCases := []struct {
		name            string
		setExpectations func(mockMovieRepo mocks.MockMovieRepositoryI, movie *model.Movies)
		expectedError   *models.KTSError
	}{
		{
			name: "Update movie",
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, movie *model.Movies) {
				mockMovieRepo.EXPECT().UpdateMovie(movie).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error while updating movie",
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, movie *model.Movies) {
				mockMovieRepo.EXPECT().UpdateMovie(movie).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, movie *model.Movies) {
				mockMovieRepo.EXPECT().UpdateMovie(movie).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie not found",
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, movie *model.Movies) {
				mockMovieRepo.EXPECT().UpdateMovie(movie).Return(kts_errors.KTS_NOT_FOUND)
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			movieRepoMock := mocks.NewMockMovieRepositoryI(mockCtrl)

			movieController := MovieController{
				MovieRepo: movieRepoMock,
			}

			// define expectations
			tc.setExpectations(*movieRepoMock, sampleMovie)

			// WHEN
			// Call the method under test
			kts_err := movieController.UpdateMovie(sampleMovie)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}

func TestDeleteMovie(t *testing.T) {
	sampleMovie := samples.GetSampleMovieById()

	testCases := []struct {
		name            string
		setExpectations func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockReviewRepo mocks.MockReviewRepositoryI, movieId *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Delete movie",
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockReviewRepo mocks.MockReviewRepositoryI, movieId *uuid.UUID) {
				mockMovieGenreRepo.EXPECT().RemoveAllGenreCombinationWithMovie(movieId).Return(nil)
				mockMovieActorRepo.EXPECT().RemoveAllActorCombinationWithMovie(movieId).Return(nil)
				mockReviewRepo.EXPECT().DeleteReviewForMovie(movieId).Return(nil)
				mockMovieRepo.EXPECT().DeleteMovie(movieId).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "Error while deleting MovieActor",
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockReviewRepo mocks.MockReviewRepositoryI, movieId *uuid.UUID) {
				mockMovieGenreRepo.EXPECT().RemoveAllGenreCombinationWithMovie(movieId).Return(nil)
				mockMovieActorRepo.EXPECT().RemoveAllActorCombinationWithMovie(movieId).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while deleting MovieGenre",
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockReviewRepo mocks.MockReviewRepositoryI, movieId *uuid.UUID) {
				mockMovieGenreRepo.EXPECT().RemoveAllGenreCombinationWithMovie(movieId).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie not found",
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockReviewRepo mocks.MockReviewRepositoryI, movieId *uuid.UUID) {
				mockMovieGenreRepo.EXPECT().RemoveAllGenreCombinationWithMovie(movieId).Return(nil)
				mockMovieActorRepo.EXPECT().RemoveAllActorCombinationWithMovie(movieId).Return(nil)
				mockReviewRepo.EXPECT().DeleteReviewForMovie(movieId).Return(nil)
				mockMovieRepo.EXPECT().DeleteMovie(movieId).Return(kts_errors.KTS_NOT_FOUND)
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "MovieReview internal error",
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockReviewRepo mocks.MockReviewRepositoryI, movieId *uuid.UUID) {
				mockMovieGenreRepo.EXPECT().RemoveAllGenreCombinationWithMovie(movieId).Return(nil)
				mockMovieActorRepo.EXPECT().RemoveAllActorCombinationWithMovie(movieId).Return(nil)
				mockReviewRepo.EXPECT().DeleteReviewForMovie(movieId).Return(kts_errors.KTS_INTERNAL_ERROR)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			movieRepoMock := mocks.NewMockMovieRepositoryI(mockCtrl)
			genreRepoMock := mocks.NewMockMovieGenreRepositoryI(mockCtrl)
			actorRepoMock := mocks.NewMockMovieActorRepositoryI(mockCtrl)
			reviewRepoMock := mocks.NewMockReviewRepositoryI(mockCtrl)

			movieController := MovieController{
				MovieRepo:         movieRepoMock,
				MovieGenreRepo:    genreRepoMock,
				MovieActorRepo:    actorRepoMock,
				ReviewRepo:        reviewRepoMock,
			}

			// define expectations
			tc.setExpectations(*movieRepoMock, *genreRepoMock, *actorRepoMock, *reviewRepoMock, sampleMovie.ID)

			// WHEN
			// Call the method under test
			kts_err := movieController.DeleteMovie(sampleMovie.ID)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)
		})
	}
}
