package repositories

import (
	"context"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/managers"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/google/uuid"
)

type TripRepo interface {
	GetMovies(ctx context.Context) []*models.Movie
	
	// CRUD
	GetMovieById(ctx context.Context, movieId *uuid.UUID) *models.Movie
	CreateMovie(ctx context.Context, movie *models.Movie) *models.Movie
	UpdateMovie(ctx context.Context, movie *models.Movie) *models.Movie
	DeleteMovie(ctx context.Context, movieId *uuid.UUID) *models.Movie

	// Validation
	ValidateIfMovieExists(ctx context.Context, movieId *uuid.UUID) bool

	GetActorsFromMovie(ctx context.Context, movieId *uuid.UUID) []*models.Actor
	GetProducerFromMovie(ctx context.Context, movieId *uuid.UUID) *models.Producer
}

type TripRepository struct {
	DatabaseMgr managers.DatabaseManagerI
}
