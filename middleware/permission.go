package middleware

import "github.com/gin-gonic/gin"

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic Permission

		// Validate Success
		c.Next()
	}
}
