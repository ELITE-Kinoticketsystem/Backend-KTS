package samples

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
)

func GetSampleTicket() *models.TicketDTO {
	ticket := models.TicketDTO{}

	id := uuid.MustParse("b1fd6028-4421-4d24-9cb1-b4fb84f180f9")
	seatId := uuid.MustParse("ac78c852-9b7d-4f08-996e-d606a54d4f38")
	eventId := uuid.MustParse("7ef0b48d-0696-4269-a713-1dfaed8f3249")
	orderId := uuid.MustParse("94c917a5-45e1-4550-be40-3a2de4688951")
	cinemaHallId := uuid.MustParse("11eea0aa-cf4c-b23c-bc67-0242ac120003")
	userId := uuid.MustParse("25c666f0-ee2b-42e1-854b-fde3c412d758")
	paymentMethodId := uuid.MustParse("d0c71fb8-f08c-4957-90f2-95fabcdadc45")
	seatCategoryId := uuid.MustParse("ec33b774-7e96-4428-818c-8cad0827dee0")
	description := "Test Description"

	timeStart, _ := time.Parse("2006-01-02T15:04:05.999999-07:00", "2023-12-31T10:52:24.108196+01:00")
	timeEnd, _ := time.Parse("2006-01-02T15:04:05.999999-07:00", "2023-12-31T10:52:24.108196+01:00")

	ticket = models.TicketDTO{
		ID:        &id,
		Validated: *new(bool),
		Price:     10,
		Seats: &model.Seats{
			ID:             &seatId,
			RowNr:          1,
			ColumnNr:       1,
			SeatCategoryID: &seatCategoryId,
			CinemaHallID:   &cinemaHallId,
			Type:           "Test Type",
		},
		Event: &model.Events{
			ID:           &eventId,
			Title:        "Test Event",
			Start:        timeStart,
			End:          timeEnd,
			Description:  &description,
			EventType:    "Test EventType",
			CinemaHallID: &cinemaHallId,
		},
		Order: &model.Orders{
			ID:              &orderId,
			Totalprice:      1000,
			IsPaid:          false,
			PaymentMethodID: &paymentMethodId,
			UserID:          &userId,
		},
	}

	return &ticket
}

func GetSampleCreateTicket() *model.Tickets {
	ticket := model.Tickets{}

	id := uuid.New()
	eventSeatId := uuid.New()
	priceCategoryID := uuid.New()
	orderId := uuid.New()

	ticket = model.Tickets{
		ID:              &id,
		Validated:       *new(bool),
		Price:           10,
		OrderID:         &orderId,
		PriceCategoryID: &priceCategoryID,
		EventSeatID:     &eventSeatId,
	}

	return &ticket
}
