// api-go\src\middleware\rate_limit.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// rpsは1秒あたりに許可するリクエスト数、burstは短期間で許可するリクエストの最大数を指定
func RateLimitMiddleware(rps int, burst int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(rps), burst)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too Many Requests",
			})
			return
		}
		c.Next()
	}
}
