package pokemon

import (
	"github.com/gin-gonic/gin"
	"github.com/pokemon/controller/pokemon"
)

func PokemonRoute(api *gin.RouterGroup) {
	apiGroup := api.Group("/pokemon")
	apiGroup.GET("/all", pokemon.GetAllPokemon)

	apiGroup.GET("/:id", pokemon.GetPokemon)

	apiGroup.POST("", pokemon.CreatePokemon)

	apiGroup.PUT("/:id", pokemon.UpdatePokemon)
	apiGroup.DELETE("/:id", pokemon.DeletePokemon)

}
