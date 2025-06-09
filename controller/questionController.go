package controller

import (
	"go_backend/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"go_backend/models"
)

func ContributeQuestion(c *gin.Context){
	var input models.Question

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateQuestion(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
	c.JSON(http.StatusOK, gin.H{"message": "Question created successfully"})
	}
}

func GetAllQuestions(c *gin.Context){
	questions, err := services.FindAllQuestions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, questions)
}

func GetQuestionById(c *gin.Context){
	id := c.Param("id")
	question, err := services.FindQuestionById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, question)
}