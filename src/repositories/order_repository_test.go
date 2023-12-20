package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	order := &model.Orders{
		ID:              utils.NewUUID(),
		Totalprice:      100,
		IsPaid:          false,
		PaymentMethodID: utils.NewUUID(),
		UserID:          utils.NewUUID(),
	}

	query := "INSERT INTO `KinoTicketSystem`.orders .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectOrderID   bool
		expectedError   *models.KTSError
	}{
		{
			name: "Create order",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), order.Totalprice, order.IsPaid, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectOrderID: true,
			expectedError: nil,
		},
		{
			name: "Create order - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), order.Totalprice, order.IsPaid, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("error"))
			},
			expectOrderID: false,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Create order - no rows affected",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(query).WithArgs(sqlmock.AnyArg(), order.Totalprice, order.IsPaid, sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectOrderID: false,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			orderRepo := &OrderRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			orderID, ktsErr := orderRepo.CreateOrder(order)

			if ktsErr != tc.expectedError {
				t.Errorf("Unexpected error: %v", ktsErr)
			}

			if tc.expectOrderID && orderID == nil {
				t.Error("Expected order ID, got nil")
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestGetOrderById(t *testing.T) {
	order := (*samples.GetGetOrderDto())[0]

	query := "SELECT .* FROM `KinoTicketSystem`.orders .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectOrder     *models.GetOrderDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Get order",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{
					"orders.id", "orders.totalprice", "orders.is_paid", "orders.payment_method_id", "orders.user_id", "tickets.id", "tickets.validated", "tickets.price", "tickets.price_category_id", "tickets.order_id", "tickets.event_seat_id", "seats.id", "seats.row_nr", "seats.column_nr", "seats.seat_category_id", "seats.cinema_hall_id", "seats.type"}).
					AddRow(order.Order.ID, order.Order.Totalprice, order.Order.IsPaid, order.Order.PaymentMethodID, order.Order.UserID, order.Tickets[0].Ticket.ID, order.Tickets[0].Ticket.Validated, order.Tickets[0].Ticket.Price, order.Tickets[0].Ticket.PriceCategoryID, order.Tickets[0].Ticket.OrderID, order.Tickets[0].Ticket.EventSeatID, order.Tickets[0].Seat.ID, order.Tickets[0].Seat.RowNr, order.Tickets[0].Seat.ColumnNr, order.Tickets[0].Seat.SeatCategoryID, order.Tickets[0].Seat.CinemaHallID, order.Tickets[0].Seat.Type))
			},
			expectOrder:   &order,
			expectedError: nil,
		},
		{
			name: "Get order - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.New("error"))
			},
			expectOrder:   nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
		{
			name: "Get order - no rows",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{}))
			},
			expectOrder:   nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			orderRepo := &OrderRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			order, kts_err := orderRepo.GetOrderById(utils.NewUUID(), utils.NewUUID())

			if kts_err != tc.expectedError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			assert.Equal(t, tc.expectOrder, order)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}

}

func TestGetOrders(t *testing.T) {
	orders := *samples.GetGetOrderDto()

	query := "SELECT .* FROM `KinoTicketSystem`.orders .*"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectOrder     *[]models.GetOrderDTO
		expectedError   *models.KTSError
	}{
		{
			name: "Get order",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{
					"orders.id", "orders.totalprice", "orders.is_paid", "orders.payment_method_id", "orders.user_id", "tickets.id", "tickets.validated", "tickets.price", "tickets.price_category_id", "tickets.order_id", "tickets.event_seat_id", "seats.id", "seats.row_nr", "seats.column_nr", "seats.seat_category_id", "seats.cinema_hall_id", "seats.type"}).
					AddRow(orders[0].Order.ID, orders[0].Order.Totalprice, orders[0].Order.IsPaid, orders[0].Order.PaymentMethodID, orders[0].Order.UserID, orders[0].Tickets[0].Ticket.ID, orders[0].Tickets[0].Ticket.Validated, orders[0].Tickets[0].Ticket.Price, orders[0].Tickets[0].Ticket.PriceCategoryID, orders[0].Tickets[0].Ticket.OrderID, orders[0].Tickets[0].Ticket.EventSeatID, orders[0].Tickets[0].Seat.ID, orders[0].Tickets[0].Seat.RowNr, orders[0].Tickets[0].Seat.ColumnNr, orders[0].Tickets[0].Seat.SeatCategoryID, orders[0].Tickets[0].Seat.CinemaHallID, orders[0].Tickets[0].Seat.Type).
					AddRow(orders[1].Order.ID, orders[1].Order.Totalprice, orders[1].Order.IsPaid, orders[1].Order.PaymentMethodID, orders[1].Order.UserID, orders[1].Tickets[0].Ticket.ID, orders[1].Tickets[0].Ticket.Validated, orders[1].Tickets[0].Ticket.Price, orders[1].Tickets[0].Ticket.PriceCategoryID, orders[1].Tickets[0].Ticket.OrderID, orders[1].Tickets[0].Ticket.EventSeatID, orders[1].Tickets[0].Seat.ID, orders[1].Tickets[0].Seat.RowNr, orders[1].Tickets[0].Seat.ColumnNr, orders[1].Tickets[0].Seat.SeatCategoryID, orders[1].Tickets[0].Seat.CinemaHallID, orders[1].Tickets[0].Seat.Type))
			},
			expectOrder:   &orders,
			expectedError: nil,
		},
		{
			name: "Get order - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs(sqlmock.AnyArg()).WillReturnError(errors.New("error"))
			},
			expectOrder:   nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			orderRepo := &OrderRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			orders, kts_err := orderRepo.GetOrders(utils.NewUUID())

			if kts_err != tc.expectedError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			assert.Equal(t, tc.expectOrder, orders)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}

}
