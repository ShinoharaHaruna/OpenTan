package middleware

import (
	"OpenTan/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	defaultMaxRequestSize = 200 * 1024 // 200KB default
)

// RequestSizeLimiter returns a middleware that limits the size of incoming requests
func RequestSizeLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := config.Get()
		maxSize := cfg.MaxRequestSize
		
		// Use default if not configured
		if maxSize <= 0 {
			maxSize = defaultMaxRequestSize
		}

		// Check Content-Length header first for efficiency
		if contentLength := c.Request.ContentLength; contentLength > 0 {
			if contentLength > int64(maxSize) {
				c.JSON(413, gin.H{
					"error": "Request too large",
					"message": "The request exceeds the maximum allowed size",
				})
				c.Abort()
				return
			}
		}

		// For requests without Content-Length or chunked encoding, limit the reader
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxSize))
		
		c.Next()
	}
}