package repositories

import (
	"errors"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/google/uuid"
)

type TheaterRepoI interface {
	GetSeatsForCinemaHall(cinemaHallId *uuid.UUID) ([]model.Seats, error)
}

type TheatreRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (tr *TheatreRepository) GetSeatsForCinemaHall(cinemaHallId *uuid.UUID) ([]model.Seats, error) {
	// TODO: implement
	return nil, errors.New("not implemented")
}
