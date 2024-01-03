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

// @Summary Get price category by id
// @Description Get price category by id
// @Tags PriceCategories
// @Accept  json
// @Produce  json
// @Param id path string true "Price category ID"
// @Success 200 {object} model.PriceCategories
// @Failure 500 {object} models.KTSErrorMessage
// @Router /price-categories/{id} [get]
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

// @Summary Get price categories
// @Description Get price categories
// @Tags PriceCategories
// @Accept  json
// @Produce  json
// @Success 200 {array} model.PriceCategories
// @Failure 500 {object} models.KTSErrorMessage
// @Router /price-categories [get]
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

// @Summary Create price category
// @Description Create price category
// @Tags PriceCategories
// @Accept  json
// @Produce  json
// @Param priceCategory body model.PriceCategories true "Price category data"
// @Success 200 {object} uuid.UUID
// @Failure 500 {object} models.KTSErrorMessage
// @Router /price-categories [post]
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

// @Summary Update price category
// @Description Update price category
// @Tags PriceCategories
// @Accept  json
// @Produce  json
// @Param priceCategory body model.PriceCategories true "Price category data"
// @Success 200 {object} uuid.UUID
// @Failure 500 {object} models.KTSErrorMessage
// @Router /price-categories [put]
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

// @Summary Delete price category
// @Description Delete price category
// @Tags PriceCategories
// @Accept  json
// @Produce  json
// @Param id path string true "Price category ID"
// @Success 200 {string} string "Deleted"
// @Failure 500 {object} models.KTSErrorMessage
// @Router /price-categories/{id} [delete]
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
