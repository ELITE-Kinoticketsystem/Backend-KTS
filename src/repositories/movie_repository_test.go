package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetMovies(t *testing.T) {
	sampleMovies := samples.GetSampleMovies()

	query := "SELECT movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\" FROM `KinoTicketSystem`.movies;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedMovies  *[]model.Movies
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows([]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}),
				)
			},
			expectedMovies: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows(
						[]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"},
					).AddRow(
						(*sampleMovies)[0].ID, (*sampleMovies)[0].Title, (*sampleMovies)[0].Description, (*sampleMovies)[0].BannerPicURL, (*sampleMovies)[0].CoverPicURL, (*sampleMovies)[0].TrailerURL, (*sampleMovies)[0].Rating, (*sampleMovies)[0].ReleaseDate, (*sampleMovies)[0].TimeInMin, (*sampleMovies)[0].Fsk,
					).AddRow(
						(*sampleMovies)[1].ID, (*sampleMovies)[1].Title, (*sampleMovies)[1].Description, (*sampleMovies)[1].BannerPicURL, (*sampleMovies)[1].CoverPicURL, (*sampleMovies)[1].TrailerURL, (*sampleMovies)[1].Rating, (*sampleMovies)[1].ReleaseDate, (*sampleMovies)[1].TimeInMin, (*sampleMovies)[1].Fsk,
					),
				)
			},
			expectedMovies: sampleMovies,
			expectedError:  nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedMovies: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the MovieRepository with the mock database connection
			movieRepo := MovieRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method under test
			movies, kts_err := movieRepo.GetMovies()

			// Verify the results
			assert.Equal(t, tc.expectedMovies, movies)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestGetMovieByName(t *testing.T) {
	sampleMovie := samples.GetSampleMovieById()

	name := sampleMovie.Title

	query := "SELECT movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\" FROM `KinoTicketSystem`.movies WHERE movies.title = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, name *string)
		expectedMovie   *model.Movies
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock, name *string) {
				mock.ExpectQuery(query).WithArgs(name).WillReturnRows(
					sqlmock.NewRows([]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}),
				)
			},
			expectedMovie: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Single movie",
			setExpectations: func(mock sqlmock.Sqlmock, name *string) {
				mock.ExpectQuery(query).WithArgs(name).WillReturnRows(
					sqlmock.NewRows(
						[]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"},
					).AddRow(
						&sampleMovie.ID, &sampleMovie.Title, &sampleMovie.Description, &sampleMovie.BannerPicURL, &sampleMovie.CoverPicURL, &sampleMovie.TrailerURL, &sampleMovie.Rating, &sampleMovie.ReleaseDate, &sampleMovie.TimeInMin, &sampleMovie.Fsk,
					),
				)
			},
			expectedMovie: sampleMovie,
			expectedError: nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mock sqlmock.Sqlmock, name *string) {
				mock.ExpectQuery(query).WithArgs(name).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedMovie: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the MovieRepository with the mock database connection
			movieRepo := MovieRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &name)

			// Call the method under test
			movies, kts_err := movieRepo.GetMovieByName(&name)

			// Verify the results
			assert.Equal(t, tc.expectedMovie, movies)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestCreateMovie(t *testing.T) {
	sampleMovie := samples.GetSampleMovieById()

	query := "INSERT INTO `KinoTicketSystem`.movies (id, title, description, banner_pic_url, cover_pic_url, trailer_url, rating, release_date, time_in_min, fsk) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movie *model.Movies)
		expectedMovieId bool
		expectedError   *models.KTSError
	}{
		{
			name: "Create movie",
			setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedMovieId: true,
			expectedError:   nil,
		},
		{
			name: "Error while creating movie",
			setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedMovieId: false,
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedMovieId: false,
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie not found",
			setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedMovieId: false,
			expectedError:   kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the MovieRepository with the mock database connection
			movieRepo := MovieRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, sampleMovie)

			// Call the method under test
			movieId, kts_err := movieRepo.CreateMovie(sampleMovie)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			if tc.expectedMovieId && movieId == nil {
				t.Error("Expected actor ID, got nil")
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

// Update Movie
func TestUpdateMovie(t *testing.T) {
	sampleMovie := samples.GetSampleMovieById()

	query := "UPDATE `KinoTicketSystem`.movies SET title = ?, description = ?, banner_pic_url = ?, cover_pic_url = ?, trailer_url = ?, rating = ?, release_date = CAST(? AS DATE), time_in_min = ?, fsk = ? WHERE movies.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movie *model.Movies)
		expectedError   *models.KTSError
	}{
		{
			name: "Update movie",
			setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
				mock.ExpectExec(query).WithArgs(movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk, utils.EqUUID(movie.ID)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while updating movie",
			setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
				mock.ExpectExec(query).WithArgs(movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk, utils.EqUUID(movie.ID)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
				mock.ExpectExec(query).WithArgs(movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk, utils.EqUUID(movie.ID)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie not found",
			setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
				mock.ExpectExec(query).WithArgs(movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk, utils.EqUUID(movie.ID)).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the MovieRepository with the mock database connection
			movieRepo := MovieRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, sampleMovie)

			// Call the method under test
			kts_err := movieRepo.UpdateMovie(sampleMovie)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

// Delete Movie
func TestDeleteMovie(t *testing.T) {
	movieId := myid.New()

	query := "DELETE FROM `KinoTicketSystem`.movies WHERE movies.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movieId *myid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Delete movie",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(*movieId)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while deleting movie",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(*movieId)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(*movieId)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Movie not found",
			setExpectations: func(mock sqlmock.Sqlmock, movieId *myid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(*movieId)).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the MovieRepository with the mock database connection
			movieRepo := MovieRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &movieId)

			// Call the method under test
			kts_err := movieRepo.DeleteMovie(&movieId)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

// GetMoviesWithGenres
func TestGetMoviesWithGenres(t *testing.T) {
	sampleMoviesWithGenres := samples.GetSampleMoviesWithGenres()

	query := "SELECT movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\", genres.id AS \"genres.id\", genres.genre_name AS \"genres.genre_name\" FROM `KinoTicketSystem`.movies LEFT JOIN `KinoTicketSystem`.movie_genres ON (movies.id = movie_genres.movie_id) LEFT JOIN `KinoTicketSystem`.genres ON (genres.id = movie_genres.genre_id);"

	testCases := []struct {
		name                    string
		setExpectations         func(mock sqlmock.Sqlmock)
		expectedMoviesWithGerne *[]models.MovieWithGenres
		expectedError           *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows([]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "genres.id", "genres.genre_name"}),
				)
			},
			expectedMoviesWithGerne: nil,
			expectedError:           kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple movies",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows(
						[]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "genres.id", "genres.genre_name"},
					).AddRow(
						(*sampleMoviesWithGenres)[0].ID, (*sampleMoviesWithGenres)[0].Title, (*sampleMoviesWithGenres)[0].Description, (*sampleMoviesWithGenres)[0].BannerPicURL, (*sampleMoviesWithGenres)[0].CoverPicURL, (*sampleMoviesWithGenres)[0].TrailerURL, (*sampleMoviesWithGenres)[0].Rating, (*sampleMoviesWithGenres)[0].ReleaseDate, (*sampleMoviesWithGenres)[0].TimeInMin, (*sampleMoviesWithGenres)[0].Fsk, (*sampleMoviesWithGenres)[0].Genres[0].ID, (*sampleMoviesWithGenres)[0].Genres[0].GenreName,
					).AddRow(
						(*sampleMoviesWithGenres)[0].ID, (*sampleMoviesWithGenres)[0].Title, (*sampleMoviesWithGenres)[0].Description, (*sampleMoviesWithGenres)[0].BannerPicURL, (*sampleMoviesWithGenres)[0].CoverPicURL, (*sampleMoviesWithGenres)[0].TrailerURL, (*sampleMoviesWithGenres)[0].Rating, (*sampleMoviesWithGenres)[0].ReleaseDate, (*sampleMoviesWithGenres)[0].TimeInMin, (*sampleMoviesWithGenres)[0].Fsk, (*sampleMoviesWithGenres)[0].Genres[1].ID, (*sampleMoviesWithGenres)[0].Genres[1].GenreName,
					).AddRow(
						(*sampleMoviesWithGenres)[1].ID, (*sampleMoviesWithGenres)[1].Title, (*sampleMoviesWithGenres)[1].Description, (*sampleMoviesWithGenres)[1].BannerPicURL, (*sampleMoviesWithGenres)[1].CoverPicURL, (*sampleMoviesWithGenres)[1].TrailerURL, (*sampleMoviesWithGenres)[1].Rating, (*sampleMoviesWithGenres)[1].ReleaseDate, (*sampleMoviesWithGenres)[1].TimeInMin, (*sampleMoviesWithGenres)[1].Fsk, (*sampleMoviesWithGenres)[1].Genres[0].ID, (*sampleMoviesWithGenres)[1].Genres[0].GenreName,
					),
				)
			},
			expectedMoviesWithGerne: sampleMoviesWithGenres,
			expectedError:           nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedMoviesWithGerne: nil,
			expectedError:           kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the MovieRepository with the mock database connection
			movieRepo := MovieRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method under test
			movies, kts_err := movieRepo.GetMoviesWithGenres()

			// Verify the results
			assert.Equal(t, tc.expectedMoviesWithGerne, movies)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestGetMovieById(t *testing.T) {
	smplFullMovie := samples.GetSampleMovieByIdWithEverything()

	id := &smplFullMovie.ID

	query := "SELECT movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\", genres.id AS \"genres.id\", genres.genre_name AS \"genres.genre_name\", actors.id AS \"actors.id\", actors.name AS \"actors.name\", actors.birthdate AS \"actors.birthdate\", actors.description AS \"actors.description\", producers.id AS \"producers.id\", producers.name AS \"producers.name\", producers.birthdate AS \"producers.birthdate\", producers.description AS \"producers.description\", reviews.id AS \"reviews.id\", reviews.rating AS \"reviews.rating\", reviews.comment AS \"reviews.comment\", reviews.datetime AS \"reviews.datetime\", reviews.is_spoiler AS \"reviews.is_spoiler\", reviews.user_id AS \"reviews.user_id\", reviews.movie_id AS \"reviews.movie_id\", users.username AS \"users.username\" FROM `KinoTicketSystem`.movies LEFT JOIN `KinoTicketSystem`.movie_genres ON (movie_genres.movie_id = movies.id) LEFT JOIN `KinoTicketSystem`.genres ON (genres.id = movie_genres.genre_id) LEFT JOIN `KinoTicketSystem`.movie_actors ON (movie_actors.movie_id = movies.id) LEFT JOIN `KinoTicketSystem`.actors ON (actors.id = movie_actors.actor_id) LEFT JOIN `KinoTicketSystem`.movie_producers ON (movie_producers.movie_id = movies.id) LEFT JOIN `KinoTicketSystem`.producers ON (producers.id = movie_producers.producer_id) LEFT JOIN `KinoTicketSystem`.reviews ON (reviews.movie_id = movies.id) LEFT JOIN `KinoTicketSystem`.users ON (users.id = reviews.user_id) WHERE movies.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, id *myid.UUID)
		expectedMovie   *models.MovieWithEverything
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock, id *myid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(*id)).WillReturnRows(
					sqlmock.NewRows([]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "genres.id", "genres.genre_name", "actors.id", "actors.name", "actors.birthdate", "actors.description", "producers.id", "producers.name", "producers.birthdate", "producers.description", "reviews.id", "reviews.rating", "reviews.comment", "reviews.datetime", "reviews.is_spoiler", "reviews.user_id", "reviews.movie_id", "users.username"}),
				)
			},
			expectedMovie: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Single movie",
			setExpectations: func(mock sqlmock.Sqlmock, id *myid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(*id)).WillReturnRows(
					sqlmock.NewRows(
						[]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "genres.id", "genres.genre_name", "actors.id", "actors.name", "actors.birthdate", "actors.description", "producers.id", "producers.name", "producers.birthdate", "producers.description", "reviews.id", "reviews.rating", "reviews.comment", "reviews.datetime", "reviews.is_spoiler", "reviews.user_id", "reviews.movie_id", "users.username"},
					).AddRow(
						&smplFullMovie.ID, &smplFullMovie.Title, &smplFullMovie.Description, &smplFullMovie.BannerPicURL, &smplFullMovie.CoverPicURL, &smplFullMovie.TrailerURL, &smplFullMovie.Rating, &smplFullMovie.ReleaseDate, &smplFullMovie.TimeInMin, &smplFullMovie.Fsk, &smplFullMovie.Genres[0].ID, &smplFullMovie.Genres[0].GenreName, &smplFullMovie.Actors[0].ID, &smplFullMovie.Actors[0].Name, &smplFullMovie.Actors[0].Birthdate, &smplFullMovie.Actors[0].Description, &smplFullMovie.Producers[0].ID, &smplFullMovie.Producers[0].Name, &smplFullMovie.Producers[0].Birthdate, &smplFullMovie.Producers[0].Description, &smplFullMovie.Reviews[0].Review.ID, &smplFullMovie.Reviews[0].Review.Rating, &smplFullMovie.Reviews[0].Review.Comment, &smplFullMovie.Reviews[0].Review.Datetime, &smplFullMovie.Reviews[0].Review.IsSpoiler, &smplFullMovie.Reviews[0].Review.UserID, &smplFullMovie.Reviews[0].Review.MovieID, &smplFullMovie.Reviews[0].Username,
					),
				)
			},
			expectedMovie: smplFullMovie,
			expectedError: nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mock sqlmock.Sqlmock, id *myid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(*id)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedMovie: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the MovieRepository with the mock database connection
			movieRepo := MovieRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, id)

			// Call the method under test
			movies, kts_err := movieRepo.GetMovieById(id)

			// Verify the results
			assert.Equal(t, tc.expectedMovie, movies)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}
