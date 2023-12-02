//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Movies struct {
	ID           uuid.UUID `sql:"primary_key"`
	Title        string
	Description  string
	BannerPicURL *string
	CoverPicURL  *string
	TrailerURL   *string
	Rating       *float64
	ReleaseDate  time.Time
	TimeInMin    int32
	Fsk          int32
}
