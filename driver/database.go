package driver

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	var err error
	// dsn := "u280225155_gym:Rahasia12345@tcp(194.163.35.1)/u280225155_gym?charset=utf8mb4&parseTime=true&loc=Local"
	// dsnLocal := "root:@tcp(host.docker.internal:3306)/gym?charset=utf8mb4&parseTime=true&loc=Local"

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username, config.DB_Password, config.DB_Host, config.DB_Port, config.DB_Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
