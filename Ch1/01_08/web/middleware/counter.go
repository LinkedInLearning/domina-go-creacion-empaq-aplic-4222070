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

// count es una variable global que se incrementa con cada petición
var count Count

const RequestCountKey = "requestCount"

func Counter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		count.Increment()

		// Añadimos el valor del contador a la petición
		ctx.Set(RequestCountKey, count.count)

		ctx.Next()
	}
}
