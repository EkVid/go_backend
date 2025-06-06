package main

import (
	"go_backend/db"
	"go_backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.Connect();

	r := gin.Default()

	routes.AuthRoutes(r)

	r.Run(":8080")
}
