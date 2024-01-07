package repositories

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type MovieRepositoryI interface {
	// Movie
	GetMovies() (*[]model.Movies, *models.KTSError)
	GetMovieById(movieId *uuid.UUID) (*models.MovieWithEverything, *models.KTSError)
	GetMovieByName(movieName *string) (*model.Movies, *models.KTSError)
	CreateMovie(movie *model.Movies) (*uuid.UUID, *models.KTSError)
	UpdateMovie(movie *model.Movies) *models.KTSError
	DeleteMovie(movieId *uuid.UUID) *models.KTSError

	// All Movies with all Genres - Grouped by Movie
	GetMoviesWithGenres() (*[]models.MovieWithGenres, *models.KTSError)
}

type MovieRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

// Movie
func (mr *MovieRepository) GetMovies() (*[]model.Movies, *models.KTSError) {
	var movies []model.Movies

	// Create the query
	stmt := mysql.SELECT(
		table.Movies.AllColumns,
	).FROM(
		table.Movies,
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &movies)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(movies) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &movies, nil
}

func (mr *MovieRepository) GetMovieByName(movieName *string) (*model.Movies, *models.KTSError) {
	// Prepare vairables
	var movie model.Movies

	// Create the query
	stmt := mysql.SELECT(
		table.Movies.AllColumns,
	).FROM(
		table.Movies,
	).WHERE(
		table.Movies.Title.EQ(utils.MySqlString(*movieName)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &movie)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, kts_errors.KTS_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &movie, nil
}

func (mr *MovieRepository) CreateMovie(movie *model.Movies) (*uuid.UUID, *models.KTSError) {
	newId := uuid.New()
	movie.ID = &newId

	// Create the insert statement
	insertQuery := table.Movies.INSERT(
		table.Movies.ID,
		table.Movies.Title,
		table.Movies.Description,
		table.Movies.BannerPicURL,
		table.Movies.CoverPicURL,
		table.Movies.TrailerURL,
		table.Movies.Rating,
		table.Movies.ReleaseDate,
		table.Movies.TimeInMin,
		table.Movies.Fsk,
	).VALUES(
		utils.MysqlUuid(movie.ID),
		movie.Title,
		movie.Description,
		movie.BannerPicURL,
		movie.CoverPicURL,
		movie.TrailerURL,
		movie.Rating,
		movie.ReleaseDate,
		movie.TimeInMin,
		movie.Fsk,
	)

	// Execute the query
	rows, err := insertQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAff, err := rows.RowsAffected()
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAff == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return movie.ID, nil
}

func (mr *MovieRepository) UpdateMovie(movie *model.Movies) *models.KTSError {
	// Create the update statement
	updateQuery := table.Movies.UPDATE().
		SET(
			table.Movies.Title.SET(mysql.String(movie.Title)),
			table.Movies.Description.SET(mysql.String(movie.Description)),
			table.Movies.BannerPicURL.SET(mysql.String(*movie.BannerPicURL)),
			table.Movies.CoverPicURL.SET(mysql.String(*movie.CoverPicURL)),
			table.Movies.TrailerURL.SET(mysql.String(*movie.TrailerURL)),
			table.Movies.Rating.SET(mysql.Float(*movie.Rating)),
			table.Movies.ReleaseDate.SET(mysql.DateT(movie.ReleaseDate)),
			table.Movies.TimeInMin.SET(mysql.Int32(movie.TimeInMin)),
			table.Movies.Fsk.SET(mysql.Int32(movie.Fsk)),
		).WHERE(
		table.Movies.ID.EQ(utils.MysqlUuid(movie.ID)),
	)

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
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}

func (mr *MovieRepository) DeleteMovie(movieId *uuid.UUID) *models.KTSError {
	// Create the delete statement
	deleteQuery := table.Movies.DELETE().WHERE(table.Movies.ID.EQ(utils.MysqlUuid(movieId)))

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
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}

// All Movies with all Genres - Grouped by Movie
func (mr *MovieRepository) GetMoviesWithGenres() (*[]models.MovieWithGenres, *models.KTSError) {
	var moviesWithGenres []models.MovieWithGenres

	// Create the query
	stmt := mysql.SELECT(
		table.Movies.AllColumns,
		table.Genres.AllColumns,
	).FROM(
		table.Movies.
			LEFT_JOIN(table.MovieGenres, table.Movies.ID.EQ(table.MovieGenres.MovieID)).
			LEFT_JOIN(table.Genres, table.Genres.ID.EQ(table.MovieGenres.GenreID)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &moviesWithGenres)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(moviesWithGenres) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &moviesWithGenres, nil
}

func (mr *MovieRepository) GetMovieById(movieId *uuid.UUID) (*models.MovieWithEverything, *models.KTSError) {
	var movie models.MovieWithEverything

	// Create the query
	stmt := mysql.SELECT(
		table.Movies.AllColumns,
		table.Genres.AllColumns,
		table.Actors.AllColumns,
		table.Producers.AllColumns,
		table.Reviews.AllColumns,
		table.Users.Username,
	).FROM(
		table.Movies.
			LEFT_JOIN(table.MovieGenres, table.MovieGenres.MovieID.EQ(table.Movies.ID)).
			LEFT_JOIN(table.Genres, table.Genres.ID.EQ(table.MovieGenres.GenreID)).
			LEFT_JOIN(table.MovieActors, table.MovieActors.MovieID.EQ(table.Movies.ID)).
			LEFT_JOIN(table.Actors, table.Actors.ID.EQ(table.MovieActors.ActorID)).
			LEFT_JOIN(table.MovieProducers, table.MovieProducers.MovieID.EQ(table.Movies.ID)).
			LEFT_JOIN(table.Producers, table.Producers.ID.EQ(table.MovieProducers.ProducerID)).
			LEFT_JOIN(table.Reviews, table.Reviews.MovieID.EQ(table.Movies.ID)).
			LEFT_JOIN(table.Users, table.Users.ID.EQ(table.Reviews.UserID)),
	).WHERE(
		table.Movies.ID.EQ(utils.MysqlUuid(movieId)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &movie)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, kts_errors.KTS_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &movie, nil
}
