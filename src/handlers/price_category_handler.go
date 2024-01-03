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

func GetPriceCategoryByIdHandler(priceCategoriesController controllers.PriceCategoryControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		priceCategoryId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}
		priceCategory, kts_err := priceCategoriesController.GetPriceCategoryById(&priceCategoryId)
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, priceCategory)
	}
}

func GetPriceCategoriesHandler(priceCategoriesController controllers.PriceCategoryControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		priceCategories, err := priceCategoriesController.GetPriceCategories()
		if err != nil {
			utils.HandleErrorAndAbort(c, err)
			return
		}

		c.JSON(http.StatusOK, priceCategories)
	}
}

func CreatePriceCategoryHandler(priceCategoriesController controllers.PriceCategoryControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var priceCategory model.PriceCategories
		if err := c.ShouldBindJSON(&priceCategory); err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		priceCategoryID, kts_err := priceCategoriesController.CreatePriceCategory(&priceCategory)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusCreated, priceCategoryID)
	}
}

func UpdatePriceCategoryHandler(priceCategoriesController controllers.PriceCategoryControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		var priceCategory model.PriceCategories
		if err := c.ShouldBindJSON(&priceCategory); err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		priceCategoryID, kts_err := priceCategoriesController.UpdatePriceCategory(&priceCategory)
		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusCreated, priceCategoryID)
	}
}

func DeletePriceCategoryHandler(priceCategoriesController controllers.PriceCategoryControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		priceCategoryId, err := uuid.Parse(c.Param("id"))
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}
		kts_err := priceCategoriesController.DeletePriceCategory(&priceCategoryId)
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, "Deleted")
	}
}
