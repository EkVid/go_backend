package routes

import (
	"go_backend/controller"
	"github.com/gin-gonic/gin"
)

func QuestionRoutes(router *gin.Engine) {
	question := router.Group("/questions")
	{
		question.GET("/all", controller.GetAllQuestions)
		question.POST("/contribute", controller.ContributeQuestion)
		question.GET("/:id", controller.GetQuestionById)
	}
}