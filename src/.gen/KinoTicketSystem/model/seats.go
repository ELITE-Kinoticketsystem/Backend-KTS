//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Seats struct {
	ID             []byte `sql:"primary_key"`
	RowNr          int32
	ColumnNr       int32
	SeatCategoryID []byte
	CinemaHallID   []byte
}