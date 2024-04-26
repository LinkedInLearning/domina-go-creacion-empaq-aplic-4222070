package middleware

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type Count struct {
	count int
	mu    sync.Mutex
}

func (c *Count) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

var count Count

const RequestCountKey = "requestCount"

func Counter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		count.Increment()

		ctx.Set(RequestCountKey, count.count)

		ctx.Next()
	}
}
