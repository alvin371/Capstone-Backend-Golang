package driver

import (
	news "capstone/backend/features/News/data"
	user "capstone/backend/features/User"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "u280225155_gym:Rahasia12345@tcp(194.163.35.1)/u280225155_gym?charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
	DB.AutoMigrate(&news.News{}, &user.User{})
}
