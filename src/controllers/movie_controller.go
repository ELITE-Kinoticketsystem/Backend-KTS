package controllers

import (
	"log"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/repositories"
	"github.com/google/uuid"
)

type MovieControllerI interface {
	// Movie
	GetMovies() (*[]model.Movies, *models.KTSError)
	GetMovieById(movieId *uuid.UUID) (*models.MovieWithEverything, *models.KTSError)
	GetMovieByName(name *string) (*model.Movies, *models.KTSError)
	CreateMovie(movie *models.MovieDTOCreate) (*uuid.UUID, *models.KTSError)
	UpdateMovie(movie *model.Movies) *models.KTSError
	DeleteMovie(movieId *uuid.UUID) *models.KTSError

	// Combine Movie and Genre
	// AddMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError
	// RemoveMovieGenre(movieId *uuid.UUID, genreId *uuid.UUID) *models.KTSError

	// All Movies with all Genres - Grouped by Movie
	GetMoviesWithGenres() (*[]models.MovieWithGenres, *models.KTSError)
}

type MovieController struct {
	MovieRepo         repositories.MovieRepositoryI
	MovieGenreRepo    repositories.MovieGenreRepositoryI
	MovieActorRepo    repositories.MovieActorRepositoryI
	MovieProducerRepo repositories.MovieProducerRepositoryI
}

// Movie
func (mc *MovieController) GetMovies() (*[]model.Movies, *models.KTSError) {
	movies, kts_errors := mc.MovieRepo.GetMovies()
	if kts_errors != nil {
		return nil, kts_errors
	}
	return movies, nil
}

func (mc *MovieController) GetMovieByName(name *string) (*model.Movies, *models.KTSError) {
	movie, kts_errors := mc.MovieRepo.GetMovieByName(name)
	if kts_errors != nil {
		return nil, kts_errors
	}
	return movie, nil
}

func (mc *MovieController) CreateMovie(movie *models.MovieDTOCreate) (*uuid.UUID, *models.KTSError) {
	if movie.Movies == (model.Movies{}) {
		log.Print("Movie is nil")
		return nil, kts_errors.KTS_BAD_REQUEST
	}

	// Movie
	movieId, kts_errors := mc.MovieRepo.CreateMovie(&movie.Movies)
	if kts_errors != nil {
		log.Print("Movie was not created")
		return nil, kts_errors
	}

	// Add genre to movie
	movieGenres := movie.GenresID
	for _, movieGenre := range movieGenres {
		kts_err := mc.MovieGenreRepo.AddMovieGenre(movieId, movieGenre.ID)

		if kts_err != nil {
			log.Print("Genre was not added to movie")
			return nil, kts_err
		}
	}

	// Add actors to movie
	movieActors := movie.ActorsID
	log.Print("MovieActors: ", movieActors)
	for _, movieActor := range movieActors {
		kts_err := mc.MovieActorRepo.AddMovieActor(movieId, movieActor.ID)

		if kts_err != nil {
			log.Print("Actor was not added to movie")
			return nil, kts_err
		}
	}

	// Add producers to movie
	movieProducers := movie.ProducersID
	for _, movieProducer := range movieProducers {
		kts_err := mc.MovieProducerRepo.AddMovieProducer(movieId, movieProducer.ID)

		if kts_err != nil {
			log.Print("Producer was not added to movie")
			return nil, kts_err
		}
	}

	log.Print("Movie was created")
	return movieId, nil
}

func (mc *MovieController) UpdateMovie(movie *model.Movies) *models.KTSError {
	return mc.MovieRepo.UpdateMovie(movie)
}

func (mc *MovieController) DeleteMovie(movieId *uuid.UUID) *models.KTSError {
	// MovieGenre
	kts_errors := mc.MovieGenreRepo.RemoveAllGenreCombinationWithMovie(movieId)
	if kts_errors != nil {
		return kts_errors
	}

	// MovieActor
	kts_errors = mc.MovieActorRepo.RemoveAllActorCombinationWithMovie(movieId)
	if kts_errors != nil {
		return kts_errors
	}

	// MovieProducer
	kts_errors = mc.MovieProducerRepo.RemoveAllProducerCombinationWithMovie(movieId)
	if kts_errors != nil {
		return kts_errors
	}

	// Delete Movie
	kts_errors = mc.MovieRepo.DeleteMovie(movieId)
	if kts_errors != nil {
		return kts_errors
	}

	return nil
}

// All Movies with all Genres - Grouped by Movie
func (mc *MovieController) GetMoviesWithGenres() (*[]models.MovieWithGenres, *models.KTSError) {
	movies, kts_errors := mc.MovieRepo.GetMoviesWithGenres()
	if kts_errors != nil {
		return nil, kts_errors
	}
	return movies, nil
}

func (mc *MovieController) GetMovieById(movieId *uuid.UUID) (*models.MovieWithEverything, *models.KTSError) {
	movie, kts_errors := mc.MovieRepo.GetMovieById(movieId)
	if kts_errors != nil {
		return nil, kts_errors
	}
	return movie, nil
}
