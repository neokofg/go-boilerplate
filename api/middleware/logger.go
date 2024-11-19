package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logger(sugar *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		sugar.Infow("HTTP Request",
			"path", path,
			"method", c.Request.Method,
			"status", c.Writer.Status(),
			"latency", time.Since(start),
		)
	}
}
