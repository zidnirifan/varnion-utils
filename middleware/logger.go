package middleware

import "github.com/gin-gonic/gin"

func Logger(SkipPath ...string) gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: SkipPath,
	})
}
