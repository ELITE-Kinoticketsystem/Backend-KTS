package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

func GetOrder(priceCategories *[]model.PriceCategories, eventSeats *[]models.GetEventSeatsDTO) *models.CreateOrderDTO {

	return &models.CreateOrderDTO{
		EventSeatPriceCategories: []struct {
			EventSeatId     *uuid.UUID
			PriceCategoryId *uuid.UUID
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
		PaymentMethodID: nil,
	}
}

func GetOrderDTO() *models.CreateOrderDTO {
	return &models.CreateOrderDTO{
		EventSeatPriceCategories: []struct {
			EventSeatId     *uuid.UUID
			PriceCategoryId *uuid.UUID
		}{
			{
				EventSeatId:     utils.NewUUID(),
				PriceCategoryId: utils.NewUUID(),
			},
		},
		PaymentMethodID: nil,
	}
}

func GetGetOrderDto() *[]models.GetOrderDTO {
	orderId := utils.NewUUID()

	return &[]models.GetOrderDTO{
		{
			Order: model.Orders{
				ID:         orderId,
				Totalprice: 0,
				IsPaid:     false,
			},
			Tickets: []struct {
				Ticket model.Tickets
				Seat   model.Seats
			}{
				{
					Ticket: model.Tickets{
						ID:      utils.NewUUID(),
						OrderID: orderId,
					},
					Seat: model.Seats{
						ID: utils.NewUUID(),
					},
				},
			},
		},
	}
}
