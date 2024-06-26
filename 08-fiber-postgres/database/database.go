package database

import (
	"log"
	"os"

	"fiber-postgres/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=maugarola dbname=bookapp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		os.Exit(1)
	}

	db.AutoMigrate(&models.Book{}, &models.Category{})

	DB = db
}
