package repositories

import (
	"log"

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
		WHERE(table.EventSeats.EventID.EQ(utils.MysqlUuid(eventId)))

	err := stmt.Query(esr.DatabaseManager.GetDatabaseConnection(), &eventSeats)

	log.Println(stmt.DebugSql())

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(eventSeats) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &eventSeats, nil
}
