package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/.gen/KinoTicketSystem/model"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
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

func GetMovieById(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := uuid.MustParse(c.Param("id"))
		movie, kts_err := movieCtrl.GetMovieById(&movieId)
		if kts_err != nil {
			c.JSON(kts_err.Status, kts_err)
			return
		}
		c.JSON(http.StatusOK, movie)
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
		var movie *model.Movies
		err := c.ShouldBindJSON(&movie)
		if err != nil ||
			utils.ContainsEmptyString(
				movie.Title, movie.Description, *movie.BannerPicURL, *movie.CoverPicURL, *movie.TrailerURL,
			) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		kts_err := movieCtrl.CreateMovie(movie)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusCreated, movie)
	}
}

func UpdateMovie(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie *model.Movies
		err := c.ShouldBindJSON(&movie)
		if err != nil ||
			utils.ContainsEmptyString(
				movie.Title, movie.Description, *movie.BannerPicURL, *movie.CoverPicURL, *movie.TrailerURL,
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
		movieId := uuid.MustParse(c.Param("movieId"))
		kts_err := movieCtrl.DeleteMovie(&movieId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.Status(http.StatusNoContent)
	}
}

// Combine Movie and Genre
// func AddMovieGenre(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		movieId := uuid.MustParse(c.Param("movieId"))
// 		genreId := uuid.MustParse(c.Param("genreId"))

// 		kts_err := movieCtrl.AddMovieGenre(&movieId, &genreId)
// 		if kts_err != nil {
// 			utils.HandleErrorAndAbort(c, kts_err)
// 			return
// 		}
// 		c.Status(http.StatusNoContent)
// 	}
// }

// func RemoveMovieGenre(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		movieId := uuid.MustParse(c.Param("movieId"))
// 		genreId := uuid.MustParse(c.Param("genreId"))

// 		kts_err := movieCtrl.RemoveMovieGenre(movieId, genreId)
// 		if kts_err != nil {
// 			utils.HandleErrorAndAbort(c, kts_err)
// 			return
// 		}
// 		c.Status(http.StatusNoContent)
// 	}
// }

func GetMovieByIdWithGenre(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := uuid.MustParse(c.Param("movieId"))

		movie, kts_err := movieCtrl.GetMovieByIdWithGenre(&movieId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, movie)
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

func GetMovieByIdWithEverything(movieCtrl controllers.MovieControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		movieId := uuid.MustParse(c.Param("id"))

		movie, kts_err := movieCtrl.GetMovieByIdWithEverything(&movieId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}
		c.JSON(http.StatusOK, movie)
	}
}
