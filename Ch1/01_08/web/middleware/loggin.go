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
		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()
		log.Printf("Request completed - HTTP Status: %d, in %v\n", status, latency)
	}
}
