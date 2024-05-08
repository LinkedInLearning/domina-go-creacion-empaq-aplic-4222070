package web

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/linkedinlearning/domina-go/service/pokemon"
	pokerror "github.com/linkedinlearning/domina-go/service/web/error"
)

func GetPokemon(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		err := pokerror.NewHttpError("Named parameter not found", "name parameter is required", http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	p, err := pokemon.Get(name)
	if err != nil {
		err = pokerror.NewHttpError("Error getting pokemon", err.Error(), http.StatusInternalServerError)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"pokemon": p})
}

func GetTypes(ctx *gin.Context) {
	res, err := pokemon.GetTypes()
	if err != nil {
		err = pokerror.NewHttpError("Error getting pokemon types", err.Error(), http.StatusInternalServerError)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"types": res})
}
