package repositories

import (
	"encoding/json"
	"errors"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTicketById(t *testing.T) {
	sampleTicket := utils.GetSampleTicket()

	prettyTicket, err := json.Marshal(sampleTicket)
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	log.Print(string(prettyTicket))

	ticketId := sampleTicket.ID
	log.Print(ticketId)

	query := "SELECT tickets.id AS \"tickets.id\", tickets.validated AS \"tickets.validated\", tickets.price AS \"tickets.price\", seats.id AS \"seats.id\", seats.row_nr AS \"seats.row_nr\", seats.column_nr AS \"seats.column_nr\", seats.seat_category_id AS \"seats.seat_category_id\", seats.cinema_hall_id AS \"seats.cinema_hall_id\", seats.type AS \"seats.type\", events.id AS \"events.id\", events.title AS \"events.title\", events.start AS \"events.start\", events.end AS \"events.end\", events.description AS \"events.description\", events.event_type AS \"events.event_type\", events.cinema_hall_id AS \"events.cinema_hall_id\", orders.id AS \"orders.id\", orders.totalprice AS \"orders.totalprice\", orders.is_paid AS \"orders.is_paid\", orders.payment_method_id AS \"orders.payment_method_id\", orders.user_id AS \"orders.user_id\" FROM `KinoTicketSystem`.tickets INNER JOIN `KinoTicketSystem`.event_seats ON (event_seats.id = tickets.event_seat_id) INNER JOIN `KinoTicketSystem`.seats ON (seats.id = event_seats.seat_id) INNER JOIN `KinoTicketSystem`.events ON (events.id = event_seats.event_id) INNER JOIN `KinoTicketSystem`.orders ON (orders.id = tickets.order_id) WHERE tickets.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, id *uuid.UUID)
		expectedTicket  *models.TicketDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Ticket found",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(ticketId)).WillReturnRows(
					sqlmock.NewRows(
						[]string{"ticket.id", "ticket.validated", "ticket.price", "seats.id", "seats.row_nr", "seats.column_nr", "seats.seat_category_id", "seats.cinema_hall_id", "seats.type", "events.id", "events.title", "events.start", "events.end", "events.description", "events.event_type", "events.cinema_hall_id", "orders.id", "orders.totalprice", "orders.is_paid", "orders.payment_method_id", "orders.user_id"},
					).AddRow(
						&sampleTicket.ID, &sampleTicket.Validated, &sampleTicket.Price, &sampleTicket.Seats.ID, &sampleTicket.Seats.RowNr, &sampleTicket.Seats.ColumnNr, &sampleTicket.Seats.SeatCategoryID, &sampleTicket.Seats.CinemaHallID, &sampleTicket.Seats.Type, &sampleTicket.Event.ID, &sampleTicket.Event.Title, &sampleTicket.Event.Start, &sampleTicket.Event.End, &sampleTicket.Event.Description, &sampleTicket.Event.EventType, &sampleTicket.Event.CinemaHallID, &sampleTicket.Order.ID, &sampleTicket.Order.Totalprice, &sampleTicket.Order.IsPaid, &sampleTicket.Order.PaymentMethodID, &sampleTicket.Order.UserID,
					),
				)
			},
			expectedTicket: sampleTicket,
			expectedError:  nil,
		},
		{
			name: "Ticket not found",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(ticketId)).WillReturnRows(
					sqlmock.NewRows([]string{"ticket.id", "ticket.validated", "ticket.price", "seats.id", "seats.row_nr", "seats.column_nr", "seats.seat_category_id", "seats.cinema_hall_id", "seats.type", "events.id", "events.title", "events.start", "events.end", "events.description", "events.event_type", "events.cinema_hall_id", "orders.id", "orders.totalprice", "orders.is_paid", "orders.payment_method_id", "orders.user_id"}),
				)
			},
			expectedTicket: nil,
			expectedError:  kts_errors.KTS_NOT_FOUND,
		},
		{
			name: "Error while querying Ticket",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectQuery(query).WithArgs(utils.EqUUID(ticketId)).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedTicket: nil,
			expectedError:  kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the TicketRepository with the mock database connection
			ticketRepo := TicketRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, ticketId)

			// WHEN
			// Call the method under test
			ticket, kts_err := ticketRepo.GetTicketById(ticketId)

			// THEN
			// Verify the results
			assert.Equal(t, tc.expectedTicket, ticket)
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestCreateTicket(t *testing.T) {
	sampleTicket := utils.GetSampleCreateTicket()

	query := "INSERT INTO `KinoTicketSystem`.tickets (id, validated, price, price_category_id, order_id, event_seat_id) VALUES (?, ?, ?, ?, ?, ?);"

	testCases := []struct {
		name             string
		setExpectations  func(mock sqlmock.Sqlmock, ticket *model.Tickets)
		expectedTicketID bool
		expectedError    *models.KTSError
	}{
		{
			name: "Ticket created",
			setExpectations: func(mock sqlmock.Sqlmock, ticket *model.Tickets) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), ticket.Validated, ticket.Price, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedTicketID: true,
			expectedError:    nil,
		},
		{
			name: "Error while creating Ticket",
			setExpectations: func(mock sqlmock.Sqlmock, ticket *model.Tickets) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), ticket.Validated, ticket.Price, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedTicketID: false,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, ticket *model.Tickets) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), ticket.Validated, ticket.Price, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedTicketID: false,
			expectedError:    kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "PriceCategory not found",
			setExpectations: func(mock sqlmock.Sqlmock, ticket *model.Tickets) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), ticket.Validated, ticket.Price, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedTicketID: false,
			expectedError:    kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the PriceCategoryRepository with the mock database connection
			ticketRepo := TicketRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, sampleTicket)

			// WHEN
			ticketID, kts_err := ticketRepo.CreateTicket(sampleTicket)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)
			if tc.expectedTicketID && ticketID == nil {
				t.Error("Expected ticket ID, got nil")
			}

			// Verify that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}
}

func TestValidateTicket(t *testing.T) {
	ticketID := utils.NewUUID()

	query := "UPDATE `KinoTicketSystem`.tickets SET validated = ? WHERE tickets.id = ?;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock, id *uuid.UUID)
		expectedError   *models.KTSError
	}{
		{
			name: "Validated Ticket successfully",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(true, utils.EqUUID(ticketID)).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while validating ticket",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(true, sqlmock.AnyArg()).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Error while converting rows affected",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(true, sqlmock.AnyArg()).WillReturnResult(
					sqlmock.NewErrorResult(errors.New("rows affected conversion did not work")),
				)
			},
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Ticket not found",
			setExpectations: func(mock sqlmock.Sqlmock, id *uuid.UUID) {
				mock.ExpectExec(query).WithArgs(true, sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: kts_errors.KTS_NOT_FOUND,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			// Create a new instance of the PriceCategoryRepository with the mock database connection
			ticketRepo := TicketRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, ticketID)

			// WHEN
			kts_err := ticketRepo.ValidateTicket(ticketID)

			// THEN
			assert.Equal(t, tc.expectedError, kts_err)

			// Verify that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
