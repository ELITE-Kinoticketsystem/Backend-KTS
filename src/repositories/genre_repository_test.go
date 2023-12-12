package repositories

import (
	"errors"
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
			name: "Error while querying genres",
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

			// Create a new instance of the genreRepository with the mock database connection
			genreRepo := GenreRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method under test
			genres, kts_err := genreRepo.GetGenres()

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
	sampleGenre := utils.GetSampleGenre()

	genreName := sampleGenre.GenreName

	query := "\nSELECT genres.id AS \"genres.id\",\n     genres.genre_name AS \"genres.genre_name\"\nFROM `KinoTicketSystem`.genres\nWHERE genres.genre_name = \\?;\n"
	// query := "SELECT genres.id AS \"genres.id\", genres.genre_name AS \"genres.genre_name\" FROM `KinoTicketSystem`.genres WHERE genres.genre_name = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, genreName *string)
		expectedGenre   *model.Genres
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock, genreName *string) {
				mock.ExpectQuery(query).WithArgs(genreName).WillReturnRows(
					sqlmock.NewRows([]string{"genres.id", "genres.genre_name"}),
				)
			},
			expectedGenre: nil,
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Single genre",
			setExpectations: func(mock sqlmock.Sqlmock, genreName *string) {
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
			name: "Error while querying genres",
			setExpectations: func(mock sqlmock.Sqlmock, genreName *string) {
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

			// Create a new instance of the genreRepository with the mock database connection
			genreRepo := GenreRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &genreName)

			// Call the method under test
			genre, kts_err := genreRepo.GetGenreByName(&genreName)

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

func TestCreateGenre(t *testing.T) {
	sampleGenre := utils.GetSampleGenre()

	genreName := sampleGenre.GenreName

	query := "INSERT INTO `KinoTicketSystem`.genres (genre_name) VALUES (?);"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, genreName *string)
		expectedError   *models.KTSError
	}{
		{
			name: "Successful creation",
			setExpectations: func(mock sqlmock.Sqlmock, genreName *string) {
				mock.ExpectExec(query).WithArgs(genreName).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while creating genre",
			setExpectations: func(mock sqlmock.Sqlmock, genreName *string) {
				mock.ExpectExec(query).WithArgs(genreName).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, genreName *string) {
				mock.ExpectExec(query).WithArgs(genreName).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Genre not found",
			setExpectations: func(mock sqlmock.Sqlmock, genreName *string) {
				mock.ExpectExec(query).WithArgs(genreName).WillReturnResult(sqlmock.NewResult(1, 0))
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

			// Create a new instance of the genreRepository with the mock database connection
			genreRepo := GenreRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &genreName)

			// Call the method under test
			kts_err := genreRepo.CreateGenre(&genreName)

			// Verify the results
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestUpdateGenre(t *testing.T) {
	sampleGenre := utils.GetSampleGenre()

	query := "UPDATE `KinoTicketSystem`.genres SET genre_name = ? WHERE genres.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, genre *model.Genres)
		expectedError   *models.KTSError
	}{
		{
			name: "Successful update",
			setExpectations: func(mock sqlmock.Sqlmock, genre *model.Genres) {
				mock.ExpectExec(query).WithArgs(sampleGenre.GenreName, utils.EqUUID(sampleGenre.ID)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while updating genre",
			setExpectations: func(mock sqlmock.Sqlmock, genre *model.Genres) {
				mock.ExpectExec(query).WithArgs(sampleGenre.GenreName, utils.EqUUID(sampleGenre.ID)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, genre *model.Genres) {
				mock.ExpectExec(query).WithArgs(sampleGenre.GenreName, utils.EqUUID(sampleGenre.ID)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Genre not found",
			setExpectations: func(mock sqlmock.Sqlmock, genre *model.Genres) {
				mock.ExpectExec(query).WithArgs(sampleGenre.GenreName, utils.EqUUID(sampleGenre.ID)).WillReturnResult(sqlmock.NewResult(1, 0))
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

			// Create a new instance of the genreRepository with the mock database connection

			genreRepo := GenreRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, sampleGenre)

			// Call the method under test
			kts_err := genreRepo.UpdateGenre(sampleGenre)

			// Verify the results
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestDeleteGenre(t *testing.T) {
	genreId := uuid.New()

	query := "DELETE FROM `KinoTicketSystem`.genres WHERE genres.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, genre *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Successful deletion",
			setExpectations: func(mock sqlmock.Sqlmock, genre *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&genreId)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while deleting genre",
			setExpectations: func(mock sqlmock.Sqlmock, genre *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&genreId)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, genre *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&genreId)).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Genre not found",
			setExpectations: func(mock sqlmock.Sqlmock, genre *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(utils.EqUUID(&genreId)).WillReturnResult(sqlmock.NewResult(1, 0))
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

			// Create a new instance of the genreRepository with the mock database connection

			genreRepo := GenreRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, &genreId)

			// Call the method under test
			kts_err := genreRepo.DeleteGenre(&genreId)

			// Verify the results
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

	query := "SELECT movies.id AS \"movies.id\",\n     movies.title AS \"movies.title\",\n     movies.description AS \"movies.description\",\n     movies.banner_pic_url AS \"movies.banner_pic_url\",\n     movies.cover_pic_url AS \"movies.cover_pic_url\",\n     movies.trailer_url AS \"movies.trailer_url\",\n     movies.rating AS \"movies.rating\",\n     movies.release_date AS \"movies.release_date\",\n     movies.time_in_min AS \"movies.time_in_min\",\n     movies.fsk AS \"movies.fsk\",\n     genres.id AS \"genres.id\",\n     genres.genre_name AS \"genres.genre_name\"\nFROM `KinoTicketSystem`.movie_genres\n     INNER JOIN `KinoTicketSystem`.movies ON (movies.id = movie_genres.movie_id)\n     INNER JOIN `KinoTicketSystem`.genres ON (genres.id = movie_genres.genre_id);\n"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectedGenres  *[]models.GenreWithMovies
		expectedError   *models.KTSError
	}{
		{
			name: "Empty result",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WillReturnRows(
					sqlmock.NewRows([]string{"genres.id", "genres.genre_name", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"}),
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
						[]string{"genres.id", "genres.genre_name", "movies.id", "movies.title", "movies.description", "movies.banner_pic_url", "movies.cover_pic_url", "movies.trailer_url", "movies.rating", "movies.release_date", "movies.time_in_min", "movies.fsk"},
					).AddRow(
						(*sampleGenresWithMovies)[0].ID, (*sampleGenresWithMovies)[0].GenreName, (*sampleGenresWithMovies)[0].Movies[0].ID, (*sampleGenresWithMovies)[0].Movies[0].Title, (*sampleGenresWithMovies)[0].Movies[0].Description, (*sampleGenresWithMovies)[0].Movies[0].BannerPicURL, (*sampleGenresWithMovies)[0].Movies[0].CoverPicURL, (*sampleGenresWithMovies)[0].Movies[0].TrailerURL, (*sampleGenresWithMovies)[0].Movies[0].Rating, (*sampleGenresWithMovies)[0].Movies[0].ReleaseDate, (*sampleGenresWithMovies)[0].Movies[0].TimeInMin, (*sampleGenresWithMovies)[0].Movies[0].Fsk,
					).AddRow(
						(*sampleGenresWithMovies)[0].ID, (*sampleGenresWithMovies)[0].GenreName, (*sampleGenresWithMovies)[0].Movies[1].ID, (*sampleGenresWithMovies)[0].Movies[1].Title, (*sampleGenresWithMovies)[0].Movies[1].Description, (*sampleGenresWithMovies)[0].Movies[1].BannerPicURL, (*sampleGenresWithMovies)[0].Movies[1].CoverPicURL, (*sampleGenresWithMovies)[0].Movies[1].TrailerURL, (*sampleGenresWithMovies)[0].Movies[1].Rating, (*sampleGenresWithMovies)[0].Movies[1].ReleaseDate, (*sampleGenresWithMovies)[0].Movies[1].TimeInMin, (*sampleGenresWithMovies)[0].Movies[1].Fsk,
					).AddRow(
						(*sampleGenresWithMovies)[1].ID, (*sampleGenresWithMovies)[1].GenreName, (*sampleGenresWithMovies)[1].Movies[0].ID, (*sampleGenresWithMovies)[1].Movies[0].Title, (*sampleGenresWithMovies)[1].Movies[0].Description, (*sampleGenresWithMovies)[1].Movies[0].BannerPicURL, (*sampleGenresWithMovies)[1].Movies[0].CoverPicURL, (*sampleGenresWithMovies)[1].Movies[0].TrailerURL, (*sampleGenresWithMovies)[1].Movies[0].Rating, (*sampleGenresWithMovies)[1].Movies[0].ReleaseDate, (*sampleGenresWithMovies)[1].Movies[0].TimeInMin, (*sampleGenresWithMovies)[1].Movies[0].Fsk,
					).AddRow(
						(*sampleGenresWithMovies)[1].ID, (*sampleGenresWithMovies)[1].GenreName, (*sampleGenresWithMovies)[1].Movies[1].ID, (*sampleGenresWithMovies)[1].Movies[1].Title, (*sampleGenresWithMovies)[1].Movies[1].Description, (*sampleGenresWithMovies)[1].Movies[1].BannerPicURL, (*sampleGenresWithMovies)[1].Movies[1].CoverPicURL, (*sampleGenresWithMovies)[1].Movies[1].TrailerURL, (*sampleGenresWithMovies)[1].Movies[1].Rating, (*sampleGenresWithMovies)[1].Movies[1].ReleaseDate, (*sampleGenresWithMovies)[1].Movies[1].TimeInMin, (*sampleGenresWithMovies)[1].Movies[1].Fsk,
					),
				)
			},
			expectedGenres: sampleGenresWithMovies,
			expectedError:  nil,
		},
		{
			name: "Error while querying genres",
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
			// db, mock, err := sqlmock.New()
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the genreRepository with the mock database connection
			genreRepo := GenreRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			// Call the method under test
			genres, kts_err := genreRepo.GetGenresWithMovies()

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
