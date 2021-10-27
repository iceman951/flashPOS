package routes

import (
	usercontroller "flash-pos.com/flash-pos-api/controllers/user"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users")

	routerGroup.GET("/", usercontroller.GetAll)

	routerGroup.POST("/register", usercontroller.Register)

	routerGroup.POST("/login", usercontroller.Login)

}
