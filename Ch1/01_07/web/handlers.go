package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/linkedinlearning/domina-go/web/pokemon"
)

func GetPokemon(r *gin.Context) {
	name := r.Param("name")
	p, err := pokemon.Get(name)
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	r.JSON(http.StatusOK, gin.H{"pokemon": p})
}

func GetTypes(r *gin.Context) {
	res, err := pokemon.GetTypes()
	if err != nil {
		r.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	r.JSON(http.StatusOK, gin.H{"types": res})
}
