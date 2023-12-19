package models

import "github.com/google/uuid"

type CreateOrderDTO struct {
	EventSeatPriceCategories []struct {
		EventSeatId     *uuid.UUID
		PriceCategoryId *uuid.UUID
	}

	PaymentMethodID *uuid.UUID
}
