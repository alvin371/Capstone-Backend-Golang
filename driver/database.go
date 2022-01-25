package driver

import (
	news "capstone/backend/features/News/data"
	user "capstone/backend/features/User"
	onlineClass "capstone/backend/features/onlineClass"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsnLocal := "root:@tcp(localhost)/gym?charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsnLocal), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
	DB.AutoMigrate(&news.News{}, &onlineClass.OnlineClassCore{}, &user.User{})
}
