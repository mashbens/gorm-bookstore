package conf

import (
	"books/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	dsn := "root:pass@tcp(127.0.0.1:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	initMgrate()
}

func initMgrate() {
	DB.AutoMigrate(&models.Book{})
}

func DBManager() *gorm.DB {
	return DB
}
