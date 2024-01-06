package repositories

import (
	"log"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"
)

type TheaterRepoI interface {
	CreateTheatre(theatre model.Theatres) *models.KTSError
	GetTheatres() (*[]model.Theatres, *models.KTSError)
	CreateCinemaHall(cinemaHall model.CinemaHalls) *models.KTSError
	CreateSeat(seat model.Seats) *models.KTSError
	GetSeatCategories() ([]model.SeatCategories, *models.KTSError)
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
		table.Theatres.LogoURL,
		table.Theatres.AddressID,
	).VALUES(
		utils.MysqlUuid(theatre.ID),
		theatre.Name,
		utils.MySqlStringPtr(theatre.LogoURL),
		utils.MysqlUuid(theatre.AddressID),
	)

	_, err := stmt.Exec(tr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	return nil
}

func (tr *TheatreRepository) GetTheatres() (*[]model.Theatres, *models.KTSError) {
	var theatres []model.Theatres
	stmt := mysql.SELECT(
		table.Theatres.AllColumns,
	).FROM(table.Theatres)

	err := stmt.Query(tr.DatabaseManager.GetDatabaseConnection(), &theatres)

	if err != nil {
		log.Println(err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &theatres, nil
}

func (tr *TheatreRepository) CreateCinemaHall(cinemaHall model.CinemaHalls) *models.KTSError {
	stmt := table.CinemaHalls.INSERT(
		table.CinemaHalls.ID,
		table.CinemaHalls.Name,
		table.CinemaHalls.Capacity,
		table.CinemaHalls.TheatreID,
	).VALUES(
		utils.MysqlUuid(cinemaHall.ID),
		cinemaHall.Name,
		cinemaHall.Capacity,
		utils.MysqlUuid(cinemaHall.TheatreID),
	)

	_, err := stmt.Exec(tr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	return nil
}

func (tr *TheatreRepository) CreateSeat(seat model.Seats) *models.KTSError {
	stmt := table.Seats.INSERT(
		table.Seats.ID,
		table.Seats.RowNr,
		table.Seats.ColumnNr,
		table.Seats.VisibleRowNr,
		table.Seats.VisibleColumnNr,
		table.Seats.SeatCategoryID,
		table.Seats.CinemaHallID,
		table.Seats.Type,
	).VALUES(
		utils.MysqlUuid(seat.ID),
		seat.RowNr,
		seat.ColumnNr,
		seat.VisibleRowNr,
		seat.VisibleColumnNr,
		utils.MysqlUuid(seat.SeatCategoryID),
		utils.MysqlUuid(seat.CinemaHallID),
		seat.Type,
	)

	_, err := stmt.Exec(tr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

func (tr *TheatreRepository) GetSeatCategories() ([]model.SeatCategories, *models.KTSError) {
	var seatCategories []model.SeatCategories

	stmt := mysql.SELECT(
		table.SeatCategories.ID,
		table.SeatCategories.CategoryName,
	).FROM(
		table.SeatCategories,
	)

	err := stmt.Query(tr.DatabaseManager.GetDatabaseConnection(), &seatCategories)
	if err != nil {
		if err.Error() == "jet: sql: no rows in result set" {
			return nil, kts_errors.KTS_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return seatCategories, nil
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
	).VALUES(
		utils.MysqlUuid(address.ID),
		address.Street,
		address.StreetNr,
		address.Zipcode,
		address.City,
		address.Country,
	)
	_, err := insertStmt.Exec(tr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}
	return nil
}
