package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/google/uuid"
)

type MovieRepoI interface {
	CreateMovie(movie *schemas.Movie) (*schemas.Movie, error)
	GetMovieById(id *uuid.UUID) (*schemas.Movie, error)
	GetAllMovies() ([]schemas.Movie, error)
	UpdateMovie(movie *schemas.Movie) error
	DeleteMovie(id *uuid.UUID) error

	GetGenreById(id *uuid.UUID) (*schemas.Genre, error)
	GetGenreByName(name string) (*schemas.Genre, error)
	GetAllGenres() ([]schemas.Genre, error)
	CreateGenre(genre *schemas.Genre) (*schemas.Genre, error)

	AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) error
}
