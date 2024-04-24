package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	pokerror "github.com/linkedinlearning/domina-go/web/web/error"
)

func AuthToken() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")

	// We want to make sure the server token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(ctx *gin.Context) {
		// read the token from the Authorization header
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			err := pokerror.NewHttpError("API token required", "You must pass the api_token", http.StatusUnauthorized)
			ctx.Error(err)
			return
		}

		// compare the token from the header with the server token
		if token != requiredToken {
			err := pokerror.NewHttpError("Invalid API token", "You must pass a valid api_token", http.StatusUnauthorized)
			ctx.Error(err)
			return
		}

		ctx.Next()
	}
}
