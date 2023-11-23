package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	Id          *uuid.UUID `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ReleaseDate time.Time  `json:"releaseDate"`
	TimeInMin   int        `json:"timeInMin"`
	Fsk         int        `json:"fsk"`
}

type Genre struct {
	Id        *uuid.UUID `json:"id"`
	GenreName string     `json:"genreName"`
}

type MovieGenre struct {
	MovieId *uuid.UUID `json:"movieId"`
	GenreId *uuid.UUID `json:"genreId"`
}
