package schemas

import "github.com/google/uuid"

type Theatre struct {
	Id        *uuid.UUID `json:"id"`
	Name      string     `json:"name"`
	AddressId *uuid.UUID `json:"addressId"`
}
