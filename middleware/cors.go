package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

const (
	ALLOW_ORIGIN string = "ALLOW_ORIGIN"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv(ALLOW_ORIGIN) != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv(ALLOW_ORIGIN))
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Referer, User-Agent, Content-Type, Content-Length, Accept-Language, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
