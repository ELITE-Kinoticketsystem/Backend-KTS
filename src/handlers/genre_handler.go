package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetGenres(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		genres, kts_err := genreCtrl.GetGenres()
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, genres)
	}
}

func GetGenreByName(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		genre, kts_err := genreCtrl.GetGenreByName(&name)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
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

		genreId, kts_err := genreCtrl.CreateGenre(&genre.GenreName)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		genre.ID = genreId

		c.JSON(http.StatusCreated, genre)
	}
}

func UpdateGenre(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var genre model.Genres
		err := c.ShouldBindJSON(&genre)
		if err != nil || utils.ContainsEmptyString(genre.GenreName) || utils.ContainsEmptyString(genre.ID.String()) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		kts_err := genreCtrl.UpdateGenre(&genre)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, genre)
	}
}

func DeleteGenre(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		genreId := uuid.MustParse(id)

		kts_err := genreCtrl.DeleteGenre(&genreId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.Status(http.StatusOK)
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
