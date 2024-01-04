package models

import (
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/google/uuid"
)

type MovieWithGenres struct {
	model.Movies

	Genres []struct {
		model.Genres
	}
}

type GenreWithMovies struct {
	model.Genres

	Movies []struct {
		model.Movies
	}
}

type MovieWithEverything struct {
	model.Movies

	Genres []struct {
		model.Genres
	}

	Actors []struct {
		model.Actors
	}

	Producers []struct {
		model.Producers
	}

	Reviews []struct {
		Review   model.Reviews
		Username string `alias:"users.username"`
	}
}

type MovieDTOCreate struct {
	model.Movies

	GenresID []struct {
		ID *uuid.UUID
	}

	ActorsID []struct {
		ID *uuid.UUID
	}

	ProducersID []struct {
		ID *uuid.UUID
	}
}

type CreateReviewRequest struct {
	Rating    int32
	Comment   string
	Datetime  time.Time
	IsSpoiler bool
	MovieID   string
}

type DeleteReviewRequest struct {
	ID string
}

type DeleteResponse struct {
	Message string `json:"message"`
}