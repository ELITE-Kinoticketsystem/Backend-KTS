package schemas

import "github.com/google/uuid"

type Movie struct {
	Id          *uuid.UUID `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ReleaseDate string     `json:"releaseDate"`
	TimeInMin   int        `json:"timeInMin"`
	FskId       *uuid.UUID `json:"fskId"`
}

type FSK struct {
	Id  *uuid.UUID `json:"id"`
	Age int        `json:"age"`
}
