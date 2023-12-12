package middlewares

import (
	"context"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if Authorization header is set
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_UNAUTHORIZED)
			return
		}

		// Check if Authorization header is valid
		tokenString, err := utils.ExtractToken(authHeader)
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_UNAUTHORIZED)
			return
		}

		userId, err := utils.ValidateToken(tokenString)
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_UNAUTHORIZED)
			return
		}

		// add userId to request context
		ctx := context.WithValue(c.Request.Context(), "userId", userId)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
