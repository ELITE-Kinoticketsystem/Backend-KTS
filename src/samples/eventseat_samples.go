package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
)

func GetEventSeatsDTO() *[]models.GetEventSeatsDTO {
	eventId := myid.NewUUID()

	return &[]models.GetEventSeatsDTO{
		{
			EventSeat: model.EventSeats{
				ID:           myid.New(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      *eventId,
				SeatID:       myid.New(),
			},
			Seat: model.Seats{
				ID:             myid.New(),
				RowNr:          1,
				ColumnNr:       1,
				SeatCategoryID: myid.New(),
			},
			SeatCategory: model.SeatCategories{
				ID:           myid.New(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				EventID:        *eventId,
				SeatCategoryID: myid.New(),
				Price:          100,
			},
		},
		{
			EventSeat: model.EventSeats{
				ID:           myid.New(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      *eventId,
				SeatID:       myid.New(),
			},
			Seat: model.Seats{
				ID:             myid.New(),
				RowNr:          1,
				ColumnNr:       2,
				SeatCategoryID: myid.New(),
			},
			SeatCategory: model.SeatCategories{
				ID:           myid.New(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				EventID:        *eventId,
				SeatCategoryID: myid.New(),
				Price:          100,
			},
		},
	}
}

func GetGetSlectedSeatsDTO() *[]models.GetSlectedSeatsDTO {
	eventId := myid.NewUUID()

	return &[]models.GetSlectedSeatsDTO{
		{
			EventSeat: model.EventSeats{
				ID:           myid.New(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      *eventId,
				SeatID:       myid.New(),
			},
			Seat: model.Seats{
				ID:             myid.New(),
				RowNr:          1,
				ColumnNr:       1,
				SeatCategoryID: myid.New(),
			},
			SeatCategory: model.SeatCategories{
				ID:           myid.New(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				EventID:        *eventId,
				SeatCategoryID: myid.New(),
				Price:          100,
			},
		},
		{
			EventSeat: model.EventSeats{
				ID:           myid.New(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      *eventId,
				SeatID:       myid.New(),
			},
			Seat: model.Seats{
				ID:             myid.New(),
				RowNr:          1,
				ColumnNr:       2,
				SeatCategoryID: myid.New(),
			},
			SeatCategory: model.SeatCategories{
				ID:           myid.New(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				EventID:        *eventId,
				SeatCategoryID: myid.New(),
				Price:          100,
			},
		},
	}
}
