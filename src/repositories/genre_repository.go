package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type GenreRepositoryI interface {
	// Genre
	GetGenres() (*[]model.Genres, *models.KTSError)
	GetGenreByName(name *string) (*model.Genres, *models.KTSError)
	CreateGenre(name *string) *models.KTSError
	UpdateGenre(genre *model.Genres) *models.KTSError
	DeleteGenre(genreId *uuid.UUID) *models.KTSError

	// All Movies with all Genres - Grouped by Genre
	GetGenresWithMovies() (*[]models.GenreWithMovies, *models.KTSError)

	// One Genre with all Movies
	GetGenreByNameWithMovies(genreName *string) (*models.GenreWithMovies, *models.KTSError)
}

type GenreRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

// Genre
func (mr *GenreRepository) GetGenres() (*[]model.Genres, *models.KTSError) {
	var genres []model.Genres

	// Create the query
	stmt := mysql.SELECT(
		table.Genres.AllColumns,
	).FROM(
		table.Genres,
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &genres)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(genres) == 0 {
		return nil, kts_errors.KTS_MOVIE_NOT_FOUND
	}

	return &genres, nil
}

func (mr *GenreRepository) GetGenreByName(name *string) (*model.Genres, *models.KTSError) {
	var genre model.Genres

	// Create the query
	stmt := mysql.SELECT(
		table.Genres.AllColumns,
	).FROM(
		table.Genres,
	).WHERE(
		table.Genres.GenreName.EQ(utils.MySqlString(*name)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &genre)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, kts_errors.KTS_MOVIE_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &genre, nil
}

func (mr *GenreRepository) CreateGenre(name *string) *models.KTSError {
	// Create the insert statement
	insertQuery := table.Genres.INSERT(table.Genres.GenreName).
		VALUES(name)

	// Execute the query
	rows, err := insertQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return kts_errors.KTS_MOVIE_NOT_FOUND
	}

	return nil
}

func (mr *GenreRepository) UpdateGenre(genre *model.Genres) *models.KTSError {

	// Create the update statement
	updateQuery := table.Genres.UPDATE(table.Genres.GenreName).
		SET(genre.GenreName).
		WHERE(table.Genres.ID.EQ(utils.MysqlUuid(genre.ID)))

	// Execute the query
	rows, err := updateQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return kts_errors.KTS_MOVIE_NOT_FOUND
	}

	return nil
}

func (mr *GenreRepository) DeleteGenre(genreId *uuid.UUID) *models.KTSError {
	

	// Create the delete statement
	deleteQuery := table.Genres.DELETE().
	WHERE(table.Genres.ID.EQ(utils.MysqlUuid(genreId)))

	// Execute the query
	rows, err := deleteQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return kts_errors.KTS_MOVIE_NOT_FOUND
	}

	return nil
}

// One Genre with all table.Movies
func (mr *GenreRepository) GetGenreByNameWithMovies(genreName *string) (*models.GenreWithMovies, *models.KTSError) {
	var movies models.GenreWithMovies

	// Create the query
	stmt := mysql.SELECT(
		table.Movies.AllColumns,
		table.Genres.AllColumns,
	).FROM(
		table.MovieGenres.
			INNER_JOIN(table.Movies, table.Movies.ID.EQ(table.MovieGenres.MovieID)).
			INNER_JOIN(table.Genres, table.Genres.ID.EQ(table.MovieGenres.GenreID)),
	).WHERE(
		table.Genres.GenreName.EQ(utils.MySqlString(*genreName)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &movies)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, kts_errors.KTS_MOVIE_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &movies, nil
}

// All Movies with all Genres - Grouped by Genre
func (mr *GenreRepository) GetGenresWithMovies() (*[]models.GenreWithMovies, *models.KTSError) {
	var genresWithMovies []models.GenreWithMovies

	// Create the query
	stmt := mysql.SELECT(
		table.Movies.AllColumns,
		table.Genres.AllColumns,
	).FROM(
		table.MovieGenres.
			INNER_JOIN(table.Movies, table.Movies.ID.EQ(table.MovieGenres.MovieID)).
			INNER_JOIN(table.Genres, table.Genres.ID.EQ(table.MovieGenres.GenreID)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &genresWithMovies)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(genresWithMovies) == 0 {
		return nil, kts_errors.KTS_MOVIE_NOT_FOUND
	}

	return &genresWithMovies, nil
}
