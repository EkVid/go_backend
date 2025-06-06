package services

import (
	"go_backend/models"
	"go_backend/db"
)

func GetUserByEmail(Email string) (models.User, error){
	var user models.User
	res := db.DB.Where("email = ?", Email).First(&user)
	return user, res.Error
}

func Authenticate(Email string, Password string) bool {
	user, err := GetUserByEmail(Email)
	if err != nil {
		// User not found or DB error
		return false
	}
	return Password == user.Password
}

func CreateUser(Email string, Username string) error {

	user := models.User{
		Email: Email,
		Username: Username,
	}

	res := db.DB.Create(&user)
	return res.Error
}