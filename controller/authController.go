package controller

import (
	"go_backend/services"
	"net/http"
	"github.com/gin-gonic/gin"
)


type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context){
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if services.Authenticate(input.Email, input.Password) {
		c.SetCookie("isAuthenticated", "true", 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"success": true})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	}
}

func Register(c *gin.Context){
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	_, err := services.GetUserByEmail(input.Email)
	if err == nil {
		// User already exists
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists with this email"})
		return
	}

	err = services.CreateUser(input.Email, input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func Logout(c *gin.Context){
	c.SetCookie("isAuthenticated", "false", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"success": true})
}