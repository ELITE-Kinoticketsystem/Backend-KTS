package myid

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

func (u UUID) Value() (driver.Value, error) {
	return u.MarshalBinary()
}

func Parse(str string) (UUID, error) {
	uuid, err := uuid.Parse(str)
	if err != nil {
		return UUID{}, err
	}
	return UUID{uuid}, nil
}

func MustParse(str string) UUID {
	uuid := uuid.MustParse(str)
	return UUID{uuid}
}

func New() UUID {
	uuid := uuid.New()
	return UUID{uuid}
}

func NewUUID() *UUID {
	uuid := uuid.New()
	return &UUID{uuid}
}
