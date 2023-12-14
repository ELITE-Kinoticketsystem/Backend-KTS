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

func TestCreateMovie(t *testing.T) {
	sampleMovie := utils.GetSampleMovieDTOCreate()

	testCases := []struct {
		name             string
		movieDTOModel    *models.MovieDTOCreate
		setExpectations  func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockMovieProducerRepo mocks.MockMovieProducerRepositoryI, movie *models.MovieDTOCreate)
		expectedMoviesId bool
		expectedError    *models.KTSError
	}{
		{ // Done
			name:          "Bad Request",
			movieDTOModel: &models.MovieDTOCreate{},
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockMovieProducerRepo mocks.MockMovieProducerRepositoryI, movie *models.MovieDTOCreate) {

			},
			expectedMoviesId: false,
			expectedError:    kts_errors.KTS_BAD_REQUEST,
		},
		{
			name: "Movie failed",
			movieDTOModel: &models.MovieDTOCreate{
				Movies: sampleMovie.Movies,
			},
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockMovieProducerRepo mocks.MockMovieProducerRepositoryI, movie *models.MovieDTOCreate) {
				mockMovieRepo.EXPECT().CreateMovie(&movie.Movies).Return(nil, kts_errors.KTS_INTERNAL_ERROR)
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
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockMovieProducerRepo mocks.MockMovieProducerRepositoryI, movie *models.MovieDTOCreate) {
				mockMovieRepo.EXPECT().CreateMovie(&movie.Movies).Return(sampleMovie.ID, nil)
				mockMovieGenreRepo.EXPECT().AddMovieGenre(sampleMovie.ID, gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
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
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockMovieProducerRepo mocks.MockMovieProducerRepositoryI, movie *models.MovieDTOCreate) {
				mockMovieRepo.EXPECT().CreateMovie(&movie.Movies).Return(sampleMovie.ID, nil)
				mockMovieGenreRepo.EXPECT().AddMovieGenre(sampleMovie.ID, gomock.Any()).Return(nil)
				mockMovieActorRepo.EXPECT().AddMovieActor(sampleMovie.ID, gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
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
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockMovieProducerRepo mocks.MockMovieProducerRepositoryI, movie *models.MovieDTOCreate) {
				mockMovieRepo.EXPECT().CreateMovie(&movie.Movies).Return(sampleMovie.ID, nil)
				mockMovieGenreRepo.EXPECT().AddMovieGenre(sampleMovie.ID, gomock.Any()).Return(nil)
				mockMovieActorRepo.EXPECT().AddMovieActor(sampleMovie.ID, gomock.Any()).Return(nil)
				mockMovieProducerRepo.EXPECT().AddMovieProducer(sampleMovie.ID, gomock.Any()).Return(kts_errors.KTS_INTERNAL_ERROR)
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
			setExpectations: func(mockMovieRepo mocks.MockMovieRepositoryI, mockMovieGenreRepo mocks.MockMovieGenreRepositoryI, mockMovieActorRepo mocks.MockMovieActorRepositoryI, mockMovieProducerRepo mocks.MockMovieProducerRepositoryI, movie *models.MovieDTOCreate) {
				mockMovieRepo.EXPECT().CreateMovie(&movie.Movies).Return(sampleMovie.ID, nil)
				mockMovieGenreRepo.EXPECT().AddMovieGenre(sampleMovie.ID, gomock.Any()).Return(nil)
				mockMovieActorRepo.EXPECT().AddMovieActor(sampleMovie.ID, gomock.Any()).Return(nil)
				mockMovieProducerRepo.EXPECT().AddMovieProducer(sampleMovie.ID, gomock.Any()).Return(nil)
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
			producerRepoMock := mocks.NewMockMovieProducerRepositoryI(mockCtrl)

			movieController := MovieController{
				MovieRepo:         movieRepoMock,
				MovieGenreRepo:    genreRepoMock,
				MovieActorRepo:    actorRepoMock,
				MovieProducerRepo: producerRepoMock,
			}

			// define expectations
			tc.setExpectations(*movieRepoMock, *genreRepoMock, *actorRepoMock, *producerRepoMock, tc.movieDTOModel)

			// WHEN
			// call RegisterUser with registrationData
			movieId, kts_err := movieController.CreateMovie(tc.movieDTOModel)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)

			if tc.expectedMoviesId && movieId == nil {
				t.Error("Expected genre ID, got nil")
			}
		})
	}
}
