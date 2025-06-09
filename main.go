package main

import (
	"go_backend/db"
	"go_backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
	"github.com/joho/godotenv"
	"os"
	"log"
)

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Access environment variables
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	// connect to db
	db.Connect(dbUser, dbPassword, dbName);

	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"}, 
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.AuthRoutes(r)
	routes.QuestionRoutes(r)

	r.Run(":8080")
}
