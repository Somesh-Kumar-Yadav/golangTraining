package main

import (
	"github.com/gin-gonic/gin"
	route "github.com/pokemon/routes"
)

func main() {
	r := gin.Default()

	routes := r.Group("/")
	route.Route(routes)

	r.Run(":8080")
}
