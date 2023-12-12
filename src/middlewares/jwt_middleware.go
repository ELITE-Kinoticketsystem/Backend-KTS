package middlewares

import (
	"context"
	"net/http"

	kts_errors "github.com/ELITE-Kinoticketsystem/Backend-KTS/src/errors"
	"github.com/ELITE-Kinoticketsystem/Backend-KTS/src/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check if cookie is set
		token, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				utils.HandleErrorAndAbort(c, kts_errors.KTS_UNAUTHORIZED)
				return
			}
			utils.HandleErrorAndAbort(c, kts_errors.KTS_BAD_REQUEST)
		}

		userId, err := utils.ValidateToken(token)
		if err != nil {
			utils.HandleErrorAndAbort(c, kts_errors.KTS_UNAUTHORIZED)
			return
		}

		// add userId to request context
		ctx := context.WithValue(c.Request.Context(), utils.UserIdKey, userId)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
