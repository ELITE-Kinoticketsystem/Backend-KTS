package repositories

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/samples"
	"github.com/stretchr/testify/assert"
)

func TestGetOrdersForStats(t *testing.T) {
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
				mock.ExpectQuery(query).WithArgs().WillReturnRows(sqlmock.NewRows([]string{
					"orders.id", "orders.totalprice", "orders.is_paid", "orders.payment_method_id", "orders.user_id", "tickets.id", "tickets.validated", "tickets.price", "tickets.price_category_id", "tickets.order_id", "tickets.event_seat_id", "price_categories.id", "price_categories.price", "price_categories.category_name", "seats.id", "seats.row_nr", "seats.column_nr", "seats.seat_category_id", "seats.cinema_hall_id", "seats.type", "seat_categories.id", "seat_categories.category_name"}).
					AddRow(orders[0].Order.ID, orders[0].Order.Totalprice, orders[0].Order.IsPaid, orders[0].Order.PaymentMethodID, orders[0].Order.UserID, orders[0].Tickets[0].Ticket.ID, orders[0].Tickets[0].Ticket.Validated, orders[0].Tickets[0].Ticket.Price, orders[0].Tickets[0].Ticket.PriceCategoryID, orders[0].Tickets[0].Ticket.OrderID, orders[0].Tickets[0].Ticket.EventSeatID, orders[0].Tickets[0].PriceCategory.ID, orders[0].Tickets[0].PriceCategory.Price, orders[0].Tickets[0].PriceCategory.CategoryName, orders[0].Tickets[0].Seat.ID, orders[0].Tickets[0].Seat.RowNr, orders[0].Tickets[0].Seat.ColumnNr, orders[0].Tickets[0].Seat.SeatCategoryID, orders[0].Tickets[0].Seat.CinemaHallID, orders[0].Tickets[0].Seat.Type, orders[0].Tickets[0].SeatCategory.ID, orders[0].Tickets[0].SeatCategory.CategoryName).
					AddRow(orders[1].Order.ID, orders[1].Order.Totalprice, orders[1].Order.IsPaid, orders[1].Order.PaymentMethodID, orders[1].Order.UserID, orders[1].Tickets[0].Ticket.ID, orders[1].Tickets[0].Ticket.Validated, orders[1].Tickets[0].Ticket.Price, orders[1].Tickets[0].Ticket.PriceCategoryID, orders[1].Tickets[0].Ticket.OrderID, orders[1].Tickets[0].Ticket.EventSeatID, orders[1].Tickets[0].PriceCategory.ID, orders[1].Tickets[0].PriceCategory.Price, orders[1].Tickets[0].PriceCategory.CategoryName, orders[1].Tickets[0].Seat.ID, orders[1].Tickets[0].Seat.RowNr, orders[1].Tickets[0].Seat.ColumnNr, orders[1].Tickets[0].Seat.SeatCategoryID, orders[1].Tickets[0].Seat.CinemaHallID, orders[1].Tickets[0].Seat.Type, orders[1].Tickets[0].SeatCategory.ID, orders[1].Tickets[0].SeatCategory.CategoryName))
			},
			expectOrder:   &orders,
			expectedError: nil,
		},
		{
			name: "Get order - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs().WillReturnError(sqlmock.ErrCancelled)
			},
			expectOrder:   nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			statsRepo := &StatsRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			orders, kts_err := statsRepo.GetOrdersForStats()

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

func TestGetTotalVisits(t *testing.T) {

	days := *samples.GetSampleDayVisitsStats()
	months := *samples.GetSampleMonthVisitsStats()
	years := *samples.GetSampleYearVisitsStats()
	startTime, _ := time.Parse("2006-01-01", "2019-01-01")
	endTime, _ := time.Parse("2006-01-01", "2024-01-01")

	testCases := []struct {
		name            string
		start           time.Time
		end             time.Time
		in              string
		setExpectations func(mock sqlmock.Sqlmock, startTime time.Time, endTime time.Time, in string)
		expectedStats   *[]models.StatsVisits
		expectedError   *models.KTSError
	}{
		{
			name:  "Get COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) in Days",
			start: startTime,
			end:   endTime,
			in:    "DAY",
			setExpectations: func(mock sqlmock.Sqlmock, startTime time.Time, endTime time.Time, in string) {
				mock.ExpectQuery(
					"SELECT COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) FROM `KinoTicketSystem`.tickets LEFT JOIN `KinoTicketSystem`.orders ON (orders.id = tickets.order_id) LEFT JOIN `KinoTicketSystem`.event_seats ON (event_seats.id = tickets.event_seat_id) LEFT JOIN `KinoTicketSystem`.events ON (events.id = event_seats.event_id) WHERE events.end BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME) GROUP BY DAY(events.end) ORDER BY MIN(events.end);",
				).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"COUNT(tickets.id)", "MIN(events.end)", "SUM(orders.totalprice)",
						},
					).
						AddRow(
							days[0].Count, days[0].Date, days[0].Revenue,
						).
						AddRow(
							days[1].Count, days[1].Date, days[1].Revenue,
						),
				)
			},
			expectedStats: &days,
			expectedError: nil,
		},
		{
			name:  "Get COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) in Months",
			start: startTime,
			end:   endTime,
			in:    "MONTH",
			setExpectations: func(mock sqlmock.Sqlmock, startTime time.Time, endTime time.Time, in string) {
				mock.ExpectQuery(
					"SELECT COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) FROM `KinoTicketSystem`.tickets LEFT JOIN `KinoTicketSystem`.orders ON (orders.id = tickets.order_id) LEFT JOIN `KinoTicketSystem`.event_seats ON (event_seats.id = tickets.event_seat_id) LEFT JOIN `KinoTicketSystem`.events ON (events.id = event_seats.event_id) WHERE events.end BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME) GROUP BY MONTH(events.end) ORDER BY MIN(events.end);",
				).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"COUNT(tickets.id)", "MIN(events.end)", "SUM(orders.totalprice)",
						},
					).
						AddRow(
							months[0].Count, months[0].Date, months[0].Revenue,
						).
						AddRow(
							months[1].Count, months[1].Date, months[1].Revenue,
						),
				)
			},
			expectedStats: &months,
			expectedError: nil,
		},
		{
			name:  "Get COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) in Years",
			start: startTime,
			end:   endTime,
			in:    "YEAR",
			setExpectations: func(mock sqlmock.Sqlmock, startTime time.Time, endTime time.Time, in string) {
				mock.ExpectQuery(
					"SELECT COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) FROM `KinoTicketSystem`.tickets LEFT JOIN `KinoTicketSystem`.orders ON (orders.id = tickets.order_id) LEFT JOIN `KinoTicketSystem`.event_seats ON (event_seats.id = tickets.event_seat_id) LEFT JOIN `KinoTicketSystem`.events ON (events.id = event_seats.event_id) WHERE events.end BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME) GROUP BY YEAR(events.end) ORDER BY MIN(events.end);",
				).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"COUNT(tickets.id)", "MIN(events.end)", "SUM(orders.totalprice)",
						},
					).
						AddRow(
							years[0].Count, years[0].Date, years[0].Revenue,
						).
						AddRow(
							years[1].Count, years[1].Date, years[1].Revenue,
						),
				)
			},
			expectedStats: &years,
			expectedError: nil,
		},
		{
			name:  "Get COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) in Days - Error",
			start: startTime,
			end:   endTime,
			in:    "DAY",
			setExpectations: func(mock sqlmock.Sqlmock, startTime time.Time, endTime time.Time, in string) {
				mock.ExpectQuery(
					"SELECT COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) FROM `KinoTicketSystem`.tickets LEFT JOIN `KinoTicketSystem`.orders ON (orders.id = tickets.order_id) LEFT JOIN `KinoTicketSystem`.event_seats ON (event_seats.id = tickets.event_seat_id) LEFT JOIN `KinoTicketSystem`.events ON (events.id = event_seats.event_id) WHERE events.end BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME) GROUP BY DAY(events.end) ORDER BY MIN(events.end);",
				).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedStats: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			statsRepo := &StatsRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, tc.start, tc.end, tc.in)

			totalVisits, kts_err := statsRepo.GetTotalVisits(tc.start, tc.end, tc.in)

			if kts_err != tc.expectedError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			assert.Equal(t, tc.expectedStats, totalVisits)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}

}

func TestGetTotalVisitsForTheatre(t *testing.T) {

	days := *samples.GetSampleDayVisitsStats()
	months := *samples.GetSampleMonthVisitsStats()
	years := *samples.GetSampleYearVisitsStats()
	startTime, _ := time.Parse("2006-01-01", "2019-01-01")
	endTime, _ := time.Parse("2006-01-01", "2024-01-01")

	testCases := []struct {
		name            string
		start           time.Time
		end             time.Time
		in              string
		theatreName     string
		setExpectations func(mock sqlmock.Sqlmock, startTime time.Time, endTime time.Time, in string, theatreName string)
		expectedStats   *[]models.StatsVisits
		expectedError   *models.KTSError
	}{
		{
			name:        "Get COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) in Days",
			start:       startTime,
			end:         endTime,
			in:          "DAY",
			theatreName: "Theatre1",
			setExpectations: func(mock sqlmock.Sqlmock, startTime time.Time, endTime time.Time, in string, theatreName string) {
				mock.ExpectQuery(
					"SELECT COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) FROM `KinoTicketSystem`.tickets LEFT JOIN `KinoTicketSystem`.orders ON (orders.id = tickets.order_id) LEFT JOIN `KinoTicketSystem`.event_seats ON (event_seats.id = tickets.event_seat_id) LEFT JOIN `KinoTicketSystem`.events ON (events.id = event_seats.event_id) LEFT JOIN `KinoTicketSystem`.cinema_halls ON (cinema_halls.id = events.cinema_hall_id) LEFT JOIN `KinoTicketSystem`.theatres ON (theatres.id = cinema_halls.theatre_id) WHERE (events.end BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME)) AND (theatres.name = ?) GROUP BY DAY(events.end) ORDER BY MIN(events.end);",
				).WithArgs(startTime, endTime, theatreName).WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"COUNT(tickets.id)", "MIN(events.end)", "SUM(orders.totalprice)",
						},
					).
						AddRow(
							days[0].Count, days[0].Date, days[0].Revenue,
						).
						AddRow(
							days[1].Count, days[1].Date, days[1].Revenue,
						),
				)
			},
			expectedStats: &days,
			expectedError: nil,
		},
		{
			name:        "Get COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) in Months",
			start:       startTime,
			end:         endTime,
			in:          "MONTH",
			theatreName: "Theatre1",
			setExpectations: func(mock sqlmock.Sqlmock, startTime time.Time, endTime time.Time, in string, theatreName string) {
				mock.ExpectQuery(
					"SELECT COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) FROM `KinoTicketSystem`.tickets LEFT JOIN `KinoTicketSystem`.orders ON (orders.id = tickets.order_id) LEFT JOIN `KinoTicketSystem`.event_seats ON (event_seats.id = tickets.event_seat_id) LEFT JOIN `KinoTicketSystem`.events ON (events.id = event_seats.event_id) LEFT JOIN `KinoTicketSystem`.cinema_halls ON (cinema_halls.id = events.cinema_hall_id) LEFT JOIN `KinoTicketSystem`.theatres ON (theatres.id = cinema_halls.theatre_id) WHERE (events.end BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME)) AND (theatres.name = ?) GROUP BY MONTH(events.end) ORDER BY MIN(events.end);",
				).WithArgs(startTime, endTime, theatreName).WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"COUNT(tickets.id)", "MIN(events.end)", "SUM(orders.totalprice)",
						},
					).
						AddRow(
							months[0].Count, months[0].Date, months[0].Revenue,
						).
						AddRow(
							months[1].Count, months[1].Date, months[1].Revenue,
						),
				)
			},
			expectedStats: &months,
			expectedError: nil,
		},
		{
			name:        "Get COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) in Years",
			start:       startTime,
			end:         endTime,
			in:          "YEAR",
			theatreName: "Theatre1",
			setExpectations: func(mock sqlmock.Sqlmock, startTime time.Time, endTime time.Time, in string, theatreName string) {
				mock.ExpectQuery(
					"SELECT COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) FROM `KinoTicketSystem`.tickets LEFT JOIN `KinoTicketSystem`.orders ON (orders.id = tickets.order_id) LEFT JOIN `KinoTicketSystem`.event_seats ON (event_seats.id = tickets.event_seat_id) LEFT JOIN `KinoTicketSystem`.events ON (events.id = event_seats.event_id) LEFT JOIN `KinoTicketSystem`.cinema_halls ON (cinema_halls.id = events.cinema_hall_id) LEFT JOIN `KinoTicketSystem`.theatres ON (theatres.id = cinema_halls.theatre_id) WHERE (events.end BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME)) AND (theatres.name = ?) GROUP BY YEAR(events.end) ORDER BY MIN(events.end);",
				).WithArgs(startTime, endTime, theatreName).WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"COUNT(tickets.id)", "MIN(events.end)", "SUM(orders.totalprice)",
						},
					).
						AddRow(
							years[0].Count, years[0].Date, years[0].Revenue,
						).
						AddRow(
							years[1].Count, years[1].Date, years[1].Revenue,
						),
				)
			},
			expectedStats: &years,
			expectedError: nil,
		},
		{
			name:        "Get COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) in Days - Error",
			start:       startTime,
			end:         endTime,
			in:          "DAY",
			theatreName: "Theatre1",
			setExpectations: func(mock sqlmock.Sqlmock, startTime time.Time, endTime time.Time, in string, theatreName string) {
				mock.ExpectQuery(
					"SELECT COUNT(tickets.id), MIN(events.end), SUM(orders.totalprice) FROM `KinoTicketSystem`.tickets LEFT JOIN `KinoTicketSystem`.orders ON (orders.id = tickets.order_id) LEFT JOIN `KinoTicketSystem`.event_seats ON (event_seats.id = tickets.event_seat_id) LEFT JOIN `KinoTicketSystem`.events ON (events.id = event_seats.event_id) LEFT JOIN `KinoTicketSystem`.cinema_halls ON (cinema_halls.id = events.cinema_hall_id) LEFT JOIN `KinoTicketSystem`.theatres ON (theatres.id = cinema_halls.theatre_id) WHERE (events.end BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME)) AND (theatres.name = ?) GROUP BY DAY(events.end) ORDER BY MIN(events.end);",
				).WithArgs(startTime, endTime, theatreName).WillReturnError(sqlmock.ErrCancelled)
			},
			expectedStats: nil,
			expectedError: kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			statsRepo := &StatsRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock, tc.start, tc.end, tc.in, tc.theatreName)

			totalVisits, kts_err := statsRepo.GetTotalVisitsForTheatre(tc.start, tc.end, tc.in, tc.theatreName)

			if kts_err != tc.expectedError {
				t.Errorf("Unexpected error: %v", kts_err)
			}

			assert.Equal(t, tc.expectedStats, totalVisits)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}

}

//

func TestGetMoviesSortedByTicketAmount(t *testing.T) {
	moviesSortedByTickets := *samples.GetSampleEventWithTicketCount()

	query := "SELECT events.title AS \"events.title\", COUNT(tickets.id) FROM `KinoTicketSystem`.events INNER JOIN `KinoTicketSystem`.event_seats ON (event_seats.event_id = events.id) INNER JOIN `KinoTicketSystem`.tickets ON (tickets.event_seat_id = event_seats.id) GROUP BY events.title ORDER BY COUNT(tickets.id) DESC;"

	testCases := []struct {
		name                      string
		setExpectations           func(mock sqlmock.Sqlmock)
		expectMoviesByTicketCount *[]models.GetEventWithTicketCount
		expectedError             *models.KTSError
	}{
		{
			name: "Get order",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs().WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"events.title", "COUNT(tickets.id)",
						},
					).
						AddRow(
							moviesSortedByTickets[0].EventName, moviesSortedByTickets[0].TicketCount,
						).
						AddRow(
							moviesSortedByTickets[1].EventName, moviesSortedByTickets[1].TicketCount,
						),
				)
			},
			expectMoviesByTicketCount: &moviesSortedByTickets,
			expectedError:             nil,
		},
		{
			name: "Get order - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs().WillReturnError(sqlmock.ErrCancelled)
			},
			expectMoviesByTicketCount: nil,
			expectedError:             kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			statsRepo := &StatsRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			moviesSortedByTicket, kts_err := statsRepo.GetMoviesSortedByTicketAmount()

			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectMoviesByTicketCount, moviesSortedByTicket)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}

}

func TestGetAllEventsTitle(t *testing.T) {
	sampleAllEvents := *samples.GetSampleAllEvents()

	query := "SELECT events.title AS \"events.title\" FROM `KinoTicketSystem`.events GROUP BY events.title;"

	testCases := []struct {
		name            string
		setExpectations func(mock sqlmock.Sqlmock)
		expectAllEvents *[]models.GetEventsTitle
		expectedError   *models.KTSError
	}{
		{
			name: "Get order",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs().WillReturnRows(
					sqlmock.NewRows(
						[]string{
							"events.title",
						},
					).
						AddRow(
							sampleAllEvents[0].EventName,
						).
						AddRow(
							sampleAllEvents[1].EventName,
						).
						AddRow(
							sampleAllEvents[2].EventName,
						),
				)
			},
			expectAllEvents: &sampleAllEvents,
			expectedError:   nil,
		},
		{
			name: "Get order - error",
			setExpectations: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(query).WithArgs().WillReturnError(sqlmock.ErrCancelled)
			},
			expectAllEvents: nil,
			expectedError:   kts_errors.KTS_INTERNAL_ERROR,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock database connection: %v", err)
			}
			defer db.Close()

			statsRepo := &StatsRepository{
				DatabaseManager: &managers.DatabaseManager{
					Connection: db,
				},
			}

			tc.setExpectations(mock)

			allEvents, kts_err := statsRepo.GetAllEventsTitle()

			assert.Equal(t, tc.expectedError, kts_err)
			assert.Equal(t, tc.expectAllEvents, allEvents)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}

		})
	}

}
