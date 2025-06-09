package services

import (
	"go_backend/models"
	"go_backend/db"
)

func CreateQuestion(question models.Question) error {
	return db.DB.Create(&question).Error
}

func FindAllQuestions() ([]models.Question, error) {
	var questions []models.Question
	res := db.DB.Find(&questions)
	return questions, res.Error
}

func FindQuestionById(id string) (models.Question, error) {
	var question models.Question
	res := db.DB.Where("id = ?", id).First(&question)
	return question, res.Error
}