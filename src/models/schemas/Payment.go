package schemas

import "github.com/google/uuid"

type PaymentMethod struct {
	Id     *uuid.UUID `json:"id"`
	Method string     `json:"method"`
}
