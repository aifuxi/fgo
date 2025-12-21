package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMySQL() {
	dsn := "root:123456@tcp(127.0.0.1:8306)/fgo?charset=utf8mb4&parseTime=True&loc=Local"

	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
		return
	}

	log.Println("Connected to MySQL")
}

func GetDB() *gorm.DB {
	return db
}
