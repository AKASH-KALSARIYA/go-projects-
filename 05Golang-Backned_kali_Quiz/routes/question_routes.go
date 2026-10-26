package routes

import (
	"quiz-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	// Admin CRUD APIs
	r.POST("/questions", controllers.CreateQuestion)       // Create
	r.GET("/questions", controllers.GetAllQuestions)       // Read All
	r.GET("/questions/:id", controllers.GetQuestionByID)   // Read One
	r.PUT("/questions/:id", controllers.UpdateQuestion)    // Update
	r.DELETE("/questions/:id", controllers.DeleteQuestion) // Delete
}
