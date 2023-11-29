package repositories

import (
	"errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/google/uuid"
)

type MovieRepoI interface {
	CreateMovie(movie *schemas.Movie) error

	GetGenreByName(name string) (*schemas.Genre, error)
	CreateGenre(genre *schemas.Genre) error

	AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) error
}

type MovieRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (mr *MovieRepository) AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (mr *MovieRepository) CreateGenre(genre *schemas.Genre) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (mr *MovieRepository) GetGenreByName(name string) (*schemas.Genre, error) {
	// TODO implement
	return nil, errors.New("not implemented")
}

func (mr *MovieRepository) CreateMovie(movie *schemas.Movie) error {
	// TODO: implement
	return errors.New("not implemented")
}
