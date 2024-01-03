package models

import "github.com/google/uuid"

type IdResponse struct {
	Id *uuid.UUID `json:"id"`
}
