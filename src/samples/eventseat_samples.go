package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
)


func GetGetSlectedSeatsDTO() *[]models.GetSlectedSeatsDTO {
	eventId := utils.NewUUID()

	return &[]models.GetSlectedSeatsDTO{
		{
			EventSeat: model.EventSeats{
				ID:           utils.NewUUID(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      eventId,
				SeatID:       utils.NewUUID(),
			},
			Seat: model.Seats{
				ID:             utils.NewUUID(),
				RowNr:          1,
				ColumnNr:       1,
				SeatCategoryID: utils.NewUUID(),
			},
			SeatCategory: model.SeatCategories{
				ID:           utils.NewUUID(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				EventID:        eventId,
				SeatCategoryID: utils.NewUUID(),
				Price:          100,
			},
		},
		{
			EventSeat: model.EventSeats{
				ID:           utils.NewUUID(),
				Booked:       false,
				BlockedUntil: nil,
				UserID:       nil,
				EventID:      eventId,
				SeatID:       utils.NewUUID(),
			},
			Seat: model.Seats{
				ID:             utils.NewUUID(),
				RowNr:          1,
				ColumnNr:       2,
				SeatCategoryID: utils.NewUUID(),
			},
			SeatCategory: model.SeatCategories{
				ID:           utils.NewUUID(),
				CategoryName: "standard",
			},
			EventSeatCategory: model.EventSeatCategories{
				EventID:        eventId,
				SeatCategoryID: utils.NewUUID(),
				Price:          100,
			},
		},
	}
}
