package main

import (
	"flash-pos.com/flash-pos-api/configs"
	"flash-pos.com/flash-pos-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := setUpRouter()
	router.Run(":4000")
}

func setUpRouter() *gin.Engine {

	configs.Connection()

	router := gin.Default()

	apiV1 := router.Group("/api/v1")

	routes.InitUserRoute(apiV1)

	return router
}
