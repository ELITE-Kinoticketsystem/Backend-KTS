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
	BlockEventSeatIfAvailable(eventId *uuid.UUID, seatId *uuid.UUID, userId *uuid.UUID, blockedUntil *time.Time) *models.KTSError
	UnblockEventSeat(eventId *uuid.UUID, seatId *uuid.UUID, userId *uuid.UUID) *models.KTSError
	UpdateBlockedUntilTimeForUserEventSeats(eventId *uuid.UUID, userId *uuid.UUID, blockedUntil *time.Time) *models.KTSError
	GetSelectedSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[]models.GetEventSeatsDTO, *models.KTSError)
	UpdateEventSeat(eventSeat *model.EventSeats) *models.KTSError
}

type EventSeatRepository struct {
	DatabaseManager *managers.DatabaseManager
}

func (esr *EventSeatRepository) GetEventSeats(eventId *uuid.UUID) (*[]models.GetEventSeatsDTO, *models.KTSError) {
	eventSeats := []models.GetEventSeatsDTO{}

	stmt := mysql.SELECT(
		table.EventSeats.AllColumns,
		table.Seats.AllColumns,
		table.SeatCategories.AllColumns,
		table.EventSeatCategories.AllColumns,
	).FROM(table.EventSeats.
		LEFT_JOIN(table.Seats, table.EventSeats.SeatID.EQ(table.Seats.ID)).
		LEFT_JOIN(table.SeatCategories, table.Seats.SeatCategoryID.EQ(table.SeatCategories.ID)).
		LEFT_JOIN(table.EventSeatCategories, table.EventSeatCategories.EventID.EQ(table.EventSeats.EventID).AND(table.EventSeatCategories.SeatCategoryID.EQ(table.Seats.SeatCategoryID)))).
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId))).ORDER_BY(table.Seats.ColumnNr.ASC(), table.Seats.RowNr.ASC())

	err := stmt.Query(esr.DatabaseManager.GetDatabaseConnection(), &eventSeats)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(eventSeats) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &eventSeats, nil
}

func (esr *EventSeatRepository) BlockEventSeatIfAvailable(eventId *uuid.UUID, seatId *uuid.UUID, userId *uuid.UUID, blockedUntil *time.Time) *models.KTSError {

	stmt := table.EventSeats.UPDATE(table.EventSeats.BlockedUntil, table.EventSeats.UserID).
		SET(blockedUntil, userId).
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)).AND(table.EventSeats.SeatID.EQ(utils.MysqlUuid(seatId))).
			AND(table.EventSeats.Booked.EQ(mysql.Bool(false))).
			AND(table.EventSeats.BlockedUntil.IS_NULL().OR(table.EventSeats.BlockedUntil.LT(utils.MysqlTimeNow())).OR(table.EventSeats.UserID.IS_NULL().OR(table.EventSeats.UserID.EQ(utils.MysqlUuid(userId))))))

	result, err := stmt.Exec(esr.DatabaseManager.GetDatabaseConnection())

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return kts_errors.KTS_CONFLICT
	}

	return nil
}

func (esr *EventSeatRepository) UpdateBlockedUntilTimeForUserEventSeats(eventId *uuid.UUID, userId *uuid.UUID, blockedUntil *time.Time) *models.KTSError {
	stmt := table.EventSeats.UPDATE(table.EventSeats.BlockedUntil).
		SET(blockedUntil).
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)).AND(table.EventSeats.UserID.EQ(utils.MysqlUuid(userId))))

	_, err := stmt.Exec(esr.DatabaseManager.GetDatabaseConnection())

	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	return nil
}

func (esr *EventSeatRepository) UnblockEventSeat(eventId *uuid.UUID, seatId *uuid.UUID, userId *uuid.UUID) *models.KTSError {
	stmt := table.EventSeats.UPDATE(table.EventSeats.BlockedUntil, table.EventSeats.UserID).
		SET(nil, nil).
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)).AND(table.EventSeats.SeatID.EQ(utils.MysqlUuid(seatId))).
			AND(table.EventSeats.UserID.EQ(utils.MysqlUuid(userId))))

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

func (esr *EventSeatRepository) GetSelectedSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[]models.GetEventSeatsDTO, *models.KTSError) {
	selectedSeats := []models.GetEventSeatsDTO{}

	stmt := mysql.SELECT(
		table.EventSeats.AllColumns,
		table.Seats.AllColumns,
		table.SeatCategories.AllColumns,
		table.EventSeatCategories.AllColumns,
	).FROM(table.EventSeats.
		LEFT_JOIN(table.Seats, table.EventSeats.SeatID.EQ(table.Seats.ID)).
		LEFT_JOIN(table.SeatCategories, table.Seats.SeatCategoryID.EQ(table.SeatCategories.ID)).
		LEFT_JOIN(table.EventSeatCategories, table.EventSeatCategories.EventID.EQ(table.EventSeats.EventID).AND(table.EventSeatCategories.SeatCategoryID.EQ(table.Seats.SeatCategoryID)))).
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)).AND(table.EventSeats.Booked.IS_FALSE()).AND(table.EventSeats.BlockedUntil.GT(mysql.CURRENT_TIMESTAMP()).AND(table.EventSeats.UserID.EQ(utils.MysqlUuid(userId))))).
		ORDER_BY(table.Seats.ColumnNr.ASC(), table.Seats.RowNr.ASC())

	err := stmt.Query(esr.DatabaseManager.GetDatabaseConnection(), &selectedSeats)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(selectedSeats) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &selectedSeats, nil
}

func (esr *EventSeatRepository) UpdateEventSeat(eventSeat *model.EventSeats) *models.KTSError {
	stmt := table.EventSeats.UPDATE(table.EventSeats.AllColumns).
		SET(
			utils.MysqlUuid(eventSeat.ID),
			eventSeat.Booked,
			eventSeat.BlockedUntil,
			utils.MysqlUuid(eventSeat.UserID),
			utils.MysqlUuid(eventSeat.SeatID),
			utils.MysqlUuid(eventSeat.EventID),
		).
		WHERE(table.EventSeats.ID.EQ(utils.MysqlUuid(eventSeat.ID)))

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
