package repositories

import (
	. "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"

	. "github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type MovieRepoI interface {
	// Movie
	GetMovies() (*[]model.Movies, *models.KTSError)
	GetMovieById(movieId uuid.UUID) (*model.Movies, *models.KTSError)
	CreateMovie(movie model.Movies) *models.KTSError
	UpdateMovie(movie model.Movies) *models.KTSError
	DeleteMovie(movieId uuid.UUID) *models.KTSError

	// Genre
	GetGenres() (*[]model.Genres, *models.KTSError)
	GetGenreByName(name string) (*model.Genres, *models.KTSError)
	CreateGenre(name string) *models.KTSError

	// Combine Movie and Genre
	AddMovieGenre(movieId uuid.UUID, genreId uuid.UUID) *models.KTSError

	// One Movie with all Genres
	GetMovieByIdWithGenre(movieId uuid.UUID) (*models.MovieWithGenres, *models.KTSError)
	// One Genre with all Movies
	GetGenreWithMovies(genreName string) (*models.GenreWithMovies, *models.KTSError)
	// All Movies with all Genres - Grouped by Genre
	GetGenresWithMovies() (*[]models.GenreWithMovies, *models.KTSError)
	// All Movies with all Genres - Grouped by Movie
	GetMoviesWithGenres() ([]models.MovieWithGenres, *models.KTSError)
}

type MovieRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

// Movie
func (mr *MovieRepository) GetMovies() (*[]model.Movies, *models.KTSError) {
	var movies []model.Movies

	// Create the query
	stmt := SELECT(
		Movies.AllColumns,
	).FROM(
		Movies,
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

func (mr *MovieRepository) GetMovieById(movieId uuid.UUID) (*model.Movies, *models.KTSError) {
	// Prepare vairables
	var movie model.Movies
	binary_id, _ := movieId.MarshalBinary()

	// Create the query
	stmt := SELECT(
		Movies.AllColumns,
	).FROM(
		Movies,
	).WHERE(
		Movies.ID.EQ(String(string(binary_id))),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &movie)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &movie, nil
}

func (mr *MovieRepository) CreateMovie(movie model.Movies) *models.KTSError {
	// Create the insert statement
	insertQuery := Movies.INSERT(Movies.Title, Movies.Description, Movies.BannerPicURL, Movies.CoverPicURL, Movies.TrailerURL, Movies.Rating, Movies.ReleaseDate, Movies.TimeInMin, Movies.Fsk).
		VALUES(movie.Title, movie.Description, movie.BannerPicURL, movie.CoverPicURL, movie.TrailerURL, movie.Rating, movie.ReleaseDate, movie.TimeInMin, movie.Fsk)

	// Execute the query
	_, err := insertQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

func (mr *MovieRepository) UpdateMovie(movie model.Movies) *models.KTSError {
	// Prepare variables
	binaryID, _ := movie.ID.MarshalBinary()

	// Create the update statement
	updateQuery := Movies.UPDATE().
		SET(
			Movies.Title.SET(String(movie.Title)),
			Movies.Description.SET(String(movie.Description)),
			Movies.BannerPicURL.SET(String(*movie.BannerPicURL)),
			Movies.CoverPicURL.SET(String(*movie.CoverPicURL)),
			Movies.TrailerURL.SET(String(*movie.TrailerURL)),
			Movies.Rating.SET(Float(*movie.Rating)),
			Movies.ReleaseDate.SET(Date(movie.ReleaseDate.Year(), movie.ReleaseDate.Month(), movie.ReleaseDate.Day())),
			Movies.TimeInMin.SET(Int32(movie.TimeInMin)),
			Movies.Fsk.SET(Int32(movie.Fsk)),
		).WHERE(
		Movies.ID.EQ(String(string(binaryID))),
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

func (mr *MovieRepository) DeleteMovie(movieId uuid.UUID) *models.KTSError {
	binaryID, _ := movieId.MarshalBinary()

	// Create the delete statement
	deleteQuery := Movies.DELETE().WHERE(Movies.ID.EQ(String(string(binaryID))))

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

// Genre
func (mr *MovieRepository) GetGenres() (*[]model.Genres, *models.KTSError) {
	var genres []model.Genres

	// Create the query
	stmt := SELECT(
		Genres.AllColumns,
	).FROM(
		Genres,
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &genres)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(genres) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &genres, nil
}

func (mr *MovieRepository) GetGenreByName(name string) (*model.Genres, *models.KTSError) {
	var genre model.Genres

	// Create the query
	stmt := SELECT(
		Genres.AllColumns,
	).FROM(
		Genres,
	).WHERE(
		Genres.GenreName.EQ(String(name)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &genre)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &genre, nil
}

func (mr *MovieRepository) CreateGenre(name string) *models.KTSError {
	// Create the insert statement
	insertQuery := Genres.INSERT(Genres.GenreName).
		VALUES(name)

	// Execute the query
	_, err := insertQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

// Combine Movie and Genre
func (mr *MovieRepository) AddMovieGenre(movieId uuid.UUID, genreId uuid.UUID) *models.KTSError {

	binary_movie_id, _ := movieId.MarshalBinary()
	binary_genre_id, _ := genreId.MarshalBinary()

	// Create the insert statement
	insertQuery := MovieGenres.INSERT(MovieGenres.MovieID, MovieGenres.GenreID).
		VALUES(String(string(binary_movie_id)), String(string(binary_genre_id)))

	// Execute the query
	_, err := insertQuery.Exec(mr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

// One Movie with all Genres
func (mr *MovieRepository) GetMovieByIdWithGenre(movieId uuid.UUID) (*models.MovieWithGenres, *models.KTSError) {
	var movie models.MovieWithGenres

	binary_id, _ := movieId.MarshalBinary()

	// Create the query
	stmt := SELECT(
		Movies.AllColumns,
		Genres.AllColumns,
	).FROM(
		MovieGenres.
			INNER_JOIN(Movies, Movies.ID.EQ(MovieGenres.MovieID)).
			INNER_JOIN(Genres, Genres.ID.EQ(MovieGenres.GenreID)),
	).WHERE(
		Movies.ID.EQ(String(string(binary_id))),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &movie)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &movie, nil
}

// One Genre with all Movies
func (mr *MovieRepository) GetGenreWithMovies(genreName string) (*models.GenreWithMovies, *models.KTSError) {
	var movies models.GenreWithMovies

	// Create the query
	stmt := SELECT(
		Movies.AllColumns,
		Genres.AllColumns,
	).FROM(
		MovieGenres.
			INNER_JOIN(Movies, Movies.ID.EQ(MovieGenres.MovieID)).
			INNER_JOIN(Genres, Genres.ID.EQ(MovieGenres.GenreID)),
	).WHERE(
		Genres.GenreName.EQ(String(genreName)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &movies)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &movies, nil
}

// All Movies with all Genres - Grouped by Genre
func (mr *MovieRepository) GetGenresWithMovies() (*[]models.GenreWithMovies, *models.KTSError) {
	var genresWithMovies []models.GenreWithMovies

	// Create the query
	stmt := SELECT(
		Movies.AllColumns,
		Genres.AllColumns,
	).FROM(
		MovieGenres.
			INNER_JOIN(Movies, Movies.ID.EQ(MovieGenres.MovieID)).
			INNER_JOIN(Genres, Genres.ID.EQ(MovieGenres.GenreID)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &genresWithMovies)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(genresWithMovies) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &genresWithMovies, nil
}

// All Movies with all Genres - Grouped by Movie
func (mr *MovieRepository) GetMoviesWithGenres() ([]models.MovieWithGenres, *models.KTSError) {
	var moviesWithGenres []models.MovieWithGenres

	// Create the query
	stmt := SELECT(
		Movies.AllColumns,
		Genres.AllColumns,
	).FROM(
		MovieGenres.
			INNER_JOIN(Movies, Movies.ID.EQ(MovieGenres.MovieID)).
			INNER_JOIN(Genres, Genres.ID.EQ(MovieGenres.GenreID)),
	)

	// Execute the query
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &moviesWithGenres)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(moviesWithGenres) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return moviesWithGenres, nil
}
