package handlers

import (
	"log"
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMovies(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		movies, kts_err := movieCtrl.GetMovies()
		if kts_err != nil {
			c.JSON(kts_err.Status, kts_err)
			return
		}
		c.JSON(http.StatusOK, movies)
	}
}

func GetMovieByName(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		movie, kts_err := movieCtrl.GetMovieByName(&name)
		if kts_err != nil {
			c.JSON(kts_err.Status, kts_err)
			return
		}
		c.JSON(http.StatusOK, movie)
	}
}

func CreateMovie(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie models.MovieDTOCreate
		err := c.ShouldBindJSON(&movie)
		if err != nil ||
			utils.ContainsEmptyString(
				movie.ReleaseDate.String(),
			) {
			log.Print("Failed to bind JSON")
			log.Print(err)
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		log.Print(movie)

		for _, genre := range movie.GenresID {
			if utils.ContainsEmptyString(genre.ID.String()) {
				log.Print("Genre is empty")
				utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
				return
			}
		}

		for _, actor := range movie.ActorsID {
			if utils.ContainsEmptyString(actor.ID.String()) {
				log.Print("Actor is empty")
				utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
				return
			}
		}

		for _, producer := range movie.ProducersID {
			if utils.ContainsEmptyString(producer.ID.String()) {
				log.Print("Producer is empty")
				utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
				return
			}
		}

		log.Print("Passed validation")
		movieId, kts_err := movieCtrl.CreateMovie(&movie)
		if kts_err != nil {
			log.Print("Movie was not created")
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		log.Print("Movie was created")
		c.JSON(http.StatusCreated, movieId)
	}
}

func UpdateMovie(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie *model.Movies
		err := c.ShouldBindJSON(&movie)
		if err != nil ||
			utils.ContainsEmptyString(
				movie.Title,
			) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		kts_err := movieCtrl.UpdateMovie(movie)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, movie)
	}
}

func DeleteMovie(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId, err := uuid.Parse(c.Param("movieId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		kts_err := movieCtrl.DeleteMovie(&movieId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
	}
}

func GetMoviesWithGenres(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		movies, kts_err := movieCtrl.GetMoviesWithGenres()
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, movies)
	}
}

func GetMovieById(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := uuid.MustParse(c.Param("id"))

		movie, kts_err := movieCtrl.GetMovieById(&movieId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, movie)
	}
}
