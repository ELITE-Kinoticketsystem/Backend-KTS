package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/controllers"
	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
)

func GetTotalVisitsHandler(statsCtrl controllers.StatsControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		filterBy := c.Param("filterBy")
		filterBy = strings.ToLower(filterBy)
		startTime, err := time.Parse("2006-01-02", c.Param("from"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		endTime, err := time.Parse("2006-01-02", c.Param("til"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		if startTime.After(endTime) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		if filterBy != "day" && filterBy != "month" && filterBy != "year" {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		totalVisits, kts_err := statsCtrl.GetTotalVisits(startTime, endTime, filterBy)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, totalVisits)
	}
}

func GetTotalVisitsForTheatreHandler(statsCtrl controllers.StatsControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		filterBy := c.Param("filterBy")
		filterBy = strings.ToLower(filterBy)
		startTime, err := time.Parse("2006-01-02", c.Param("from"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		endTime, err := time.Parse("2006-01-02", c.Param("til"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		theatreName := c.Param("theatreName")
		if utils.ContainsEmptyString(theatreName) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		if startTime.After(endTime) {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		if filterBy != "day" && filterBy != "month" && filterBy != "year" {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		totalVisits, kts_err := statsCtrl.GetTotalVisitsForTheatre(startTime, endTime, filterBy, theatreName)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, totalVisits)
	}
}

func GetOrdersForStatsHandler(statsCtrl controllers.StatsControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		orders, kts_err := statsCtrl.GetOrdersForStats()
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, orders)
	}
}
