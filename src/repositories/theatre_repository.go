package repositories

import (
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type TheaterRepoI interface {
	GetSeatsForCinemaHall(cinemaHallId *uuid.UUID) ([]model.Seats, *models.KTSError)
}

type TheatreRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (tr *TheatreRepository) GetSeatsForCinemaHall(cinemaHallId *uuid.UUID) ([]model.Seats, *models.KTSError) {
	stmt := table.Seats.SELECT(table.Seats.AllColumns).WHERE(table.Seats.CinemaHallID.EQ(utils.MysqlUuid(cinemaHallId)))

	var seats []model.Seats

	err := stmt.Query(tr.DatabaseManager.GetDatabaseConnection(), &seats)

	if err != nil {
		log.Println(err.Error())
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return seats, nil
}
