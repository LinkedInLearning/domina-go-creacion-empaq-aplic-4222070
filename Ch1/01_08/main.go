package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/linkedinlearning/domina-go/web/web"
	"github.com/linkedinlearning/domina-go/web/web/middleware"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error running the server: %v", err)
	}
}

func run() error {
	requiredToken := os.Getenv("API_TOKEN")

	// We want to make sure the server token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	// Creates a router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	initializeMiddleware(r)

	initializeRoutes(r)

	// By default, it listens on :8080
	return r.Run()
}

func initializeMiddleware(r *gin.Engine) {
	r.Use(
		middleware.Counter(),   // contador de peticiones
		middleware.Logging(),   // log de peticiones
		middleware.AuthToken(), // autenticación
		middleware.Error(),     // manejo de errores
	)
}

func initializeRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/pokemon/:name", web.GetPokemon)
		api.GET("/pokemon/types", web.GetTypes)
	}
}
