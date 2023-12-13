package models

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
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
