package repositories

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type TheaterRepoI interface {
	CreateTheatre(theatre model.Theatres) *models.KTSError
	GetSeatsForCinemaHall(cinemaHallId *uuid.UUID) ([]model.Seats, *models.KTSError)
	CreateAddress(address model.Addresses) *models.KTSError
}

type TheatreRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (tr *TheatreRepository) CreateTheatre(theatre model.Theatres) *models.KTSError {
	stmt := table.Theatres.INSERT(
		table.Theatres.ID,
		table.Theatres.Name,
		table.Theatres.AddressID,
	).MODEL(theatre)
	_, err := stmt.Exec(tr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	return nil
}

func (tr *TheatreRepository) GetSeatsForCinemaHall(cinemaHallId *uuid.UUID) ([]model.Seats, *models.KTSError) {
	var seats []model.Seats

	stmt := mysql.SELECT(
		table.Seats.ID,
		table.Seats.RowNr,
		table.Seats.ColumnNr,
		table.Seats.SeatCategoryID,
		table.Seats.CinemaHallID,
		table.Seats.Type,
	).FROM(
		table.Seats,
	).WHERE(
		table.Seats.CinemaHallID.EQ(utils.MysqlUuid(cinemaHallId)),
	)
	err := stmt.Query(tr.DatabaseManager.GetDatabaseConnection(), &seats)
	
	if err != nil {
		if err.Error() == "jet: sql: no rows in result set" {
			return nil, kts_errors.KTS_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return seats, nil
}

func (tr *TheatreRepository) CreateAddress(address model.Addresses) *models.KTSError {
	insertStmt := table.Addresses.INSERT(
		table.Addresses.ID,
		table.Addresses.Street,
		table.Addresses.StreetNr,
		table.Addresses.Zipcode,
		table.Addresses.City,
		table.Addresses.Country,
	).MODEL(address)
	_, err := insertStmt.Exec(tr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	return nil
}
