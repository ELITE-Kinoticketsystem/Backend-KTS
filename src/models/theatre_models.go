package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/google/uuid"
)

type CreateTheatreRequest struct {
	Name    string
	LogoUrl string
	Address struct {
		Street   string
		StreetNr string
		Zipcode  string
		City     string
		Country  string
	}
}

type CreateCinemaHallRequest struct {
	HallName string
	Width    int
	Height   int
	Seats    []struct {
		X        int
		Y        int
		Type     string
		Category string
	}
	TheatreId *uuid.UUID
}

type GetTheatreWithAddress struct {
	ID      *uuid.UUID `sql:"primary_key" alias:"theatres.id"`
	Name    string     `alias:"theatres.name"`
	LogoUrl *string    `alias:"theatres.logo_url"`
	Address model.Addresses
}
