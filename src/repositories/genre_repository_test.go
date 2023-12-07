package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetGenres(t *testing.T) {
	sampleGenres := utils.GetSampleGenres()

	query := "\nSELECT genres.id AS \"genres.id\",\n     genres.genre_name AS \"genres.genre_name\"\nFROM `KinoTicketSystem`.genres;\n"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedGenres  *[]model.Genres
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows([]string{"genres.id", "genres.genre_name"}),
				)
			},
			expectedGenres: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Multiple genres",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows(
						[]string{"genres.id", "genres.genre_name"},
					).AddRow(
						(*sampleGenres)[0].ID, (*sampleGenres)[0].GenreName,
					).AddRow(
						(*sampleGenres)[1].ID, (*sampleGenres)[1].GenreName,
					),
				)
			},
			expectedGenres: sampleGenres,
			expectedError:  nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedGenres: nil,
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
			genres, kts_err := movieRepo.GetGenres()

			// Verify the results
			assert.Equal(t, tc.expectedGenres, genres)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestGetGenreByName(t *testing.T) {
	sampleGenre := utils.GetSampleGenreByName()

	genreName := sampleGenre.GenreName

	query := "\nSELECT genres.id AS \"genres.id\",\n     genres.genre_name AS \"genres.genre_name\"\nFROM `KinoTicketSystem`.genres\nWHERE genres.genre_name = \\?;\n"
	// query := "SELECT genres.id AS \"genres.id\", genres.genre_name AS \"genres.genre_name\" FROM `KinoTicketSystem`.genres WHERE genres.genre_name = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, genreName string)
		expectedGenre   *model.Genres
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock, genreName string) {
				mock.ExpectQuery(query).WithArgs(genreName).WillReturnRows(
					sqlmock.NewRows([]string{"genres.id", "genres.genre_name"}),
				)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Single genre",
			setExpectations: func(mock sqlmock.Sqlmock, genreName string) {
				mock.ExpectQuery(query).WithArgs(genreName).WillReturnRows(
					sqlmock.NewRows(
						[]string{"genres.id", "genres.genre_name"},
					).AddRow(
						(*sampleGenre).ID, (*sampleGenre).GenreName,
					),
				)
			},
			expectedGenre: sampleGenre,
			expectedError: nil,
		},
		{
			name: "Error while querying movies",
			setExpectations: func(mock sqlmock.Sqlmock, genreName string) {
				mock.ExpectQuery(query).WithArgs(genreName).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
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

			tc.setExpectations(mock, genreName)

			// Call the method under test
			genre, kts_err := movieRepo.GetGenreByName(genreName)

			// Verify the results
			assert.Equal(t, tc.expectedGenre, genre)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

// Create Genre
// Update Genre
// Delete Genre

func TestGetGenreByNameWithMovies(t *testing.T) {
	sampleGenreByNameWithMovies := utils.GetSampleGenreByNameWithMovies()

	genreName := sampleGenreByNameWithMovies.GenreName

	// query := "SELECT genres.id AS \"genres.id\",\n" +
	// 	"genres.genre_name AS \"genres.genre_name\",\n" +
	// 	"movies.id AS \"movies.id\",\n" +
	// 	"movies.title AS \"movies.title\",\n" +
	// 	"movies.description AS \"movies.description\",\n" +
	// 	"movies.banner_pic_url AS \"movies.banner_pic_url\",\n" +
	// 	"movies.cover_pic_url AS \"movies.cover_pic_url\",\n" +
	// 	"movies.trailer_url AS \"movies.trailer_url\",\n" +
	// 	"movies.rating AS \"movies.rating\",\n" +
	// 	"movies.release_date AS \"movies.release_date\",\n" +
	// 	"movies.time_in_min AS \"movies.time_in_min\",\n" +
	// 	"movies.fsk AS \"movies.fsk\"\n" +
	// 	"FROM `KinoTicketSystem`.genres\n" +
	// 	"LEFT JOIN `KinoTicketSystem`.movie_genre ON movie_genre.genre_id = genres.id\n" +
	// 	"LEFT JOIN `KinoTicketSystem`.movies ON movies.id = movie_genre.movie_id\n" +
	// 	"WHERE genres.genre_name = \\?;\n"

	query := "\nSELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\",\n     genres.id AS \"genres.id\",\n     genres.genre_name AS \"genres.genre_name\"\nFROM `KinoTicketSystem`.movie_genres\n     INNER JOIN `KinoTicketSystem`.movies ON (movies.id = movie_genres.movie_id)\n     INNER JOIN `KinoTicketSystem`.genres ON (genres.id = movie_genres.genre_id)\nWHERE genres.genre_name = \\?;\n"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, genreName string)
		expectedGenre   *models.GenreWithMovies
		expectedError   *models.KTSError
	}{
		// {
		// 	name: "Empty result",
		// 	setExpectations: func(mock sqlmock.Sqlmock, genreName string) {
		// 		mock.ExpectQuery(query).WithArgs(genreName).WillReturnRows(
		// 			sqlmock.NewRows([]string{"genres.id", "genres.genre_name", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}),
		// 		)
		// 	},
		// 	expectedGenre: nil,
		// 	expectedError: kts_errors.KTS_NOT_FOUND,
		// },
		{
			name: "Single genre",
			setExpectations: func(mock sqlmock.Sqlmock, genreName string) {
				mock.ExpectQuery(query).WithArgs(genreName).WillReturnRows(
					sqlmock.NewRows(
						[]string{"genres.id", "genres.genre_name", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"},
					).AddRow(
						(*sampleGenreByNameWithMovies).ID, (*sampleGenreByNameWithMovies).GenreName, (*sampleGenreByNameWithMovies).Movies[0].ID, (*sampleGenreByNameWithMovies).Movies[0].Title, (*sampleGenreByNameWithMovies).Movies[0].Description, (*sampleGenreByNameWithMovies).Movies[0].BannerPicURL, (*sampleGenreByNameWithMovies).Movies[0].CoverPicURL, (*sampleGenreByNameWithMovies).Movies[0].TrailerURL, (*sampleGenreByNameWithMovies).Movies[0].Rating, (*sampleGenreByNameWithMovies).Movies[0].ReleaseDate, (*sampleGenreByNameWithMovies).Movies[0].TimeInMin, (*sampleGenreByNameWithMovies).Movies[0].Fsk, (*sampleGenreByNameWithMovies).Movies[1].ID, (*sampleGenreByNameWithMovies).Movies[1].Title, (*sampleGenreByNameWithMovies).Movies[1].Description, (*sampleGenreByNameWithMovies).Movies[1].BannerPicURL, (*sampleGenreByNameWithMovies).Movies[1].CoverPicURL, (*sampleGenreByNameWithMovies).Movies[1].TrailerURL, (*sampleGenreByNameWithMovies).Movies[1].Rating, (*sampleGenreByNameWithMovies).Movies[1].ReleaseDate, (*sampleGenreByNameWithMovies).Movies[1].TimeInMin, (*sampleGenreByNameWithMovies).Movies[1].Fsk,
					),
				)
			},
			expectedGenre: sampleGenreByNameWithMovies,
			expectedError: nil,
		},
		// {
		// 	name: "Error while querying movies",
		// 	setExpectations: func(mock sqlmock.Sqlmock, genreName string) {
		// 		mock.ExpectQuery(query).WithArgs(genreName).WillReturnError(sqlmock.ErrCancelled)
		// 	},
		// 	expectedGenre: nil,
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

			tc.setExpectations(mock, genreName)

			// Call the method under test
			genre, kts_err := movieRepo.GetGenreByNameWithMovies(genreName)

			// Verify the results
			assert.Equal(t, tc.expectedGenre, genre)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestGetGenresWithMovies(t *testing.T) {
	sampleGenresWithMovies := utils.GetSampleGenresWithMovies()

	// query := "\nSELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\",\n     genres.id AS \"genres.id\",\n     genres.genre_name AS \"genres.genre_name\"\nFROM `KinoTicketSystem`.movie_genres\n     INNER JOIN `KinoTicketSystem`.movies ON (movies.id = movie_genres.movie_id)\n     INNER JOIN `KinoTicketSystem`.genres ON (genres.id = movie_genres.genre_id);\n"
	// query := "\nSELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\",\n     genres.id AS \"genres.id\",\n     genres.genre_name AS \"genres.genre_name\"\nFROM `KinoTicketSystem`.movie_genres\n     INNER JOIN `KinoTicketSystem`.movies ON (movies.id = movie_genres.movie_id)\n     INNER JOIN `KinoTicketSystem`.genres ON (genres.id = movie_genres.genre_id);\n"
	// 	query := `
	// 	SELECT movies.id AS "movies.id",
	// 		movies.title AS "movies.title",
	// 		movies.description AS "movies.description",
	// 		movies.banner_pic_url AS "movies.banner_pic_url",
	// 		movies.cover_pic_url AS "movies.cover_pic_url",
	// 		movies.trailer_url AS "movies.trailer_url",
	// 		movies.rating AS "movies.rating",
	// 		movies.release_date AS "movies.release_date",
	// 		movies.time_in_min AS "movies.time_in_min",
	// 		movies.fsk AS "movies.fsk",
	// 		genres.id AS "genres.id",
	// 		genres.genre_name AS "genres.genre_name"
	// FROM KinoTicketSystem.movie_genres
	// 	INNER JOIN KinoTicketSystem.movies ON (movies.id = movie_genres.movie_id)
	// 	INNER JOIN KinoTicketSystem.genres ON (genres.id = movie_genres.genre_id);`

	query := "\"SELECT movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\", genres.id AS \"genres.id\", genres.genre_name AS \"genres.genre_name\" FROM `KinoTicketSystem`.movie_genres INNER JOIN `KinoTicketSystem`.movies ON (movies.id = movie_genres.movie_id) INNER JOIN `KinoTicketSystem`.genres ON (genres.id = movie_genres.genre_id);\" with expected regexp \"SELECT movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\", genres.id AS \"genres.id\", genres.genre_name AS \"genres.genre_name\" FROM KinoTicketSystem.movie_genres INNER JOIN KinoTicketSystem.movies ON (movies.id = movie_genres.movie_id) INNER JOIN KinoTicketSystem.genres ON (genres.id = movie_genres.genre_id);\""
	// query := `SELECT movies\.id AS \\"movies\.id\\", movies\.title AS \\"movies\.title\\", movies\.description AS \\"movies\.description\\", movies\.banner_pic_url AS \\"movies\.banner_pic_url\\", movies\.cover_pic_url AS \\"movies\.cover_pic_url\\", movies\.trailer_url AS \\"movies\.trailer_url\\", movies\.rating AS \\"movies\.rating\\", movies\.release_date AS \\"movies\.release_date\\", movies\.time_in_min AS \\"movies\.time_in_min\\", movies\.fsk AS \\"movies\.fsk\\", genres\.id AS \\"genres\.id\\", genres\.genre_name AS \\"genres\.genre_name\\" FROM 'KinoTicketSystem'\.movie_genres INNER JOIN 'KinoTicketSystem'\.movies ON \(movies\.id \= movie_genres\.movie_id\) INNER JOIN 'KinoTicketSystem'\.genres ON \(genres\.id \= movie_genres\.genre_id\);`

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedGenres  *[]models.GenreWithMovies
		expectedError   *models.KTSError
	}{
		// {
		// 	name: "Empty result",
		// 	setExpectations: func(mock sqlmock.Sqlmock) {
		// 		mock.ExpectQuery(query).WillReturnRows(
		// 			sqlmock.NewRows([]string{"movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "genres.id", "genres.genre_name"}),
		// 		)
		// 	},
		// 	expectedGenres: nil,
		// 	expectedError:  kts_errors.KTS_NOT_FOUND,
		// },
		{
			name: "Multiple genres",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows(
						[]string{"genres.id", "genres.genre_name", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk", "movies.id", "movies.title", "movies.description", "movies.banner_pic", "movies.cover_pic", "movies.trailer", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"},
					).AddRow(
						(*sampleGenresWithMovies)[0].ID, (*sampleGenresWithMovies)[0].GenreName, (*sampleGenresWithMovies)[0].Movies[0].ID, (*sampleGenresWithMovies)[0].Movies[0].Title, (*sampleGenresWithMovies)[0].Movies[0].Description, (*sampleGenresWithMovies)[0].Movies[0].BannerPicURL, (*sampleGenresWithMovies)[0].Movies[0].CoverPicURL, (*sampleGenresWithMovies)[0].Movies[0].TrailerURL, (*sampleGenresWithMovies)[0].Movies[0].Rating, (*sampleGenresWithMovies)[0].Movies[0].ReleaseDate, (*sampleGenresWithMovies)[0].Movies[0].TimeInMin, (*sampleGenresWithMovies)[0].Movies[0].Fsk, (*sampleGenresWithMovies)[0].Movies[1].ID, (*sampleGenresWithMovies)[0].Movies[1].Title, (*sampleGenresWithMovies)[0].Movies[1].Description, (*sampleGenresWithMovies)[0].Movies[1].BannerPicURL, (*sampleGenresWithMovies)[0].Movies[1].CoverPicURL, (*sampleGenresWithMovies)[0].Movies[1].TrailerURL, (*sampleGenresWithMovies)[0].Movies[1].Rating, (*sampleGenresWithMovies)[0].Movies[1].ReleaseDate, (*sampleGenresWithMovies)[0].Movies[1].TimeInMin, (*sampleGenresWithMovies)[0].Movies[1].Fsk,
					).AddRow(
						(*sampleGenresWithMovies)[1].ID, (*sampleGenresWithMovies)[1].GenreName, (*sampleGenresWithMovies)[1].Movies[0].ID, (*sampleGenresWithMovies)[1].Movies[0].Title, (*sampleGenresWithMovies)[1].Movies[0].Description, (*sampleGenresWithMovies)[1].Movies[0].BannerPicURL, (*sampleGenresWithMovies)[1].Movies[0].CoverPicURL, (*sampleGenresWithMovies)[1].Movies[0].TrailerURL, (*sampleGenresWithMovies)[1].Movies[0].Rating, (*sampleGenresWithMovies)[1].Movies[0].ReleaseDate, (*sampleGenresWithMovies)[1].Movies[0].TimeInMin, (*sampleGenresWithMovies)[1].Movies[0].Fsk, (*sampleGenresWithMovies)[1].Movies[1].ID, (*sampleGenresWithMovies)[1].Movies[1].Title, (*sampleGenresWithMovies)[1].Movies[1].Description, (*sampleGenresWithMovies)[1].Movies[1].BannerPicURL, (*sampleGenresWithMovies)[1].Movies[1].CoverPicURL, (*sampleGenresWithMovies)[1].Movies[1].TrailerURL, (*sampleGenresWithMovies)[1].Movies[1].Rating, (*sampleGenresWithMovies)[1].Movies[1].ReleaseDate, (*sampleGenresWithMovies)[1].Movies[1].TimeInMin, (*sampleGenresWithMovies)[1].Movies[1].Fsk,
					),
				)
			},
			expectedGenres: sampleGenresWithMovies,
			expectedError:  nil,
		},
		// {
		// 	name: "Error while querying movies",
		// 	setExpectations: func(mock sqlmock.Sqlmock) {
		// 		mock.ExpectQuery(query).WillReturnError(sqlmock.ErrCancelled)
		// 	},
		// 	expectedGenres: nil,
		// 	expectedError:  kts_errors.KTS_INTERNAL_ERROR,
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

			tc.setExpectations(mock)

			// Call the method under test
			genres, kts_err := movieRepo.GetGenresWithMovies()

			// Verify the results
			assert.Equal(t, tc.expectedGenres, genres)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}
