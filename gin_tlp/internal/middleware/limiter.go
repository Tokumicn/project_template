package middleware

import (
	"gin_tlp/pkg/app"
	"gin_tlp/pkg/errcode"
	"gin_tlp/pkg/limiter"
	"github.com/gin-gonic/gin"
)

// RateLimiter 限流器
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			// 取不出许可令牌时，本次请求提交中断
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
