//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
)

type Seats struct {
	ID             *uuid.UUID `sql:"primary_key"`
	Y              int32
	X              int32
	RowNr          int32
	ColumnNr       int32
	SeatCategoryID *uuid.UUID
	CinemaHallID   *uuid.UUID
	Type           string
}
