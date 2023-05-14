package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pokemon/routes/pokemon"
)

func Route(api *gin.RouterGroup) {
	route := api.Group("/v1")
	pokemon.PokemonRoute(route)
}
