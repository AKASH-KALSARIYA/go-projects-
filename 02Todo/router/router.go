package router

import (
	"todo/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.GET("/", controller.Getalltask)
	router.POST("/tasks", controller.CreateTask)
	router.DELETE("/deletetasks", controller.Deletealltask)
	router.PUT("/update/:id", controller.UpdateOneTask)

	return router
}
