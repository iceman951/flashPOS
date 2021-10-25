package routes

import "github.com/gin-gonic/gin"

func InitUserRoute(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")

	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "all users",
		})
	})

	routerGroup.POST("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "register",
		})
	})

}
