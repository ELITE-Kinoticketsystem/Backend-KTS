package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
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
		model.Reviews
	}
}

type MovieDTO struct {
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
