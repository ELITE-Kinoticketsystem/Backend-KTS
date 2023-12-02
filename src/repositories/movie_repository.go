package repositories

import (
	. "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"

	. "github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type MovieRepoI interface {
	CreateMovie(movie *schemas.Movie) *models.KTSError
	GetMovies() ([]models.MoviesWithGenre, *models.KTSError)
	GetMovieById(id uuid.UUID) (*schemas.Movie, *models.KTSError)
	UpdateMovie(movie *schemas.Movie) *models.KTSError
	DeleteMovie(id *uuid.UUID) *models.KTSError

	GetGenreByName(name string) (*model.Genres, *models.KTSError)
	CreateGenre(genre *model.Genres) *models.KTSError
	GetGenres() (*[]model.Genres, *models.KTSError)

	AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError

	GetMoviesByGenre(genreId *uuid.UUID) (*[]model.Movies, *models.KTSError)
	GetMoviesByGenres() (*[]model.Movies, *models.KTSError)
}

type MovieRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (mr *MovieRepository) AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError {
	// TODO: implement
	return nil
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

func (mr *MovieRepository) CreateGenre(genre *model.Genres) *models.KTSError {
	// TODO: implement
	return nil
}

func (mr *MovieRepository) CreateMovie(movie *schemas.Movie) *models.KTSError {
	// TODO: implement
	return nil
}

func (mr *MovieRepository) GetMovies() (*[]models.MoviesWithGenre, *models.KTSError) {
	var movies_with_genre []models.MoviesWithGenre

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
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &movies_with_genre)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(movies_with_genre) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &movies_with_genre, nil
}

func (mr *MovieRepository) GetMovieById(id uuid.UUID) (*models.MoviesWithGenre, *models.KTSError) {
	// Prepare vairables
	var movie_with_genre models.MoviesWithGenre
	binary_id, _ := id.MarshalBinary()

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
	err := stmt.Query(mr.DatabaseManager.GetDatabaseConnection(), &movie_with_genre)
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &movie_with_genre, nil
}

func (mr *MovieRepository) UpdateMovie(movie *schemas.Movie) *models.KTSError {
	// TODO: implement
	return nil
}

func (mr *MovieRepository) DeleteMovie(id *uuid.UUID) *models.KTSError {
	// TODO: implement
	return nil
}

func (mr *MovieRepository) GetMoviesByGenre(genreName string) (*models.GenresWithMovie, *models.KTSError) {
	var movies models.GenresWithMovie

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

func (mr *MovieRepository) GetMoviesByGenres() (*[]models.GenresWithMovie, *models.KTSError) {
	var genresWithMovies []models.GenresWithMovie

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
