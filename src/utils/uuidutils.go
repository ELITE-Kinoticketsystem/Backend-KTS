package utils

import "github.com/google/uuid"

func NewUUID() *uuid.UUID {
	uuid := uuid.New()
	return &uuid
}
