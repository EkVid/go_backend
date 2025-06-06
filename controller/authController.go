package controller

import (
	"go_backend/models"
	"go_backend/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){
	var credentials models.User

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if services.Authenticate(credentials.Email, credentials.Password) {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	}
}