package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	Id        *uuid.UUID `json:"id"`
	Timestamp time.Time  `json:"timestamp"`
	Validated bool       `json:"validated"`
	Paid      bool       `json:"paid"`
	Reserved  bool       `json:"reserved"`
	Price     int        `json:"price"` // requires conversion
	SeatId    *uuid.UUID `json:"seatId"`
	PriceCategory PriceCategory `json:"priceCategory"` // should have own table
}

type PriceCategory string

const (
	StudentDiscount PriceCategory = "StudentDiscount"
	ChildDiscount   PriceCategory = "ChildDiscount"
	ElderlyDiscount PriceCategory = "ElderlyDiscount"
	RegularPrice    PriceCategory = "RegularPrice"
)
