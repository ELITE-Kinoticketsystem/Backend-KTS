package repositories

import (
	"log"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type OrderRepoI interface {
	CreateOrder(order *model.Orders) (*uuid.UUID, *models.KTSError)
	GetOrderById(orderId *uuid.UUID, userId *uuid.UUID) (*models.GetOrderDTO, *models.KTSError)
	GetOrders(userId *uuid.UUID) (*[]models.GetOrderDTO, *models.KTSError)
}

type OrderRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (or *OrderRepository) CreateOrder(order *model.Orders) (*uuid.UUID, *models.KTSError) {

	insertStmt := table.Orders.INSERT(table.Orders.AllColumns).
		VALUES(
			utils.MysqlUuid(order.ID),
			order.Totalprice,
			order.IsPaid,
			utils.MysqlUuidOrNil(order.PaymentMethodID),
			utils.MysqlUuid(order.UserID),
		)
	log.Println(insertStmt.DebugSql())
	rows, err := insertStmt.Exec(or.DatabaseManager.GetDatabaseConnection())

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	affectedRows, err := rows.RowsAffected()

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if affectedRows == 0 {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return order.ID, nil
}

func (or *OrderRepository) GetOrderById(orderId *uuid.UUID, userId *uuid.UUID) (*models.GetOrderDTO, *models.KTSError) {
	order := models.GetOrderDTO{}

	stmt := table.Orders.SELECT(
		table.Orders.AllColumns,
		table.Tickets.AllColumns,
		table.Seats.AllColumns,
		table.Events.AllColumns,
		table.CinemaHalls.AllColumns,
		table.Theatres.AllColumns,
		table.Movies.AllColumns,
	).
		FROM(table.Orders.
			LEFT_JOIN(table.Tickets, table.Tickets.OrderID.EQ(table.Orders.ID)).
			LEFT_JOIN(table.EventSeats, table.EventSeats.ID.EQ(table.Tickets.EventSeatID)).
			LEFT_JOIN(table.Seats, table.Seats.ID.EQ(table.EventSeats.SeatID)).
			LEFT_JOIN(table.Events, table.Events.ID.EQ(table.EventSeats.EventID)).
			LEFT_JOIN(table.CinemaHalls, table.CinemaHalls.ID.EQ(table.Events.CinemaHallID)).
			LEFT_JOIN(table.Theatres, table.Theatres.ID.EQ(table.CinemaHalls.TheatreID)).
			LEFT_JOIN(table.EventMovies, table.EventMovies.EventID.EQ(table.Events.ID)).
			LEFT_JOIN(table.Movies, table.Movies.ID.EQ(table.EventMovies.MovieID)),
		).
		WHERE(table.Orders.ID.EQ(utils.MysqlUuid(orderId)).AND(table.Orders.UserID.EQ(utils.MysqlUuid(userId))))

	err := stmt.Query(or.DatabaseManager.GetDatabaseConnection(), &order)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &order, nil
}

func (or *OrderRepository) GetOrders(userId *uuid.UUID) (*[]models.GetOrderDTO, *models.KTSError) {
	orders := &[]models.GetOrderDTO{}

	stmt := table.Orders.SELECT(
		table.Orders.AllColumns,
		table.Tickets.AllColumns,
		table.Seats.AllColumns,
		table.Events.AllColumns,
		table.CinemaHalls.AllColumns,
		table.Theatres.AllColumns,
		table.Movies.AllColumns,
	).
		FROM(table.Orders.
			LEFT_JOIN(table.Tickets, table.Tickets.OrderID.EQ(table.Orders.ID)).
			LEFT_JOIN(table.EventSeats, table.EventSeats.ID.EQ(table.Tickets.EventSeatID)).
			LEFT_JOIN(table.Seats, table.Seats.ID.EQ(table.EventSeats.SeatID)).
			LEFT_JOIN(table.Events, table.Events.ID.EQ(table.EventSeats.EventID)).
			LEFT_JOIN(table.CinemaHalls, table.CinemaHalls.ID.EQ(table.Events.CinemaHallID)).
			LEFT_JOIN(table.Theatres, table.Theatres.ID.EQ(table.CinemaHalls.TheatreID)).
			LEFT_JOIN(table.EventMovies, table.EventMovies.EventID.EQ(table.Events.ID)).
			LEFT_JOIN(table.Movies, table.Movies.ID.EQ(table.EventMovies.MovieID)),
		).
		WHERE(table.Orders.UserID.EQ(utils.MysqlUuid(userId)))

	err := stmt.Query(or.DatabaseManager.GetDatabaseConnection(), orders)

	if err != nil {
		log.Println(err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return orders, nil
}
