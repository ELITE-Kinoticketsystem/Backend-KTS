package repositories

import (
	"log"
	"time"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/table"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/go-jet/jet/v2/mysql"
)

type StatsRepositoryI interface {
	GetOrdersForStats() (*[]models.GetOrderDTO, *models.KTSError)
	GetTotalVisits(startTime time.Time, endTime time.Time, in string) (*[]models.StatsStruct, *models.KTSError)
}

type StatsRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (sr *StatsRepository) GetOrdersForStats() (*[]models.GetOrderDTO, *models.KTSError) {
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
		ORDER_BY(
			table.Events.Start.DESC(),
			table.Seats.RowNr.ASC(),
			table.Seats.ColumnNr.ASC(),
		)

	err := stmt.Query(sr.DatabaseManager.GetDatabaseConnection(), orders)

	if err != nil {
		log.Println(err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return orders, nil
}

func (sr *StatsRepository) GetTotalVisits(startTime time.Time, endTime time.Time, in string) (*[]models.StatsStruct, *models.KTSError) {
	visits := []models.StatsStruct{}

	filter := in + "(events.end)"

	stmt := mysql.SELECT(
		mysql.COUNT(table.Orders.ID),
		mysql.Raw(filter).AS("Duration"),
		mysql.Raw("MIN(events.end)"),
	).FROM(
		table.Orders.
			LEFT_JOIN(table.Tickets, table.Tickets.OrderID.EQ(table.Orders.ID)).
			LEFT_JOIN(table.EventSeats, table.EventSeats.ID.EQ(table.Tickets.EventSeatID)).
			LEFT_JOIN(table.Events, table.Events.ID.EQ(table.EventSeats.EventID)),
	).WHERE(
		table.Events.End.BETWEEN(utils.GetDateTime(startTime), utils.GetDateTime(endTime)),
	).GROUP_BY(
		mysql.Raw(filter),
	).ORDER_BY(
		mysql.Raw("MIN(events.end)"),
	)

	log.Print(stmt.DebugSql())

	err := stmt.Query(sr.DatabaseManager.GetDatabaseConnection(), &visits)

	if err != nil {
		log.Println(err)
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	log.Print(visits)

	return &visits, nil
}
