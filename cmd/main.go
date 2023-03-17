package main

import (
	"github.com/gabrielgl9/removies/internal/infra/web/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)
	
	router.Run()
}