package middlewares

import (
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/models"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Request.Context().Value(models.ContextKeyUserID).(*uuid.UUID)
		if !ok {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_UNAUTHORIZED)
		}

		adminId, err := uuid.Parse("dddddddd-dddd-dddd-dddd-dddddddddddd")
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_INTERNAL_ERROR)
			return
		}

		if *userId != adminId {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_FORBIDDEN)
			return
		}
		c.Next()
	}
}
