package repositories

import (
	"errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/google/uuid"
)

type TheaterRepoI interface {
	GetSeatsForCinemaHall(cinemaHallId *uuid.UUID) ([]schemas.Seat, error)
}

type TheatreRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (tr *TheatreRepository) GetSeatsForCinemaHall(cinemaHallId *uuid.UUID) ([]schemas.Seat, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}
