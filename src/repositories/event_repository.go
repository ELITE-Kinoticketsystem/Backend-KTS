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

const (
	showing = "showing"
	special = "special event"
)

type EventRepo interface {
	CreateEvent(event *model.Events) (*uuid.UUID, *models.KTSError)
	GetEventsForMovie(movieId *uuid.UUID) ([]*model.Events, *models.KTSError)

	AddEventMovie(eventId *uuid.UUID, movieId *uuid.UUID) *models.KTSError

	GetSpecialEvents() (*[]models.GetSpecialEventsDTO, *models.KTSError)
	CreateEventSeat(eventSeat *model.EventSeats) *models.KTSError

	CreateEventSeatCategory(eventSeatCategory *model.EventSeatCategories) *models.KTSError
	GetEventById(eventId *uuid.UUID) (*models.GetSpecialEventsDTO, *models.KTSError)
}

type EventRepository struct {
	DatabaseManager managers.DatabaseManagerI
}

func (er *EventRepository) CreateEvent(event *model.Events) (*uuid.UUID, *models.KTSError) {
	event.ID = utils.NewUUID()

	stmt := table.Events.INSERT(table.Events.AllColumns).VALUES(
		utils.MysqlUuid(event.ID),
		event.Title,
		event.Start,
		event.End,
		event.Description,
		event.EventType,
		utils.MysqlUuid(event.CinemaHallID),
	)

	kts_err := utils.ExcecuteInsertStatement(stmt, er.DatabaseManager.GetDatabaseConnection())

	if kts_err != nil {
		return nil, kts_err
	}

	return event.ID, nil
}

func (er *EventRepository) CreateEventSeatCategory(eventSeatCategory *model.EventSeatCategories) *models.KTSError {
	insertStmt := table.EventSeatCategories.INSERT(table.EventSeatCategories.AllColumns).VALUES(
		utils.MysqlUuid(eventSeatCategory.EventID),
		utils.MysqlUuid(eventSeatCategory.SeatCategoryID),
		eventSeatCategory.Price,
	)

	return utils.ExcecuteInsertStatement(insertStmt, er.DatabaseManager.GetDatabaseConnection())
}

func (er *EventRepository) AddEventMovie(eventId *uuid.UUID, movieId *uuid.UUID) *models.KTSError {
	insertStmt := table.EventMovies.INSERT(table.EventMovies.EventID, table.EventMovies.MovieID).VALUES(
		utils.MysqlUuid(eventId),
		utils.MysqlUuid(movieId),
	)

	return utils.ExcecuteInsertStatement(insertStmt, er.DatabaseManager.GetDatabaseConnection())
}

func (er *EventRepository) CreateEventSeat(eventSeat *model.EventSeats) *models.KTSError {
	insertStmt := table.EventSeats.INSERT(table.EventSeats.ID, table.EventSeats.EventID, table.EventSeats.SeatID, table.EventSeats.Booked).VALUES(
		utils.MysqlUuid(eventSeat.ID),
		utils.MysqlUuid(eventSeat.EventID),
		utils.MysqlUuid(eventSeat.SeatID),
		eventSeat.Booked,
	)

	return utils.ExcecuteInsertStatement(insertStmt, er.DatabaseManager.GetDatabaseConnection())
}

func (er *EventRepository) GetSpecialEvents() (*[]models.GetSpecialEventsDTO, *models.KTSError) {
	var specialEvents []models.GetSpecialEventsDTO

	stmt := mysql.SELECT(
		table.Events.AllColumns,
		table.Movies.AllColumns,
	).
		FROM(
			table.Events.
				LEFT_JOIN(table.EventMovies, table.EventMovies.EventID.EQ(table.Events.ID)).
				LEFT_JOIN(table.Movies, table.Movies.ID.EQ(table.EventMovies.MovieID)),
		).
		WHERE(
			table.Events.EventType.EQ(utils.MySqlString(special)).AND(table.Events.Start.GT(utils.MysqlTimeNow())),
		)

	err := stmt.Query(er.DatabaseManager.GetDatabaseConnection(), &specialEvents)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	if len(specialEvents) == 0 {
		return nil, kts_errors.KTS_NOT_FOUND
	}

	return &specialEvents, nil
}

func (er *EventRepository) GetEventsForMovie(movieId *uuid.UUID) ([]*model.Events, *models.KTSError) {
	stmt := table.Events.SELECT(table.Events.AllColumns).FROM(table.Events.LEFT_JOIN(table.EventMovies, table.EventMovies.EventID.EQ(table.Events.ID))).WHERE(table.EventMovies.MovieID.EQ(utils.MysqlUuid(movieId))).WHERE(table.Events.Start.GT(utils.MysqlTimeNow())).WHERE(table.Events.EventType.EQ(utils.MySqlString(showing)))

	var events []*model.Events

	err := stmt.Query(er.DatabaseManager.GetDatabaseConnection(), &events)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return events, nil
}

func (er *EventRepository) GetEventById(eventId *uuid.UUID) (*models.GetSpecialEventsDTO, *models.KTSError) {
	var event models.GetSpecialEventsDTO

	stmt := mysql.SELECT(
		table.Events.AllColumns,
		table.Movies.AllColumns,
	).
		FROM(
			table.Events.
				LEFT_JOIN(table.EventMovies, table.EventMovies.EventID.EQ(table.Events.ID)).
				LEFT_JOIN(table.Movies, table.Movies.ID.EQ(table.EventMovies.MovieID)),
		).
		WHERE(
			table.Events.ID.EQ(utils.MysqlUuid(eventId)),
		)

	err := stmt.Query(er.DatabaseManager.GetDatabaseConnection(), &event)

	if err != nil {
		return nil, kts_errors.KTS_INTERNAL_ERROR
	}

	return &event, nil
}
