package controllers

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type MovieControllerI interface {
	// Movie
	GetMovies() (*[]model.Movies, *models.KTSError)
	GetMovieById(movieId *uuid.UUID) (*model.Movies, *models.KTSError)
	GetMovieByName(name *string) (*model.Movies, *models.KTSError)
	CreateMovie(movie *model.Movies) *models.KTSError
	UpdateMovie(movie *model.Movies) *models.KTSError
	DeleteMovie(movieId *uuid.UUID) *models.KTSError

	// Combine Movie and Genre
	// AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError
	// RemoveMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError

	// One Movie with all Genres
	GetMovieByIdWithGenre(movieId *uuid.UUID) (*models.MovieWithGenres, *models.KTSError)

	// All Movies with all Genres - Grouped by Movie
	GetMoviesWithGenres() (*[]models.MovieWithGenres, *models.KTSError)

	GetMovieByIdWithEverything(movieId *uuid.UUID) (*models.MovieWithEverything, *models.KTSError)
}

type MovieController struct {
	MovieRepo      repositories.MovieRepositoryI
	MovieGenreRepo repositories.MovieGenreRepositoryI
}

// Movie
func (mc *MovieController) GetMovies() (*[]model.Movies, *models.KTSError) {
	movies, kts_errors := mc.MovieRepo.GetMovies()
	if kts_errors != nil {
		return nil, kts_errors
	}
	return movies, nil
}

func (mc *MovieController) GetMovieById(movieId *uuid.UUID) (*model.Movies, *models.KTSError) {
	movie, kts_errors := mc.MovieRepo.GetMovieById(movieId)
	if kts_errors != nil {
		return nil, kts_errors
	}
	return movie, nil
}

func (mc *MovieController) GetMovieByName(name *string) (*model.Movies, *models.KTSError) {
	movie, kts_errors := mc.MovieRepo.GetMovieByName(name)
	if kts_errors != nil {
		return nil, kts_errors
	}
	return movie, nil
}

func (mc *MovieController) CreateMovie(movie *model.Movies) *models.KTSError {
	// Add Method AddMovieGenre

	// TODO: implement
	return kts_errors.KTS_INTERNAL_ERROR
}

func (mc *MovieController) UpdateMovie(movie *model.Movies) *models.KTSError {
	// Add Method RemoveMovieGenre

	// TODO: implement
	return kts_errors.KTS_INTERNAL_ERROR
}

func (mc *MovieController) DeleteMovie(movieId *uuid.UUID) *models.KTSError {
	// TODO: implement
	return kts_errors.KTS_INTERNAL_ERROR
}

// One Movie with all Genres
func (mc *MovieController) GetMovieByIdWithGenre(movieId *uuid.UUID) (*models.MovieWithGenres, *models.KTSError) {
	movie, kts_errors := mc.MovieRepo.GetMovieByIdWithGenre(movieId)
	if kts_errors != nil {
		return nil, kts_errors
	}
	return movie, nil
}

// All Movies with all Genres - Grouped by Movie
func (mc *MovieController) GetMoviesWithGenres() (*[]models.MovieWithGenres, *models.KTSError) {
	movies, kts_errors := mc.MovieRepo.GetMoviesWithGenres()
	if kts_errors != nil {
		return nil, kts_errors
	}
	return movies, nil
}

func (mc *MovieController) GetMovieByIdWithEverything(movieId *uuid.UUID) (*models.MovieWithEverything, *models.KTSError) {
	movie, kts_errors := mc.MovieRepo.GetMovieByIdWithEverything(movieId)
	if kts_errors != nil {
		return nil, kts_errors
	}
	return movie, nil
}
