package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	pokerror "github.com/linkedinlearning/domina-go/service/web/error"
)

func AuthToken() gin.HandlerFunc {
	requiredToken := os.Getenv("API_TOKEN")

	return func(ctx *gin.Context) {
		// el token debe ser pasado en el header de la petición
		token := ctx.Request.Header.Get("Authorization")

		// si no hay token, devolvemos un error
		if token == "" {
			err := pokerror.NewHttpError("API token required", "You must pass the api_token", http.StatusUnauthorized)
			ctx.Error(err)
			return
		}

		// si el token no coincide con el esperado, devolvemos un error
		if token != requiredToken {
			err := pokerror.NewHttpError("Invalid API token", "You must pass a valid api_token", http.StatusUnauthorized)
			ctx.Error(err)
			return
		}

		ctx.Next()
	}
}
