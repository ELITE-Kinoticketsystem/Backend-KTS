package samples

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
)

func GetOrder(priceCategories *[]model.PriceCategories, eventSeats *[]models.GetSlectedSeatsDTO, paymentMethodId *myid.UUID) *models.CreateOrderDTO {

	return &models.CreateOrderDTO{
		EventSeatPriceCategories: []struct {
			EventSeatId     *myid.UUID
			PriceCategoryId *myid.UUID
		}{
			{
				EventSeatId:     &(*eventSeats)[0].EventSeat.ID,
				PriceCategoryId: &(*priceCategories)[0].ID,
			},
			{
				EventSeatId:     &(*eventSeats)[1].EventSeat.ID,
				PriceCategoryId: &(*priceCategories)[1].ID,
			},
		},
		PaymentMethodID: paymentMethodId,
	}
}

func GetOrderDTO() *models.CreateOrderDTO {
	return &models.CreateOrderDTO{
		EventSeatPriceCategories: []struct {
			EventSeatId     *myid.UUID
			PriceCategoryId *myid.UUID
		}{
			{
				EventSeatId:     myid.NewUUID(),
				PriceCategoryId: myid.NewUUID(),
			},
		},
		PaymentMethodID: nil,
	}
}

func GetGetOrderDto() *[]models.GetOrderDTO {
	orderId := myid.New()
	order2Id := myid.New()

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
						ID:      myid.New(),
						OrderID: orderId,
					},
					Seat: model.Seats{
						ID: myid.New(),
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
				Ticket model.Tickets
				Seat   model.Seats
			}{
				{
					Ticket: model.Tickets{
						ID:      myid.New(),
						OrderID: order2Id,
					},
					Seat: model.Seats{
						ID: myid.New(),
					},
				},
			},
		},
	}
}

func GetModelOrder() *model.Orders {
	return &model.Orders{
		ID:              myid.New(),
		Totalprice:      100,
		IsPaid:          false,
		PaymentMethodID: myid.NewUUID(),
		UserID:          myid.New(),
	}
}
