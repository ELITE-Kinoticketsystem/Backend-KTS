package repositories

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
)

type EventSeatRepoI interface {
	GetEventSeats(eventId *uuid.UUID) (*[]models.GetEventSeatsDTO, *models.KTSError)
	GetEventSeat(eventSeatId *uuid.UUID) (*model.EventSeats, *models.KTSError)
	BlockEventSeat(eventSeatId *uuid.UUID, userId *uuid.UUID, blockedUntil *time.Time) *models.KTSError
	ResetTimerOnUserSeats(userId *uuid.UUID, currentTime *time.Time, blockedUntil *time.Time) *models.KTSError
}

type EventSeatRepository struct {
	DatabaseManager *managers.DatabaseManager
}

func (esr *EventSeatRepository) GetEventSeats(eventId *uuid.UUID) (*[]models.GetEventSeatsDTO, *models.KTSError) {
	eventSeats := []models.GetEventSeatsDTO{}

	stmt := mysql.SELECT(
		table.EventSeats.BlockedUntil.LT(utils.MysqlTimeNow()).AND(table.EventSeats.UserID.EQ(nil)).AND(table.EventSeats.Booked.EQ(mysql.Bool(false))).AS("available"),
		table.EventSeats.AllColumns).FROM(table.EventSeats.
		LEFT_JOIN(table.Seats, table.EventSeats.SeatID.EQ(table.Seats.ID)).
		LEFT_JOIN(table.SeatCategories, table.Seats.SeatCategoryID.EQ(table.SeatCategories.ID)),
	).WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)))

	err := stmt.Query(esr.DatabaseManager.GetDatabaseConnection(), &eventSeats)

	if err != nil {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &eventSeats, nil
}

func (esr *EventSeatRepository) GetEventSeat(eventSeatId *uuid.UUID) (*model.EventSeats, *models.KTSError) {
	eventSeat := model.EventSeats{}

	stmt := mysql.SELECT(
		table.EventSeats.AllColumns).FROM(table.EventSeats).WHERE(table.EventSeats.ID.EQ(utils.MysqlUuid(eventSeatId)))

	err := stmt.Query(esr.DatabaseManager.GetDatabaseConnection(), &eventSeat)

	if err != nil {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &eventSeat, nil
}

func (esr *EventSeatRepository) BlockEventSeat(eventSeatId *uuid.UUID, userId *uuid.UUID, blockedUntil *time.Time) *models.KTSError {

	stmt := table.EventSeats.UPDATE(table.EventSeats.BlockedUntil, table.EventSeats.UserID).SET(
		blockedUntil,
		utils.MysqlUuid(userId),
	).WHERE(table.EventSeats.ID.EQ(utils.MysqlUuid(eventSeatId))).WHERE(table.EventSeats.BlockedUntil.LT(utils.MysqlTimeNow())).WHERE(table.EventSeats.UserID.EQ(nil)).WHERE(table.EventSeats.Booked.EQ(mysql.Bool(false)))

	result, err := stmt.Exec(esr.DatabaseManager.GetDatabaseConnection())

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}

func (esr *EventSeatRepository) ResetTimerOnUserSeats(userId *uuid.UUID, currentTime *time.Time, blockedUntil *time.Time) *models.KTSError {

	stmt := table.EventSeats.UPDATE(table.EventSeats.BlockedUntil).SET(
		blockedUntil,
	).WHERE(table.EventSeats.UserID.EQ(utils.MysqlUuid(userId))).WHERE(table.EventSeats.BlockedUntil.GT(utils.MysqlTimeNow())).WHERE(table.EventSeats.Booked.EQ(mysql.Bool(false)))

	result, err := stmt.Exec(esr.DatabaseManager.GetDatabaseConnection())

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}
