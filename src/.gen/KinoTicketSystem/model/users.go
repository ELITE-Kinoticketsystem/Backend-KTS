//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Users struct {
	ID        []byte `sql:"primary_key"`
	Username  *string
	Email     string
	Password  string
	Firstname *string
	Lastname  *string
}
