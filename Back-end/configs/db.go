package configs

import (
	"fmt"

	"flash-pos.com/flash-pos-api/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	// dsn := os.Getenv("DATABASE_DSN")
	dsn := "server=flash-pos-server.database.windows.net;user id=azure-admin;password=fanhin-ice319;port=1433;database=flas-pos-db;"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("ไม่สามารถติดต่อกับ Database Server ได้")
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("ติดต่อฐานข้อมูลสำเร็จ")

	db.AutoMigrate(&models.User{})
	DB = db

}
