package main

import (
	"capstone/backend/driver"
	_middleware "capstone/backend/middlewares"
	"capstone/backend/routes"
)

// func init() {
// 	viper.SetConfigFile("app/config/config.json")
// 	err := viper.ReadInConfig()

// 	if err != nil {
// 		panic(err)
// 	}

// 	if viper.GetBool(`debug`) {
// 		log.Println("Service RUN on DEBUG mode")
// 	}
// }
// func dbMigrate(db *gorm.DB) {

// 	db.AutoMigrate(&news.News{}, &onlineClass.OnlineClassCore{}, &offlineClass.OfflineClassCore{}, &user.User{}, &bookingOffline.OfflineClassUser{}, &bookingOnline.OnlineClassUser{})
// }

func main() {

	// configDb := mysql.ConfigDB{
	// 	DB_Username: viper.GetString(`database.user`),
	// 	DB_Password: viper.GetString(`database.pass`),
	// 	DB_Host:     viper.GetString(`database.host`),
	// 	DB_Port:     viper.GetString(`database.port`),
	// 	DB_Database: viper.GetString(`database.name`),
	// }

	// db := configDb.InitDB()
	// dbMigrate(db)
	driver.InitDB()
	e := routes.New()
	// Log Middleware
	_middleware.LogMiddlewareInit(e)

	e.Start(":8000")

}
