package routes

import (
	usercontroller "flash-pos.com/flash-pos-api/controllers/user"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")

	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "all users",
		})
	})

	routerGroup.POST("/register", usercontroller.Register)

}
