//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
)

type UserMovies struct {
	UserID   myid.UUID `sql:"primary_key"`
	MovieID  myid.UUID `sql:"primary_key"`
	ListType string    `sql:"primary_key"`
}
