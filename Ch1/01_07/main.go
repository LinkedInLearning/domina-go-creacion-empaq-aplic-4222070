package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linkedinlearning/domina-go/web/web"
)

func main() {
	// Creates a router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	initializeRoutes(r)

	// By default, it listens on :8080
	r.Run()
}

func initializeRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/pokemon/:name", web.GetPokemon)
		api.GET("/pokemon/types", web.GetTypes)
	}
}
