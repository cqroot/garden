package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

type LogFields struct {
	StatusCode int
	Latency    time.Duration
	ClientIp   string
	Method     string
	Path       string
}

func CustomLogger(logFunc func(LogFields)) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		timeStamp := time.Now()
		latency := timeStamp.Sub(start)

		if raw != "" {
			path = path + "?" + raw
		}

		logFunc(LogFields{
			StatusCode: c.Writer.Status(),
			Latency:    latency,
			ClientIp:   c.ClientIP(),
			Method:     c.Request.Method,
			Path:       path,
		})
	}
}
