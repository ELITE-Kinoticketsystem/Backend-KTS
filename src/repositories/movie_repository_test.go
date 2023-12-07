package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetMovies(t *testing.T) {
	sampleMovies := utils.GetSampleMovies()

	query := "\nSELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\"\nFROM `KinoTicketSystem`.movies;\n"

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

func TestGetMovieById(t *testing.T) {
	sampleMovie := utils.GetSampleMovieById()

	id := sampleMovie.ID

	query := "SELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\"\nFROM `KinoTicketSystem`.movies\nWHERE movies.id = ?;\n"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, id *uuid.UUID)
		expectedMovie   *model.Movies
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(id)).WillReturnRows(
					sqlmock.NewRows([]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}),
				)
			},
			expectedMovie: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Single movie",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(id)).WillReturnRows(
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
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(id)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedMovie: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			// db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

func TestCreateMovie(t *testing.T) {
	sampleMovie := utils.GetSampleMovieById()

	query := "INSERT INTO `KinoTicketSystem`.movies (title, description, banner_pic_url, cover_pic_url, trailer_url, rating, release_date, time_in_min, fsk) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);\n"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, movie *model.Movies)
		expectedError   *models.KTSError
	}{
		{
			name: "Create movie",
			setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
				mock.ExpectExec(query).WithArgs(movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while creating movie",
			setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
				mock.ExpectExec(query).WithArgs(movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			// db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
			kts_err := movieRepo.CreateMovie(*sampleMovie)

			// Verify
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

// Update Movie
func TestUpdateMovie(t *testing.T) {
	sampleMovie := utils.GetSampleMovieById()

	// query := "UPDATE `KinoTicketSystem`.movies SET title = ?, description = ?, banner_pic_url = ?, cover_pic_url = ?, trailer_url = ?, rating = ?, release_date = ?, time_in_min = ?, fsk = ? WHERE movies.id = ?;\n"
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
		// {
		// 	name: "Error while updating movie",
		// 	setExpectations: func(mock sqlmock.Sqlmock, movie *model.Movies) {
		// 		mock.ExpectExec(query).WithArgs(movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk, movie.ID).WillReturnError(sqlmock.ErrCancelled)
		// 	},
		// 	expectedError: kts_errors.KTS_INTERNAL_ERROR,
		// },
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

func TestGetMovieByIdWithGenre(t *testing.T) {
	sampleMovieByIdWithGenre := utils.GetSampleMovieByIdWithGenre()

	id := sampleMovieByIdWithGenre.ID

	query := "SELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\",\n     genres.id AS \"genres.id\",\n     genres.genre_name AS \"genres.genre_name\"\nFROM `KinoTicketSystem`.movie_genres\n     INNER JOIN `KinoTicketSystem`.movies ON (movies.id = movie_genres.movie_id)\n     INNER JOIN `KinoTicketSystem`.genres ON (genres.id = movie_genres.genre_id)\nWHERE movies.id = ?;\n"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, id *uuid.UUID)
		expectedMovie   *models.MovieWithGenres
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(id)).WillReturnRows(
					sqlmock.NewRows([]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "genres.id", "genres.genre_name"}),
				)
			},
			expectedMovie: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Single movie",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(id)).WillReturnRows(
					sqlmock.NewRows(
						[]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "genres.id", "genres.genre_name"},
					).AddRow(
						&sampleMovieByIdWithGenre.ID, &sampleMovieByIdWithGenre.Title, &sampleMovieByIdWithGenre.Description, &sampleMovieByIdWithGenre.BannerPicURL, &sampleMovieByIdWithGenre.CoverPicURL, &sampleMovieByIdWithGenre.TrailerURL, &sampleMovieByIdWithGenre.Rating, &sampleMovieByIdWithGenre.ReleaseDate, &sampleMovieByIdWithGenre.TimeInMin, &sampleMovieByIdWithGenre.Fsk, &sampleMovieByIdWithGenre.Genres[0].ID, &sampleMovieByIdWithGenre.Genres[0].GenreName,
					).AddRow(
						&sampleMovieByIdWithGenre.ID, &sampleMovieByIdWithGenre.Title, &sampleMovieByIdWithGenre.Description, &sampleMovieByIdWithGenre.BannerPicURL, &sampleMovieByIdWithGenre.CoverPicURL, &sampleMovieByIdWithGenre.TrailerURL, &sampleMovieByIdWithGenre.Rating, &sampleMovieByIdWithGenre.ReleaseDate, &sampleMovieByIdWithGenre.TimeInMin, &sampleMovieByIdWithGenre.Fsk, &sampleMovieByIdWithGenre.Genres[1].ID, &sampleMovieByIdWithGenre.Genres[1].GenreName,
					),
				)
			},
			expectedMovie: sampleMovieByIdWithGenre,
			expectedError: nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(id)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedMovie: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			// db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
			movies, kts_err := movieRepo.GetMovieByIdWithGenre(id)

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

// GetMoviesWithGenres
func TestGetMoviesWithGenres(t *testing.T) {
	sampleMoviesWithGenres := utils.GetSampleMoviesWithGenres()

	query := "SELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\",\n     genres.id AS \"genres.id\",\n     genres.genre_name AS \"genres.genre_name\"\nFROM `KinoTicketSystem`.movie_genres\n     INNER JOIN `KinoTicketSystem`.movies ON (movies.id = movie_genres.movie_id)\n     INNER JOIN `KinoTicketSystem`.genres ON (genres.id = movie_genres.genre_id);\n"

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
