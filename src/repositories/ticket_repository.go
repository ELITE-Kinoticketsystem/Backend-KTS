package repositories

import (
	"encoding/json"
	"log"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"

	"github.com/go-jet/jet/v2/mysql"
)

type TicketRepositoryI interface {
	GetTicketById(id *uuid.UUID) (*model.Tickets, *models.KTSError)
	CreateTicket(ticket *model.Tickets) (*uuid.UUID, *models.KTSError)
	ValidateTicket(id *uuid.UUID) *models.KTSError
}

type TicketRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (tr *TicketRepository) GetTicketById(id *uuid.UUID) (*models.TicketDTO, *models.KTSError) {
	var ticket models.TicketDTO

	// Create the query
	stmt := mysql.SELECT(
		table.Tickets.ID.AS("tickets.id"),
		table.Tickets.Validated.AS("tickets.validated"),
		table.Tickets.Price.AS("tickets.price"),

		table.Seats.AllColumns,
		table.Events.AllColumns,
		table.Orders.AllColumns,
	).FROM(
		table.Tickets.
			INNER_JOIN(table.EventSeats, table.EventSeats.ID.EQ(table.Tickets.EventSeatID)).
			INNER_JOIN(table.Seats, table.Seats.ID.EQ(table.EventSeats.SeatID)).
			INNER_JOIN(table.Events, table.Events.ID.EQ(table.EventSeats.EventID)).
			INNER_JOIN(table.Orders, table.Orders.ID.EQ(table.Tickets.OrderID)),
	).WHERE(
		table.Tickets.ID.EQ(utils.MysqlUuid(id)),
	)

	// Execute the query
	err := stmt.Query(tr.DatabaseManager.GetDatabaseConnection(), &ticket)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, kts_errors.KTS_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	prettyTicket, err := json.Marshal(ticket)
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	log.Print(string(prettyTicket))

	return &ticket, nil
}

func (tr *TicketRepository) CreateTicket(ticket *model.Tickets) (*uuid.UUID, *models.KTSError) {
	ticket.ID = utils.NewUUID()
	// Create the query
	stmt := table.Tickets.INSERT(
		table.Tickets.AllColumns,
	).VALUES(
		utils.MysqlUuid(ticket.ID),
		ticket.Validated,
		ticket.Price,
		utils.MysqlUuid(ticket.PriceCategoryID),
		utils.MysqlUuid(ticket.OrderID),
		utils.MysqlUuid(ticket.EventSeatID),
	)

	log.Print(stmt.DebugSql())

	// Execute the query
	rows, err := stmt.Exec(tr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return ticket.ID, nil
}

func (tr *TicketRepository) ValidateTicket(id *uuid.UUID) *models.KTSError {
	// Create the query
	stmt := table.Tickets.UPDATE(
		table.Tickets.Validated,
	).SET(
		true,
	).WHERE(
		table.Tickets.ID.EQ(utils.MysqlUuid(id)),
	)

	log.Print(stmt.DebugSql())

	// Execute the query
	rows, err := stmt.Exec(tr.DatabaseManager.GetDatabaseConnection())
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return kts_errors.KTS_INTERNAL_ERROR
	}

	if rowsAffected == 0 {
		return kts_errors.KTS_NOT_FOUND
	}

	return nil
}
