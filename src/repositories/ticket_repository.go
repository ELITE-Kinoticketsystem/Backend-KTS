package repositories

import (
	"log"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/go-jet/jet/v2/mysql"
)

type TicketRepositoryI interface {
	GetTicketById(id *myid.UUID) (*models.TicketDTO, *models.KTSError)
	CreateTicket(ticket *model.Tickets) (*myid.UUID, *models.KTSError)
	ValidateTicket(id *myid.UUID) *models.KTSError
}

type TicketRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (tr *TicketRepository) GetTicketById(id *myid.UUID) (*models.TicketDTO, *models.KTSError) {
	var ticket models.TicketDTO

	// Create the query
	stmt := mysql.SELECT(
		table.Tickets.ID,
		table.Tickets.Validated,
		table.Tickets.Price,

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
		table.Tickets.ID.EQ(utils.MysqlUuid(*id)),
	)

	// Execute the query
	err := stmt.Query(tr.DatabaseManager.GetDatabaseConnection(), &ticket)
	if err != nil {
		if err.Error() == "qrm: no rows in result set" {
			return nil, kts_errors.KTS_NOT_FOUND
		}
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &ticket, nil
}

func (tr *TicketRepository) CreateTicket(ticket *model.Tickets) (*myid.UUID, *models.KTSError) {
	ticket.ID = myid.New()
	// Create the query
	stmt := table.Tickets.INSERT(
		table.Tickets.AllColumns,
	).VALUES(
		ticket.ID,
		ticket.Validated,
		ticket.Price,
		ticket.PriceCategoryID,
		ticket.OrderID,
		ticket.EventSeatID,
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
		return nil, kts_errors.KTS_CONFLICT
	}

	return &ticket.ID, nil
}

func (tr *TicketRepository) ValidateTicket(id *myid.UUID) *models.KTSError {
	// Create the query
	stmt := table.Tickets.UPDATE(
		table.Tickets.Validated,
	).SET(
		true,
	).WHERE(
		table.Tickets.ID.EQ(utils.MysqlUuid(*id)),
	)

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
		return kts_errors.KTS_CONFLICT
	}

	return nil
}
