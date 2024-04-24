package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pokerror "github.com/linkedinlearning/domina-go/web/web/error"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case pokerror.ErrHttp:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
			}
		}

		c.Next()
	}
}
