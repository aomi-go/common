package middleware

import (
	"github.com/aomi-go/common/logger"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerMiddleware(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()
		// Stop timer
		stop := time.Now()

		if raw != "" {
			path = path + "?" + raw
		}
		log.Int("statusCode", c.Writer.Status()).
			Duration("latency", stop.Sub(start)).
			String("clientIP", c.ClientIP()).
			String("method", c.Request.Method).
			String("path", path).
			Debug(c.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}
