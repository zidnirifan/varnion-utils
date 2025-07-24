package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zidnirifan/varnion-utils/authentication"
	"github.com/zidnirifan/varnion-utils/tools"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.Response{
				Status:  "Unauthorized",
				Message: "Authorization header is required",
			})
			return
		}

		// Check and extract Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.Response{
				Status:  "Unauthorized",
				Message: "Invalid Authorization header format",
			})
			return
		}

		// Check if token exists
		tokenString := parts[1]
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.Response{
				Status:  "Unauthorized",
				Message: "Token cannot be empty",
			})
			return
		}

		// Logic Authentication
		claims, err := tools.ValidateAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, tools.Response{
				Status:  "Unauthorized",
				Message: "Invalid or expired token",
			})
			return
		}

		// Store claims in context
		c.Set(authentication.UserIDKey, claims.UserID.String())

		// Validate Success
		c.Next()
	}
}
