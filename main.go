package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sneha-afk/astroauth/routes"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
