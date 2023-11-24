package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models/schemas"
	"github.com/google/uuid"
)

type TheaterRepoI interface {
	GetSeatsForCinemaHall(cinemaHallId *uuid.UUID) ([]schemas.Seat, error)
}
