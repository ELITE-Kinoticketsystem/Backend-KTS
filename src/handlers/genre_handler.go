package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/myid"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Get genres
// @Description Get genres
// @Tags Genres
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Genres
// @Failure 500 {object} models.KTSErrorMessage
// @Router /genres [get]
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

// @Summary Get genre by name
// @Description Get genre by name
// @Tags Genres
// @Accept  json
// @Produce  json
// @Param name path string true "Genre name"
// @Success 200 {object} model.Genres
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 404 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /genres/{name} [get]
func GetGenreByName(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if utils.ContainsEmptyString(name) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		genre, kts_err := genreCtrl.GetGenreByName(&name)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, genre)
	}
}

// @Summary Create genre
// @Description Create genre
// @Tags Genres
// @Accept  json
// @Produce  json
// @Param name path string true "Genre name"
// @Success 201 {object} model.Genres
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /genres/{name} [post]
func CreateGenre(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("name")
		if utils.ContainsEmptyString(name) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		genreId, kts_err := genreCtrl.CreateGenre(&name)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusCreated, genreId)
	}
}

// @Summary Update genre
// @Description Update genre
// @Tags Genres
// @Accept  json
// @Produce  json
// @Param genre body model.Genres true "Genres model"
// @Success 200 {object} model.Genres
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /genres [put]
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

// @Summary Delete genre
// @Description Delete genre
// @Tags Genres
// @Accept  json
// @Produce  json
// @Param id path string true "Genre ID"
// @Success 200
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /genres/{id} [delete]
func DeleteGenre(genreCtrl controllers.GenreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		genreId, err := myid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		kts_err := genreCtrl.DeleteGenre(&genreId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

// @Summary Get genres with movies
// @Description Get genres with movies
// @Tags Genres
// @Accept  json
// @Produce  json
// @Success 200 {array} models.GenreWithMovies
// @Failure 500 {object} models.KTSErrorMessage
// @Router /genres/movies [get]
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
