package models

import (
	"time"

	"github.com/google/uuid"
)

type MovieResponse struct {
	Id          *uuid.UUID `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ReleaseDate time.Time  `json:"releaseDate"`
	TimeInMin   int        `json:"timeInMin"`
	Fsk         int        `json:"fsk"`
	GenreName   string     `json:"genreName"`
}