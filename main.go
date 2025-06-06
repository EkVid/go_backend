package main

import (
	"go_backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.AuthRoutes(r)

	r.Run(":8080")
}
