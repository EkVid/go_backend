package db

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
	"go_backend/models"
)

var DB *gorm.DB

func Connect(){
	dsn := "host=localhost user=test_user password=test dbname=go_backend_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        panic("Failed to connect to database!")
    }

    DB = db
    fmt.Println("Database connected")
	DB.AutoMigrate(&models.User{})
}