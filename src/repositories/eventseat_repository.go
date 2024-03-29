package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/go-jet/jet/v2/mysql"
	"github.com/google/uuid"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
)

type EventSeatRepoI interface {
	GetEventSeats(eventId *uuid.UUID) (*[]models.GetEventSeatsDTO, *models.KTSError)
	GetHallDimensions(eventId *uuid.UUID) (int32, int32, *models.KTSError)
	BlockEventSeatIfAvailable(eventId *uuid.UUID, seatId *uuid.UUID, userId *uuid.UUID, blockedUntil *time.Time) *models.KTSError
	UnblockEventSeat(eventId *uuid.UUID, seatId *uuid.UUID, userId *uuid.UUID) *models.KTSError
	UnblockAllEventSeats(eventId *uuid.UUID, userId *uuid.UUID) *models.KTSError
	UpdateBlockedUntilTimeForUserEventSeats(eventId *uuid.UUID, userId *uuid.UUID, blockedUntil *time.Time) (int64, *models.KTSError)
	GetSelectedSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[]models.GetSlectedSeatsDTO, *models.KTSError)
	UpdateEventSeat(eventSeat *model.EventSeats) *models.KTSError
}

type EventSeatRepository struct {
	managers.DatabaseManagerI
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
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId))).ORDER_BY(table.Seats.RowNr.ASC(), table.Seats.ColumnNr.ASC())

	err := stmt.Query(esr.GetDatabaseConnection(), &eventSeats)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(eventSeats) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &eventSeats, nil
}

func (esr *EventSeatRepository) GetHallDimensions(eventId *uuid.UUID) (int32, int32, *models.KTSError) {
	var width int32
	var height int32

	query := "SELECT cinema_halls.width, cinema_halls.height FROM events LEFT JOIN cinema_halls ON cinema_halls.id = events.cinema_hall_id WHERE events.id = ?"

	err := esr.GetDatabaseConnection().QueryRow(query, eventId[:]).Scan(&width, &height)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, kts_errors.KTS_NOT_FOUND
		}
		return 0, 0, kts_errors.KTS_INTERNAL_ERROR
	}

	return width, height, nil
}

func (esr *EventSeatRepository) BlockEventSeatIfAvailable(eventId *uuid.UUID, seatId *uuid.UUID, userId *uuid.UUID, blockedUntil *time.Time) *models.KTSError {

	stmt := table.EventSeats.UPDATE(table.EventSeats.BlockedUntil, table.EventSeats.UserID).
		SET(utils.MysqlTime(blockedUntil), utils.MysqlUuid(userId)).
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)).AND(table.EventSeats.ID.EQ(utils.MysqlUuid(seatId))).
			AND(table.EventSeats.Booked.EQ(mysql.Bool(false))).
			AND(table.EventSeats.BlockedUntil.IS_NULL().OR(table.EventSeats.BlockedUntil.LT(utils.MysqlTimeNow())).OR(table.EventSeats.UserID.IS_NULL().OR(table.EventSeats.UserID.EQ(utils.MysqlUuid(userId))))))

	result, kts_err := stmt.Exec(esr.GetDatabaseConnection())
	if kts_err != nil {
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

func (esr *EventSeatRepository) UpdateBlockedUntilTimeForUserEventSeats(eventId *uuid.UUID, userId *uuid.UUID, blockedUntil *time.Time) (int64, *models.KTSError) {
	stmt := table.EventSeats.UPDATE(table.EventSeats.BlockedUntil).
		SET(blockedUntil).
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)).AND(table.EventSeats.UserID.EQ(utils.MysqlUuid(userId))).AND(table.EventSeats.BlockedUntil.GT(utils.MysqlTimeNow())).AND(table.EventSeats.Booked.IS_FALSE()))

	result, kts_err := stmt.Exec(esr.GetDatabaseConnection())

	if kts_err != nil {
		return 0, kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, kts_errors.KTS_INTERNAL_ERROR
	}

	return rowsAffected, nil
}

func (esr *EventSeatRepository) UnblockEventSeat(eventId *uuid.UUID, seatId *uuid.UUID, userId *uuid.UUID) *models.KTSError {
	stmt := table.EventSeats.UPDATE(table.EventSeats.BlockedUntil, table.EventSeats.UserID).
		SET(nil, nil).
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)).AND(table.EventSeats.ID.EQ(utils.MysqlUuid(seatId))).
			AND(table.EventSeats.UserID.EQ(utils.MysqlUuid(userId))).AND(table.EventSeats.BlockedUntil.GT(utils.MysqlTimeNow())).AND(table.EventSeats.Booked.IS_FALSE()))

	result, err := stmt.Exec(esr.GetDatabaseConnection())

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

func (esr *EventSeatRepository) UnblockAllEventSeats(eventId *uuid.UUID, userId *uuid.UUID) *models.KTSError {
	stmt := table.EventSeats.UPDATE(table.EventSeats.BlockedUntil, table.EventSeats.UserID).
		SET(nil, nil).
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)).
			AND(table.EventSeats.UserID.EQ(utils.MysqlUuid(userId))).AND(table.EventSeats.BlockedUntil.GT(utils.MysqlTimeNow())).AND(table.EventSeats.Booked.IS_FALSE()))

	result, err := stmt.Exec(esr.GetDatabaseConnection())

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

func (esr *EventSeatRepository) GetSelectedSeats(eventId *uuid.UUID, userId *uuid.UUID) (*[]models.GetSlectedSeatsDTO, *models.KTSError) {
	selectedSeats := []models.GetSlectedSeatsDTO{}

	stmt := mysql.SELECT(
		table.EventSeats.AllColumns,
		table.Seats.AllColumns,
		table.SeatCategories.AllColumns,
		table.EventSeatCategories.AllColumns,
		table.CinemaHalls.AllColumns,
		table.Theatres.AllColumns,
	).FROM(table.EventSeats.
		LEFT_JOIN(table.Seats, table.EventSeats.SeatID.EQ(table.Seats.ID)).
		LEFT_JOIN(table.SeatCategories, table.Seats.SeatCategoryID.EQ(table.SeatCategories.ID)).
		LEFT_JOIN(table.EventSeatCategories, table.EventSeatCategories.EventID.EQ(table.EventSeats.EventID).AND(table.EventSeatCategories.SeatCategoryID.EQ(table.Seats.SeatCategoryID))).
		LEFT_JOIN(table.Events, table.Events.ID.EQ(table.EventSeats.EventID)).
		LEFT_JOIN(table.CinemaHalls, table.CinemaHalls.ID.EQ(table.Events.CinemaHallID)).
		LEFT_JOIN(table.Theatres, table.Theatres.ID.EQ(table.CinemaHalls.TheatreID)),
	).
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)).AND(table.EventSeats.Booked.IS_FALSE()).AND(table.EventSeats.BlockedUntil.GT(mysql.CURRENT_TIMESTAMP()).AND(table.EventSeats.UserID.EQ(utils.MysqlUuid(userId))))).
		ORDER_BY(table.Seats.RowNr.ASC(), table.Seats.ColumnNr.ASC())

	log.Println(stmt.DebugSql())

	err := stmt.Query(esr.GetDatabaseConnection(), &selectedSeats)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(selectedSeats) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &selectedSeats, nil
}

func (esr *EventSeatRepository) UpdateEventSeat(eventSeat *model.EventSeats) *models.KTSError {
	eventSeatBlockUnitl := *eventSeat.BlockedUntil

	stmt := table.EventSeats.UPDATE().
		SET(
			table.EventSeats.Booked.SET(mysql.Bool(eventSeat.Booked)),
			// Assuming BlockedUntil is a *time.Time
			table.EventSeats.BlockedUntil.SET(mysql.DateTimeT(eventSeatBlockUnitl)),
			table.EventSeats.UserID.SET(utils.MysqlUuid(eventSeat.UserID)),
			table.EventSeats.SeatID.SET(utils.MysqlUuid(eventSeat.SeatID)),
			table.EventSeats.EventID.SET(utils.MysqlUuid(eventSeat.EventID)),
		).WHERE(
		table.EventSeats.ID.EQ(utils.MysqlUuid(eventSeat.ID)),
	)

	result, kts_err := stmt.Exec(esr.GetDatabaseConnection())
	if kts_err != nil {
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
