package schemas

import (
	"github.com/google/uuid"
)

type Order struct {
	Id              *uuid.UUID `json:"id"`
	TotalPrice      int        `json:"total"` // requires conversion
	IsPaid          bool       `json:"isPaid"`
	PaymentMethodId *uuid.UUID `json:"paymentMethodId"`
	UserId          *uuid.UUID `json:"userId"`
}

type PaymentMethod struct {
	Id         *uuid.UUID `json:"id"`
	MethodName string     `json:"method"`
}
