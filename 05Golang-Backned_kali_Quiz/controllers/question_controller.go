package controllers

import (
	"context"
	"net/http"
	"quiz-backend/database"
	"quiz-backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// --- ADMIN CRUD APIs ---

// CreateQuestion - POST /api/questions
func CreateQuestion(c *gin.Context) {
	var input models.Question
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = primitive.NewObjectID()
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	collection := database.GetCollection("questions")
	_, err := collection.InsertOne(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// GetAllQuestions - GET /api/questions
func GetAllQuestions(c *gin.Context) {
	var questions []models.Question
	collection := database.GetCollection("questions")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch questions"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var question models.Question
		cursor.Decode(&question)
		questions = append(questions, question)
	}

	if questions == nil {
		questions = []models.Question{}
	}

	c.JSON(http.StatusOK, questions)
}

// GetQuestionByID - GET /api/questions/:id
func GetQuestionByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var question models.Question
	collection := database.GetCollection("questions")
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&question)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}
	c.JSON(http.StatusOK, question)
}

// UpdateQuestion - PUT /api/questions/:id
func UpdateQuestion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input models.Question
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.GetCollection("questions")
	update := bson.M{
		"$set": bson.M{
			"question":   input.Question,
			"options":    input.Options,
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update question"})
		return
	}

	input.ID = id
	c.JSON(http.StatusOK, input)
}

// DeleteQuestion - DELETE /api/questions/:id
func DeleteQuestion(c *gin.Context) {
	idStr := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := database.GetCollection("questions")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete question"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}
