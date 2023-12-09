package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
)

type GenreControllerI interface {
	GetGenres() (*[]model.Genres, *models.KTSError)
	GetGenreByName(name *string) (*model.Genres, *models.KTSError)
	CreateGenre(name *string) *models.KTSError
	UpdateGenre(name *string) *models.KTSError
	DeleteGenre(name *string) *models.KTSError

	// One Genre with all Movies
	GetGenreByNameWithMovies(genreName *string) (*models.GenreWithMovies, *models.KTSError)
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

func (mc *GenreController) GetGenreByName(name *string) (*model.Genres, *models.KTSError) {
	genre, kts_errors := mc.GenreRepo.GetGenreByName(name)
	if kts_errors != nil {
		return nil, kts_errors
	}
	return genre, nil
}

func (mc *GenreController) CreateGenre(name *string) *models.KTSError {
	// TODO: implement
	return kts_errors.KTS_INTERNAL_ERROR
}

func (mc *GenreController) UpdateGenre(name *string) *models.KTSError {
	// TODO implement
	return kts_errors.KTS_INTERNAL_ERROR
}

func (mc *GenreController) DeleteGenre(name *string) *models.KTSError {
	// TODO implement
	return kts_errors.KTS_INTERNAL_ERROR
}

// One Genre with all Movies
func (mc *GenreController) GetGenreByNameWithMovies(genreName *string) (*models.GenreWithMovies, *models.KTSError) {
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
