//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Events struct {
	ID           []byte `sql:"primary_key"`
	Title        string
	Start        time.Time
	End          time.Time
	EventType    string
	CinemaHallID []byte
}