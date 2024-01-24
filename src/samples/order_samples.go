package samples

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

func GetOrder(priceCategories *[]model.PriceCategories, eventSeats *[]models.GetSlectedSeatsDTO, paymentMethodId *uuid.UUID) *models.CreateOrderDTO {

	return &models.CreateOrderDTO{
		EventSeatPriceCategories: []struct {
			EventSeatId     *uuid.UUID `binding:"required"`
			PriceCategoryId *uuid.UUID `binding:"required"`
		}{
			{
				EventSeatId:     (*eventSeats)[0].EventSeat.ID,
				PriceCategoryId: (*priceCategories)[0].ID,
			},
			{
				EventSeatId:     (*eventSeats)[1].EventSeat.ID,
				PriceCategoryId: (*priceCategories)[1].ID,
			},
		},
		PaymentMethodID: paymentMethodId,
	}
}

func GetOrderDTO() *models.CreateOrderDTO {
	return &models.CreateOrderDTO{
		EventSeatPriceCategories: []struct {
			EventSeatId     *uuid.UUID `binding:"required"`
			PriceCategoryId *uuid.UUID `binding:"required"`
		}{
			{
				EventSeatId:     utils.NewUUID(),
				PriceCategoryId: utils.NewUUID(),
			},
		},
		PaymentMethodID: utils.NewUUID(),
	}
}

func GetGetOrderDto() *[]models.GetOrderDTO {
	orderId := utils.NewUUID()
	order2Id := utils.NewUUID()

	return &[]models.GetOrderDTO{
		{
			Order: model.Orders{
				ID:         orderId,
				Totalprice: 1000,
				IsPaid:     false,
			},
			Tickets: []struct {
				Ticket        model.Tickets
				PriceCategory model.PriceCategories
				Seat          model.Seats
				SeatCategory  model.SeatCategories
			}{
				{
					Ticket: model.Tickets{
						ID:      utils.NewUUID(),
						OrderID: orderId,
					},
					PriceCategory: model.PriceCategories{
						Price:        100,
						CategoryName: "Test Price Category",
						ID:           utils.NewUUID(),
					},
					Seat: model.Seats{
						ID: utils.NewUUID(),
					},
					SeatCategory: model.SeatCategories{
						CategoryName: "Test Seat Category",
						ID:           utils.NewUUID(),
					},
				},
			},
		},
		{
			Order: model.Orders{
				ID:         order2Id,
				Totalprice: 0,
				IsPaid:     false,
			},
			Tickets: []struct {
				Ticket        model.Tickets
				PriceCategory model.PriceCategories
				Seat          model.Seats
				SeatCategory  model.SeatCategories
			}{
				{
					Ticket: model.Tickets{
						ID:      utils.NewUUID(),
						OrderID: order2Id,
					},
					Seat: model.Seats{
						ID: utils.NewUUID(),
					},
				},
			},
		},
	}
}

func GetOrderSample() models.GetOrderDTO {
	return models.GetOrderDTO{
		Order: model.Orders{
			ID:         utils.NewUUID(),
			Totalprice: 1000,
			IsPaid:     false,
		},
		Event: model.Events{
			ID:    utils.NewUUID(),
			Title: "Test Event",
			Start: time.Now(),
			End:   time.Now(),
			Is3d:  false,
		},
		CinemaHall: model.CinemaHalls{
			ID:   utils.NewUUID(),
			Name: "Test Cinema Hall",
		},
		Theatre: model.Theatres{
			ID:   utils.NewUUID(),
			Name: "Test Theatre",
		},
		Movies: []model.Movies{
			{
				ID:          utils.NewUUID(),
				Title:       "Test Movie",
				Description: "Test Description",
			},
		},
		Tickets: []struct {
			Ticket        model.Tickets
			PriceCategory model.PriceCategories
			Seat          model.Seats
			SeatCategory  model.SeatCategories
		}{
			{
				Ticket: model.Tickets{
					ID: utils.NewUUID(),
				},
				PriceCategory: model.PriceCategories{
					Price:        100,
					CategoryName: "Test Price Category",
					ID:           utils.NewUUID(),
				},
				Seat: model.Seats{
					ID:       utils.NewUUID(),
					RowNr:    1,
					ColumnNr: 1,
				},
				SeatCategory: model.SeatCategories{
					CategoryName: "Test Seat Category",
					ID:           utils.NewUUID(),
				},
			},
		},
	}

}

func GetModelOrder() *model.Orders {
	return &model.Orders{
		ID:              utils.NewUUID(),
		Totalprice:      100,
		IsPaid:          false,
		PaymentMethodID: utils.NewUUID(),
		UserID:          utils.NewUUID(),
	}
}
