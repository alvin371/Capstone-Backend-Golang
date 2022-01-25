package driver

import (
	news "capstone/backend/features/News/data"
	user "capstone/backend/features/User"
	bookingOffline "capstone/backend/features/bookingOffline"
	offlineClass "capstone/backend/features/offlineClass"
	onlineClass "capstone/backend/features/onlineClass"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// dsn := "u280225155_gym:Rahasia12345@tcp(194.163.35.1)/u280225155_gym?charset=utf8mb4&parseTime=true&loc=Local"
	dsnLocal := "root:@tcp(localhost)/gym?charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsnLocal), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = db
	DB.AutoMigrate(&news.News{}, &onlineClass.OnlineClassCore{}, &offlineClass.OfflineClassCore{}, &user.User{}, &bookingOffline.OfflineClassUser{})
}
