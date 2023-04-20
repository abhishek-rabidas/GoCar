package config

import (
	"fmt"
	"gocar/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	dsn := "root:admin@tcp(127.0.0.1:3306)/gocar?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to db")
	}

	db.AutoMigrate(&models.Car{})

	return db
}
