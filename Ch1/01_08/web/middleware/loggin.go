package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Println("Request received")

		// retrieve the request count from the context
		count := c.MustGet(RequestCountKey).(int)

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		log.Printf("Request #%d - HTTP Status: %d, in %v\n", count, status, latency)
	}
}
