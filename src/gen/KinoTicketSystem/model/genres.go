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

type Genres struct {
	ID        myid.UUID `sql:"primary_key"`
	GenreName string
}
