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

	query := "SELECT orders.id AS \"orders.id\", orders.totalprice AS \"orders.totalprice\", orders.is_paid AS \"orders.is_paid\", orders.payment_method_id AS \"orders.payment_method_id\", orders.user_id AS \"orders.user_id\", tickets.id AS \"tickets.id\", tickets.validated AS \"tickets.validated\", tickets.price AS \"tickets.price\", tickets.price_category_id AS \"tickets.price_category_id\", tickets.order_id AS \"tickets.order_id\", tickets.event_seat_id AS \"tickets.event_seat_id\", seats.id AS \"seats.id\", seats.y AS \"seats.y\", seats.x AS \"seats.x\", seats.row_nr AS \"seats.row_nr\", seats.column_nr AS \"seats.column_nr\", seats.seat_category_id AS \"seats.seat_category_id\", seats.cinema_hall_id AS \"seats.cinema_hall_id\", seats.type AS \"seats.type\", events.id AS \"events.id\", events.title AS \"events.title\", events.start AS \"events.start\", events.end AS \"events.end\", events.description AS \"events.description\", events.event_type AS \"events.event_type\", events.cinema_hall_id AS \"events.cinema_hall_id\", events.is3d AS \"events.is3d\", cinema_halls.id AS \"cinema_halls.id\", cinema_halls.name AS \"cinema_halls.name\", cinema_halls.capacity AS \"cinema_halls.capacity\", cinema_halls.theatre_id AS \"cinema_halls.theatre_id\", cinema_halls.width AS \"cinema_halls.width\", cinema_halls.height AS \"cinema_halls.height\", theatres.id AS \"theatres.id\", theatres.name AS \"theatres.name\", theatres.logo_url AS \"theatres.logo_url\", theatres.address_id AS \"theatres.address_id\", movies.id AS \"movies.id\", movies.title AS \"movies.title\", movies.description AS \"movies.description\", movies.banner_pic_url AS \"movies.banner_pic_url\", movies.cover_pic_url AS \"movies.cover_pic_url\", movies.trailer_url AS \"movies.trailer_url\", movies.rating AS \"movies.rating\", movies.release_date AS \"movies.release_date\", movies.time_in_min AS \"movies.time_in_min\", movies.fsk AS \"movies.fsk\" FROM `KinoTicketSystem`.orders LEFT JOIN `KinoTicketSystem`.tickets ON (tickets.order_id = orders.id) LEFT JOIN `KinoTicketSystem`.event_seats ON (event_seats.id = tickets.event_seat_id) LEFT JOIN `KinoTicketSystem`.seats ON (seats.id = event_seats.seat_id) LEFT JOIN `KinoTicketSystem`.events ON (events.id = event_seats.event_id) LEFT JOIN `KinoTicketSystem`.cinema_halls ON (cinema_halls.id = events.cinema_hall_id) LEFT JOIN `KinoTicketSystem`.theatres ON (theatres.id = cinema_halls.theatre_id) LEFT JOIN `KinoTicketSystem`.event_movies ON (event_movies.event_id = events.id) LEFT JOIN `KinoTicketSystem`.movies ON (movies.id = event_movies.movie_id) ORDER BY events.start DESC, seats.row_nr ASC, seats.column_nr ASC;"

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
				mock.ExpectQuery(query).WithArgs().WillReturnError(sqlmock.ErrCancelled)
			},
			expectOrder:   nil,
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
				).WithArgs().WillReturnError(sqlmock.ErrCancelled)
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
