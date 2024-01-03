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

// User has to be logged in
// @Summary Create order
// @Description Create order
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param eventId path string true "Event ID"
// @Param order body models.CreateOrderDTO true "Order data"
// @Success 200 {object} models.IdResponse
// @Failure 500 {object} models.KTSErrorMessage
// @Failure 400 {object} models.KTSErrorMessage
// @Router /events/{eventId}/reserve [post]
// @Router /events/{eventId}/book [post]
func CreateOrderHandler(orderController controllers.OrderControllerI, isReservation bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		eventId, err := uuid.Parse(c.Param("eventId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)

		createOrderDTO := models.CreateOrderDTO{}

		if err := c.ShouldBindJSON(&createOrderDTO); err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		order, kts_err := orderController.CreateOrder(createOrderDTO, &eventId, userId, isReservation)

		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(http.StatusOK, models.IdResponse{
			Id: order,
		})

	}
}

// @Summary Get order by id
// @Description Get order by id
// @Tags Orders
// @Accept  json
// @Produce  json
// @Param orderId path string true "Order ID"
// @Success 200 {object} models.GetOrderDTO
// @Failure 400 {object} models.KTSErrorMessage
// @Failure 404 {object} models.KTSErrorMessage
// @Failure 500 {object} models.KTSErrorMessage
// @Router /orders/{orderId} [get]
func GetOrderByIdHandler(orderController controllers.OrderControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId, err := uuid.Parse(c.Param("orderId"))

		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
			return
		}

		userId := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)

		order, kts_err := orderController.GetOrderById(&orderId, userId)

		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(200, order)
	}
}

// @Summary Get orders
// @Description Get orders for user
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {array} models.GetOrderDTO
// @Failure 500 {object} models.KTSErrorMessage
// @Router /orders [get]
func GetOrdersHandler(orderController controllers.OrderControllerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)

		orders, kts_err := orderController.GetOrders(userId)

		if kts_err != nil {
			utils.HandleErrorAndAbort(c, kts_err)
			return
		}

		c.JSON(200, orders)
	}
}
