package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/varnion-rnd/utils/tools"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		parts := strings.Fields(tokenString)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, tools.Response{
				Message: "Token is missing",
			})
			c.Abort()
			return
		}

		if len(parts) < 2 {
			c.JSON(http.StatusUnauthorized, tools.Response{
				Message: "Token is missing",
			})
			c.Abort()
			return
		}

		// Logic Authentication

		// Validate Success
		c.Next()
	}
}
