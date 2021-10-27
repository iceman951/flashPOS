package usercontroller

import (
	"net/http"

	"flash-pos.com/flash-pos-api/configs"
	"flash-pos.com/flash-pos-api/models"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

func GetAll(c *gin.Context) {
	var users []models.User

	configs.DB.Find(&users)
	c.JSON(200, gin.H{
		"success": true,
		"messege": "ดึงข้อมูล all users สำเร็จ",
		"data":    users,
	})
}

func Register(c *gin.Context) {

	var input InputRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	//เช็ค email ซ้ำ
	userExist := configs.DB.Where("email = ?", input.Email).First(&user)
	if userExist.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "มีผู้ใช้งานอีเมล์นี้ในระบบแล้ว"})
		return
	}

	result := configs.DB.Debug().Create(&user)

	//db error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "ลงทะเบียนสำเร็จ",
	})
}

func Login(c *gin.Context) {
	var input InputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	//เช็คว่ามีผู้ใช้นี้ในระบบเราหรือไม่
	userAccount := configs.DB.Where("email = ?", input.Email).First(&user)
	if userAccount.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบผู้ใช้งานนี้ในระบบ"})
		return
	}

	//เปรียบเทียบ password
	ok, _ := argon2.VerifyEncoded([]byte(input.Password), []byte(user.Password))
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "รหัสผ่านไม่ถูกต้อง"})
		return
	}

	c.JSON(200, gin.H{
		"success": ok,
		"messege": user.Email + " เข้าสู่ระบบสำเร็จ",
	})

}
