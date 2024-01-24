package controllers

import (
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/google/uuid"
)

type OrderControllerI interface {
	CreateOrder(createOrderDTO models.CreateOrderDTO, eventId *uuid.UUID, userId *uuid.UUID, isReservation bool) (*uuid.UUID, *models.KTSError)
	GetOrderById(orderId *uuid.UUID, userId *uuid.UUID) (*models.GetOrderDTO, *models.KTSError)
	GetOrders(userId *uuid.UUID) (*[]models.GetOrderDTO, *models.KTSError)
}

type OrderController struct {
	OrderRepo         repositories.OrderRepoI
	EventSeatRepo     repositories.EventSeatRepoI
	PriceCategoryRepo repositories.PriceCategoryRepositoryI
	TicketRepo        repositories.TicketRepositoryI
	UserRepo          repositories.UserRepositoryI
	MailMgr           managers.MailMgr
}

func (oc *OrderController) CreateOrder(createOrderDTO models.CreateOrderDTO, eventId *uuid.UUID, userId *uuid.UUID, isReservation bool) (*uuid.UUID, *models.KTSError) {
	if (isReservation) != (createOrderDTO.PaymentMethodID == nil) {
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	selectedSeats, kts_err_selectedSeats := oc.EventSeatRepo.GetSelectedSeats(eventId, userId)
	if kts_err_selectedSeats != nil {
		return nil, kts_err_selectedSeats
	}

	priceCategories, kts_err_price_categories := oc.PriceCategoryRepo.GetPriceCategories()
	if kts_err_price_categories != nil {
		return nil, kts_err_price_categories
	}

	adultPriceCategory := getPriceCategoryByName(*priceCategories, utils.ADULT)
	orderId := utils.NewUUID()

	tickets, totalPrice, kts_err_create_ticket_and_calc := createTicketsAndCalculateTotalPrice(selectedSeats, createOrderDTO, priceCategories, adultPriceCategory, orderId)

	if kts_err_create_ticket_and_calc != nil {
		return nil, kts_err_create_ticket_and_calc
	}

	order := model.Orders{
		ID:              orderId,
		UserID:          userId,
		PaymentMethodID: createOrderDTO.PaymentMethodID,
		IsPaid:          !isReservation,
		Totalprice:      totalPrice,
	}

	_, kts_err_create_order := oc.OrderRepo.CreateOrder(&order)
	if kts_err_create_order != nil {
		return nil, kts_err_create_order
	}

	for _, ticket := range tickets {
		_, err_create_ticket := oc.TicketRepo.CreateTicket(&ticket)
		if err_create_ticket != nil {
			return nil, err_create_ticket
		}
	}

	for _, seat := range *selectedSeats {
		seat.EventSeat.Booked = true
		err_update_eventseat := oc.EventSeatRepo.UpdateEventSeat(&seat.EventSeat)
		if err_update_eventseat != nil {
			return nil, err_update_eventseat
		}
	}

	oc.sendOrderConfirmationMail(orderId, userId)

	return orderId, nil
}

func (oc *OrderController) sendOrderConfirmationMail(orderId *uuid.UUID, userId *uuid.UUID) {
	order, kts_err := oc.OrderRepo.GetOrderById(orderId, userId)
	if kts_err != nil {
		return
	}

	user, kts_err := oc.UserRepo.GetUserById(userId)
	if kts_err != nil {
		return
	}

	kts_err = oc.MailMgr.SendOrderConfirmationMail(user.Email, *order)
	if kts_err != nil {
		return
	}
}

func (oc *OrderController) GetOrderById(orderId *uuid.UUID, userId *uuid.UUID) (*models.GetOrderDTO, *models.KTSError) {
	return oc.OrderRepo.GetOrderById(orderId, userId)
}

func (oc *OrderController) GetOrders(userId *uuid.UUID) (*[]models.GetOrderDTO, *models.KTSError) {
	return oc.OrderRepo.GetOrders(userId)
}

func createTicketsAndCalculateTotalPrice(slectedSeats *[]models.GetSlectedSeatsDTO, createOrderDTO models.CreateOrderDTO, priceCategories *[]model.PriceCategories, adultPriceCategory *model.PriceCategories, orderId *uuid.UUID) ([]model.Tickets, int32, *models.KTSError) {
	tickets := make([]model.Tickets, len(*slectedSeats))

	totalPrice := int32(0)

	for i, seat := range *slectedSeats {

		var priceCategory *model.PriceCategories

		for _, seatPriceCategory := range createOrderDTO.EventSeatPriceCategories {
			if seat.EventSeat.ID == nil {
				return nil, 0, kts_errors.KTS_BAD_REQUEST
			}
			if *seat.EventSeat.ID == *seatPriceCategory.EventSeatId {
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
			ID:              utils.NewUUID(),
			OrderID:         orderId,
			EventSeatID:     seat.EventSeat.ID,
			PriceCategoryID: priceCategory.ID,
			Price:           price,
			Validated:       false,
		}
	}
	return tickets, totalPrice, nil
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
		if *priceCategory.ID == *id {
			return &priceCategory
		}
	}
	return nil
}
