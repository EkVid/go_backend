package db

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"go_backend/models"
)

var DB *gorm.DB

func Connect(user, password, dbname string) {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", user, password, dbname)
   	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        panic("Failed to connect to database!")
    }

    DB = db
    fmt.Println("Database connected")
	DB.AutoMigrate(&models.User{})
}