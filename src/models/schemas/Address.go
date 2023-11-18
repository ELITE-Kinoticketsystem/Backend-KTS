package schemas

import "github.com/google/uuid"

type Address struct {
	Id       *uuid.UUID `json:"id"`
	Street   string     `json:"street"`
	Streetnr string     `json:"streetnr"`
	Zipcode  string     `json:"zipcode"`
	City     string     `json:"city"`
	Country  string     `json:"country"`
}
