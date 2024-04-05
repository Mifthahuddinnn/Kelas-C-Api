package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mahasiswa/models"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:1234@tcp(localhost:8080)/C21?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	if err := db.AutoMigrate(&models.Student{}, &models.Account{}); err != nil {
		panic("failed to migrate database schema: " + err.Error())
	}
	return db
}

func Init() {
	DB = InitDB()
}
