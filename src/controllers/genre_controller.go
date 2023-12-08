package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
)

type GenreControllerI interface {
	GetGenres() (*[]model.Genres, *models.KTSError)
	GetGenreByName(name string) (*model.Genres, *models.KTSError)
	CreateGenre(name string) *models.KTSError

	// One Genre with all Movies
	GetGenreByNameWithMovies(genreName string) (*models.GenreWithMovies, *models.KTSError)
	// All Movies with all Genres - Grouped by Genre
	GetGenresWithMovies() (*[]models.GenreWithMovies, *models.KTSError)
}

type GenreController struct {
	GenreRepo repositories.GenreRepositoryI
}

func (mc *GenreController) GetGenres() (*[]model.Genres, *models.KTSError) {
	genres, ktskts_errors := mc.GenreRepo.GetGenres()
	if ktskts_errors != nil {
		return nil, ktskts_errors
	}
	return genres, nil
}

func (mc *GenreController) GetGenreByName(name string) (*model.Genres, *models.KTSError) {
	genre, kts_errors := mc.GenreRepo.GetGenreByName(name)
	if kts_errors != nil {
		return nil, kts_errors
	}
	return genre, nil
}

func (mc *GenreController) CreateGenre(name string) *models.KTSError {
	kts_errors := mc.GenreRepo.CreateGenre(name)
	if kts_errors != nil {
		return kts_errors
	}
	return nil
}

// One Genre with all Movies
func (mc *GenreController) GetGenreByNameWithMovies(genreName string) (*models.GenreWithMovies, *models.KTSError) {
	genre, kts_errors := mc.GenreRepo.GetGenreByNameWithMovies(genreName)
	if kts_errors != nil {
		return nil, kts_errors
	}
	return genre, nil
}

// All Movies with all Genres - Grouped by Genre
func (mc *GenreController) GetGenresWithMovies() (*[]models.GenreWithMovies, *models.KTSError) {
	genres, kts_errors := mc.GenreRepo.GetGenresWithMovies()
	if kts_errors != nil {
		return nil, kts_errors
	}
	return genres, nil
}
