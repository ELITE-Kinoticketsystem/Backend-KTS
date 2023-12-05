package repositories

import (
	"errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/google/uuid"
)

type MovieRepoI interface {
	CreateMovie(movie *model.Movies) error

	GetGenreByName(name string) (*model.Genres, error)
	CreateGenre(genre *model.Genres) error

	AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) error
}

type MovieRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (mr *MovieRepository) AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (mr *MovieRepository) CreateGenre(genre *model.Genres) error {
	// TODO: implement
	return errors.New("not implemented")
}

func (mr *MovieRepository) GetGenreByName(name string) (*model.Genres, error) {
	// TODO implement
	return nil, errors.New("not implemented")
}

func (mr *MovieRepository) CreateMovie(movie *model.Movies) error {
	// TODO: implement
	return errors.New("not implemented")
}
