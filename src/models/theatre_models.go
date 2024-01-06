package models

import "github.com/google/uuid"

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
	Seats    [][]struct {
		RowNr    int
		ColumnNr int
		Type     string
		Category string
	}
	TheatreId *uuid.UUID
}
