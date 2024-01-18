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

// @Summary Get Total Visits
// @Description Get Total Visits
// @Tags Stats
// @Accept  json
// @Produce  json
// @Param filterBy path string true "Filter By"
// @Param from path string true "From"
// @Param til path string true "Til"
// @Success 200 {object} models.StatsVisitsTwoArrays
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /stats/visits/{filterBy}/{from}/{til} [get]
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

// @Summary Get Total Visits For Theatre
// @Description Get Total Visits For Theatre
// @Tags Stats
// @Accept  json
// @Produce  json
// @Param filterBy path string true "Filter By"
// @Param from path string true "From"
// @Param til path string true "Til"
// @Param theatreName path string true "Theatre Name"
// @Success 200 {object} models.StatsVisitsTwoArrays
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /stats/visits/{filterBy}/{from}/{til}/{theatreName} [get]
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

// @Summary Get All Orders
// @Description Get All Orders
// @Tags Stats
// @Accept  json
// @Produce  json
// @Success 200 {object} models.StatsVisitsTwoArrays
// @Failure 500 {object} models.KTSErrorMessage
// @Router /stats/orders/ [get]
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
