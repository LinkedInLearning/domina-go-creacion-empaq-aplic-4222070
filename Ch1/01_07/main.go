package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/linkedinlearning/domina-go/web/web"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error running the server: %v", err)
	}
}

func run() error {
	// Creates a router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	initializeRoutes(r)

	// By default, it listens on :8080
	return r.Run()
}

func initializeRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/pokemon/:name", web.GetPokemon)
		api.GET("/pokemon/types", web.GetTypes)
	}
}
