//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"time"
)

type Producers struct {
	ID          myid.UUID `sql:"primary_key"`
	Name        string
	Birthdate   time.Time
	Description string
}
