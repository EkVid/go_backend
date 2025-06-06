package services

import "go_backend/models"

var users = models.User{
	Email: "test@mail.com",
	Password: "test",
}

func Authenticate(Email string, Password string) bool{
	return Email == users.Email && Password == users.Password
}