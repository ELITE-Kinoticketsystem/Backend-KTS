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
			db, mock, err := sqlmock.New()
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

// "jet: Query: could not match actual sql: \"SELECT movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS no \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\" FROM `KinoTicketSystem`.movies WHERE movies.id =?;\"

func TestGetMovieById(t *testing.T) {
	sampleMovie := utils.GetSampleMovieById()

	id := sampleMovie.ID

	// query := "SELECT movies.id AS \"movies.id\",\n" +
	// 	"movies.title AS \"movies.title\",\n" +
	// 	"movies.description AS \"movies.description\",\n" +
	// 	"movies.banner_pic_url AS \"movies.banner_pic_url\",\n" +
	// 	"movies.cover_pic_url AS \"movies.cover_pic_url\",\n" +
	// 	"movies.trailer_url AS \"movies.trailer_url\",\n" +
	// 	"movies.rating AS \"movies.rating\",\n" +
	// 	"movies.release_date AS \"movies.release_date\",\n" +
	// 	"movies.time_in_min AS \"movies.time_in_min\",\n" +
	// 	"movies.fsk AS \"movies.fsk\"\n" +
	// 	"FROM `KinoTicketSystem`.movies\n" +
	// 	"WHERE movies.id \\= \\?;\n"
	// query := `SELECT movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS no \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\" FROM KinoTicketSystem.movies WHERE movies.id = ?;`
	// query := `\\nSELECT movies\.id AS \\"movies\.id\\",\\n movies\.title AS \\"movies\.title\\",\\n movies\.description AS \\"movies\.description\\",\\n movies\.banner_pic_url AS \\"movies\.banner_pic_url\\",\\n movies\.cover_pic_url AS \\"movies\.cover_pic_url\\",\\n movies\.trailer_url AS \\"movies\.trailer_url\\",\\n movies\.rating AS \\"movies\.rating\\",\\n movies\.release_date AS \\"movies\.release_date\\",\\n movies\.time_in_min AS \\"movies\.time_in_min\\",\\n movies\.fsk AS \\"movies\.fsk\\"\\nFROM movies\\nWHERE movies\.id \= \?;\\n`
	// query := "\nSELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\"\nFROM `KinoTicketSystem`.movies\nWHERE movies.id \\= \\?;\n"
	// query := "\nSELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\"\nFROM KinoTicketSystem.movies\nWHERE movies.id = ?;\n"
	// query := "SELECT movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS no \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\" FROM `KinoTicketSystem`.movies WHERE movies.id = ?;"
	// query := "SELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\"\nFROM `KinoTicketSystem`.movies\nWHERE movies.id = ?;"
	query := "SELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\"\nFROM `KinoTicketSystem`.movies\nWHERE movies.id = ?;\n"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, id *uuid.UUID)
		expectedMovie   *model.Movies
		expectedError   *models.KTSError
	}{
		// {
		// 	name: "Empty result",
		// 	setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
		// 		mock.ExpectQuery(query).WithArgs(id).WillReturnRows(
		// 			sqlmock.NewRows([]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}),
		// 		)
		// 	},
		// 	expectedMovie: nil,
		// 	expectedError: kts_errors.KTS_NOT_FOUND,
		// },
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
		// {
		// 	name: "Error while querying movies",
		// 	setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
		// 		mock.ExpectQuery(query).WithArgs(id).WillReturnError(sqlmock.ErrCancelled)
		// 	},
		// 	expectedMovie: nil,
		// 	expectedError: kts_errors.KTS_INTERNAL_ERROR,
		// },
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

// Create Movie
// Update Movie
// Delete Movie

func TestGetMovieByIdWithGenre(t *testing.T) {
	sampleMovieByIdWithGenre := utils.GetSampleMovieByIdWithGenre()

	id := sampleMovieByIdWithGenre.ID

	query := "SELECT movies.id AS \"movies.id\",\n" +
		"movies.title AS \"movies.title\",\n" +
		"movies.description AS \"movies.description\",\n" +
		"movies.banner_pic_url AS \"movies.banner_pic_url\",\n" +
		"movies.cover_pic_url AS \"movies.cover_pic_url\",\n" +
		"movies.trailer_url AS \"movies.trailer_url\",\n" +
		"movies.rating AS \"movies.rating\",\n" +
		"movies.release_date AS \"movies.release_date\",\n" +
		"movies.time_in_min AS \"movies.time_in_min\",\n" +
		"movies.fsk AS \"movies.fsk\",\n" +
		"genres.id AS \"genres.id\",\n" +
		"genres.genre_name AS \"genres.genre_name\"\n" +
		"FROM `KinoTicketSystem`.movies\n" +
		"LEFT JOIN `KinoTicketSystem`.movie_genre ON movie_genre.movie_id = movies.id\n" +
		"LEFT JOIN `KinoTicketSystem`.genres ON genres.id = movie_genre.genre_id\n" +
		"WHERE movies.id = \\?;\n"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, id *uuid.UUID)
		expectedMovie   *models.MovieWithGenres
		expectedError   *models.KTSError
	}{
		// {
		// 	name: "Empty result",
		// 	setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
		// 		mock.ExpectQuery(query).WithArgs(id).WillReturnRows(
		// 			sqlmock.NewRows([]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "genres.id", "genres.genre_name"}),
		// 		)
		// 	},
		// 	expectedMovie: nil,
		// 	expectedError: kts_errors.KTS_NOT_FOUND,
		// },
		{
			name: "Single movie",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(id).WillReturnRows(
					sqlmock.NewRows(
						[]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "genres.id", "genres.genre_name", "genres.id", "genres.genre_name"},
					).AddRow(
						(*sampleMovieByIdWithGenre).ID, (*sampleMovieByIdWithGenre).Title, (*sampleMovieByIdWithGenre).Description, (*sampleMovieByIdWithGenre).BannerPicURL, (*sampleMovieByIdWithGenre).CoverPicURL, (*sampleMovieByIdWithGenre).TrailerURL, (*sampleMovieByIdWithGenre).Rating, (*sampleMovieByIdWithGenre).ReleaseDate, (*sampleMovieByIdWithGenre).TimeInMin, (*sampleMovieByIdWithGenre).Fsk, (*sampleMovieByIdWithGenre).Genres[0].ID, (*sampleMovieByIdWithGenre).Genres[0].GenreName, (*sampleMovieByIdWithGenre).Genres[1].ID, (*sampleMovieByIdWithGenre).Genres[1].GenreName,
					),
				)
			},
			expectedMovie: sampleMovieByIdWithGenre,
			expectedError: nil,
		},
		// {
		// 	name: "Error while querying movies",
		// 	setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
		// 		mock.ExpectQuery(query).WithArgs(id).WillReturnError(sqlmock.ErrCancelled)
		// 	},
		// 	expectedMovie: nil,
		// 	expectedError: kts_errors.KTS_INTERNAL_ERROR,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new mock database connection
			db, mock, err := sqlmock.New()
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