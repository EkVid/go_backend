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
		return false
	}

	return Password == user.Password
}

func CreateUser(Email string, Username string) error {

	user := models.User{
		Email: Email,
		Username: Username,
		Password: "test",
	}

	res := db.DB.Create(&user)
	return res.Error
}