package usercontroller

import (
	"net/http"

	"flash-pos.com/flash-pos-api/configs"
	"flash-pos.com/flash-pos-api/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input InputRegister
	user := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	result := configs.DB.Debug().Create(&user)

	//db error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{
		"message": "register success",
	})

}
