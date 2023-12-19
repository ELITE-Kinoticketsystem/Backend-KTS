package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type OrderControllerI interface {
	CreateOrder(CreateOrderDTO models.CreateOrderDTO, eventId *uuid.UUID, userId *uuid.UUID, isReservation bool) (*uuid.UUID, *models.KTSError)
}

type OrderController struct {
	OrderRepo         repositories.OrderRepoI
	EventSeatRepo     repositories.EventSeatRepoI
	PriceCategoryRepo repositories.PriceCategoryRepositoryI
	TicketRepo        repositories.TicketRepoI
}

func (oc *OrderController) CreateOrder(CreateOrderDTO models.CreateOrderDTO, eventId *uuid.UUID, userId *uuid.UUID, isReservation bool) (*uuid.UUID, *models.KTSError) {
	if isReservation && CreateOrderDTO.PaymentMethodID != nil {
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	// Get EventSeats for User
	slectedSeats, kts_err := oc.EventSeatRepo.GetSelectedSeats(eventId, userId)
	if kts_err != nil {
		return nil, kts_err
	}

	priceCategories, kts_err := oc.PriceCategoryRepo.GetPriceCategories()

	adultPriceCategory := getPriceCategoryByName(*priceCategories, utils.ADULT)

	if kts_err != nil {
		return nil, kts_err
	}

	orderId := utils.NewUUID()

	// Create Ticket Objects and calculate total price
	tickets := make([]model.Tickets, len(*slectedSeats))

	totalPrice := int32(0)

	for i, seat := range *slectedSeats {

		var priceCategory *model.PriceCategories

		for _, seatPriceCategory := range CreateOrderDTO.EventSeatPriceCategories {
			if seat.EventSeat.ID == seatPriceCategory.EventSeatId {
				priceCategory = getPriceCategoryById(*priceCategories, seatPriceCategory.PriceCategoryId)
				break
			}
		}

		if priceCategory == nil {
			priceCategory = adultPriceCategory
		}

		price := utils.CalculatePrice(seat.EventSeatCategory.Price, priceCategory.Price)

		totalPrice += price

		tickets[i] = model.Tickets{
			ID:              utils.NewUUID(),
			OrderID:         orderId,
			EventSeatID:     seat.EventSeat.ID,
			PriceCategoryID: priceCategory.ID,
			Price:           price,
			Validated:       false,
		}
	}

	// Create Order
	order := model.Orders{
		ID:              orderId,
		UserID:          userId,
		PaymentMethodID: CreateOrderDTO.PaymentMethodID,
		IsPaid:          !isReservation,
		Totalprice:      totalPrice,
	}

	oc.OrderRepo.CreateOrder(&order)

	// With order ID and EventSeats create Tickets

	for _, ticket := range tickets {
		_, err := oc.TicketRepo.CreateTicket(&ticket)
		if err != nil {
			return nil, err
		}
	}

	// Update EventSeats to booked
	for _, seat := range *slectedSeats {
		seat.EventSeat.Booked = true
		err := oc.EventSeatRepo.UpdateEventSeat(&seat.EventSeat)
		if err != nil {
			return nil, err
		}
	}

	return orderId, nil
}

func getPriceCategoryByName(priceCategories []model.PriceCategories, name utils.PriceCategories) *model.PriceCategories {
	for _, priceCategory := range priceCategories {
		if priceCategory.CategoryName == name.String() {
			return &priceCategory
		}
	}
	return nil
}

func getPriceCategoryById(priceCategories []model.PriceCategories, id *uuid.UUID) *model.PriceCategories {
	for _, priceCategory := range priceCategories {
		if priceCategory.ID == id {
			return &priceCategory
		}
	}
	return nil
}
