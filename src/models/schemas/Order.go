package schemas

import (
	//"time"

	"github.com/google/uuid"
)

type Order struct {
	Id              *uuid.UUID `json:"id"`
	Total           int        `json:"total"` // requires conversion
	TicketId        *uuid.UUID `json:"ticketId"`
	PaymentMethodId *uuid.UUID `json:"paymentMethodId"`
	Reservation     bool       `json:"reservation"` // orderType would be better
	Booking         bool       `json:"booking"`
}
