package handlers

import (
	"net/http"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
)

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

		c.Status(http.StatusCreated)
	}
}
