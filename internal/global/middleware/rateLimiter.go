// internal/global/middleware/rateLimiter.go

package middleware

import (
	"OpenTan/config"
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

type RateLimiterConfig struct {
	Rate    rate.Limit
	Burst   int
	KeyFunc func(*gin.Context) string
	MaxWait time.Duration
}

func NewRateLimiter(cfg RateLimiterConfig) gin.HandlerFunc {
	var mu sync.Mutex
	clients := make(map[string]*rate.Limiter)
	semaphores := make(map[string]*semaphore.Weighted)

	if cfg.KeyFunc == nil {
		cfg.KeyFunc = func(c *gin.Context) string {
			return c.ClientIP()
		}
	}

	return func(c *gin.Context) {
		if c.Request.URL.Path != config.Get().Prefix+"/chat/completions" {
			// Only apply rate limiting to the completions endpoint
			c.Next()
			return
		}
		key := cfg.KeyFunc(c)

		mu.Lock()
		limiter, ok := clients[key]
		if !ok {
			limiter = rate.NewLimiter(cfg.Rate, cfg.Burst)
			clients[key] = limiter
			semaphores[key] = semaphore.NewWeighted(int64(cfg.Burst))
		}
		mu.Unlock()

		ctx, cancel := context.WithTimeout(c.Request.Context(), cfg.MaxWait)
		defer cancel()

		if err := semaphores[key].Acquire(ctx, 1); err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests (timeout)"})
			return
		}
		defer semaphores[key].Release(1)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			return
		}

		c.Next()
	}
}
