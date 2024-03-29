package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Create theatre
// @Description Create theatre
// @Tags Theatres
// @Accept  json
// @Produce  json
// @Param theatre body models.CreateTheatreRequest true "Theatre data"
// @Success 201
// @Failure 400 {object} models.KTSErrorMessage
// @Router /theatres [post]
func CreateTheatre(theatreCtrl controllers.TheatreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var theatreData models.CreateTheatreRequest
		err := c.ShouldBindJSON(&theatreData)
		if err != nil || utils.ContainsEmptyString(theatreData.Name, theatreData.Address.Street, theatreData.Address.StreetNr, theatreData.Address.Zipcode, theatreData.Address.City, theatreData.Address.Country) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		kts_err := theatreCtrl.CreateTheatre(&theatreData)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusCreated, nil)
	}
}

// @Summary Get theatres
// @Description Get theatres
// @Tags Theatres
// @Accept  json
// @Produce  json
// @Param theatre body models.GetTheatreWithAddress true "Theatre data with addresses"
// @Success 200
// @Failure 400 {object} models.KTSErrorMessage
// @Router /theatres [get]
func GetTheatres(theatreCtrl controllers.TheatreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		theatres, kts_err := theatreCtrl.GetTheatres()
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, theatres)
	}
}

func CreateCinemaHallHandler(theatreCtrl controllers.TheatreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cinemaHallData models.CreateCinemaHallRequest
		err := c.ShouldBindJSON(&cinemaHallData)
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		if utils.ContainsEmptyString(cinemaHallData.HallName) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		kts_err := theatreCtrl.CreateCinemaHall(&cinemaHallData)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusCreated, nil)
	}
}

func GetCinemaHallsForTheatreHandler(theatreCtrl controllers.TheatreControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		theatreId, err := uuid.Parse(c.Param("theatreId"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}
		cinemaHalls, kts_err := theatreCtrl.GetCinemaHallsForTheatre(&theatreId)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, cinemaHalls)
	}
}