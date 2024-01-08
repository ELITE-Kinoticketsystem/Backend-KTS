package controllers

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
)

type OrderControllerI interface {
	CreateOrder(createOrderDTO models.CreateOrderDTO, eventId *myid.UUID, userId *myid.UUID, isReservation bool) (*myid.UUID, *models.KTSError)
	GetOrderById(orderId *myid.UUID, userId *myid.UUID) (*models.GetOrderDTO, *models.KTSError)
	GetOrders(userId *myid.UUID) (*[]models.GetOrderDTO, *models.KTSError)
}

type OrderController struct {
	OrderRepo         repositories.OrderRepoI
	EventSeatRepo     repositories.EventSeatRepoI
	PriceCategoryRepo repositories.PriceCategoryRepositoryI
	TicketRepo        repositories.TicketRepositoryI
}

func (oc *OrderController) CreateOrder(createOrderDTO models.CreateOrderDTO, eventId *myid.UUID, userId *myid.UUID, isReservation bool) (*myid.UUID, *models.KTSError) {
	if (isReservation) != (createOrderDTO.PaymentMethodID == nil) {
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	selectedSeats, kts_err := oc.EventSeatRepo.GetSelectedSeats(eventId, userId)
	if kts_err != nil {
		return nil, kts_err
	}

	priceCategories, kts_err := oc.PriceCategoryRepo.GetPriceCategories()
	if kts_err != nil {
		return nil, kts_err
	}

	adultPriceCategory := getPriceCategoryByName(*priceCategories, utils.ADULT)
	orderId := myid.NewUUID()

	tickets, totalPrice := createTicketsAndCalculateTotalPrice(selectedSeats, createOrderDTO, priceCategories, adultPriceCategory, orderId)

	order := model.Orders{
		ID:              *orderId,
		UserID:          *userId,
		PaymentMethodID: createOrderDTO.PaymentMethodID,
		IsPaid:          !isReservation,
		Totalprice:      totalPrice,
	}

	_, kts_err = oc.OrderRepo.CreateOrder(&order)
	if kts_err != nil {
		return nil, kts_err
	}

	for _, ticket := range tickets {
		_, err := oc.TicketRepo.CreateTicket(&ticket)
		if err != nil {
			return nil, err
		}
	}

	for _, seat := range *selectedSeats {
		seat.EventSeat.Booked = true
		err := oc.EventSeatRepo.UpdateEventSeat(&seat.EventSeat)
		if err != nil {
			return nil, err
		}
	}

	return orderId, nil
}

func (oc *OrderController) GetOrderById(orderId *myid.UUID, userId *myid.UUID) (*models.GetOrderDTO, *models.KTSError) {
	return oc.OrderRepo.GetOrderById(orderId, userId)
}

func (oc *OrderController) GetOrders(userId *myid.UUID) (*[]models.GetOrderDTO, *models.KTSError) {
	return oc.OrderRepo.GetOrders(userId)
}

func createTicketsAndCalculateTotalPrice(slectedSeats *[]models.GetSlectedSeatsDTO, createOrderDTO models.CreateOrderDTO, priceCategories *[]model.PriceCategories, adultPriceCategory *model.PriceCategories, orderId *myid.UUID) ([]model.Tickets, int32) {
	tickets := make([]model.Tickets, len(*slectedSeats))

	totalPrice := int32(0)

	for i, seat := range *slectedSeats {

		var priceCategory *model.PriceCategories

		for _, seatPriceCategory := range createOrderDTO.EventSeatPriceCategories {
			if seat.EventSeat.ID == *seatPriceCategory.EventSeatId {
				priceCategory = getPriceCategoryById(*priceCategories, seatPriceCategory.PriceCategoryId)
				break
			}
		}

		if priceCategory == nil {
			priceCategory = adultPriceCategory
		}

		price := utils.CalculatePrice(seat.EventSeatCategory.Price, priceCategory.Price, seat.Seat.Type)

		totalPrice += price

		tickets[i] = model.Tickets{
			ID:              myid.New(),
			OrderID:         *orderId,
			EventSeatID:     seat.EventSeat.ID,
			PriceCategoryID: priceCategory.ID,
			Price:           price,
			Validated:       false,
		}
	}
	return tickets, totalPrice
}

func getPriceCategoryByName(priceCategories []model.PriceCategories, name utils.PriceCategories) *model.PriceCategories {
	for _, priceCategory := range priceCategories {
		if priceCategory.CategoryName == name.String() {
			return &priceCategory
		}
	}
	return nil
}

func getPriceCategoryById(priceCategories []model.PriceCategories, id *myid.UUID) *model.PriceCategories {
	for _, priceCategory := range priceCategories {
		if priceCategory.ID == *id {
			return &priceCategory
		}
	}
	return nil
}
