package models

import (
	"time"

	"github.com/google/uuid"
)

type MovieDTO struct {
	Id          *uuid.UUID `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ReleaseDate time.Time  `json:"releaseDate"`
	TimeInMin   int32      `json:"timeInMin"`
	Fsk         int32      `json:"fsk"`
	GenreNames  []string   `json:"genreName"`
}
