package controller

import (
	"log"
	"todo/model"
	"todo/repository"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task model.Task

	err := c.BindJSON(&task)
	if err != nil {
		log.Fatal(err)
	}
	repository.Insterttask(task)
	c.JSON(200, gin.H{
		"message": "task created",
		"data":    task,
	})
}

func Getalltask(c *gin.Context) {
	task := repository.Getalltask()
	c.JSON(200, gin.H{
		"data": task,
	})
}

func Deletealltask(c *gin.Context) {
	count := repository.Deletealltask()
	c.JSON(200, gin.H{
		"deleted": count,
	})
}

func UpdateOneTask(c *gin.Context) {

	id := c.Param("id")

	var task model.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	count := repository.UpdateOneTask(id, task)

	c.JSON(200, gin.H{
		"updated": count,
	})
}
