package middleware

import (
	"github.com/aomi-go/common/web/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

func TraceIdMiddleware(c *gin.Context) {
	// Get trace id from request header
	reqId := c.GetHeader(http.RequestId)
	if reqId == "" {
		reqId = uuid.New().String()
		reqId = strings.ReplaceAll(reqId, "-", "")
	}
	c.Set("traceId", reqId)
	c.Header(http.RequestId, reqId)
	// Process request
	c.Next()

}
