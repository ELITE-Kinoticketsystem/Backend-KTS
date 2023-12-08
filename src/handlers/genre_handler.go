package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
)

// Genre
func GetGenres(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		genres, kts_err := genreCtrl.GetGenres()
		if kts_err != nil {
			c.JSON(kts_err.Status, kts_err)
			return
		}
		c.JSON(http.StatusOK, genres)
	}
}

func GetGenreByName(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		genre, kts_err := genreCtrl.GetGenreByName(name)
		if kts_err != nil {
			c.JSON(kts_err.Status, kts_err)
			return
		}
		c.JSON(http.StatusOK, genre)
	}
}

func CreateGenre(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var genre model.Genres
		err := c.ShouldBindJSON(&genre)
		if err != nil || utils.ContainsEmptyString(genre.GenreName) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		kts_err := genreCtrl.CreateGenre(genre.GenreName)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusCreated, genre)
	}
}

func GetGenreByNameWithMovies(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		genreName := c.Param("genreName")

		genre, kts_err := genreCtrl.GetGenreByNameWithMovies(genreName)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, genre)
	}
}

func GetGenresWithMovies(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		genres, kts_err := genreCtrl.GetGenresWithMovies()
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, genres)
	}
}
