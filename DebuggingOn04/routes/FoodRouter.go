package routes

import (
	controllers "restaurant-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:Food_id", controllers.GetFood())
	incomingRoutes.POST("/foods", controllers.CreateFoods())
	incomingRoutes.PATCH("/foods/:Food_id", controllers.UpdateFood())

	incomingRoutes.DELETE("/foods/:Food_id", controllers.DeleteFood())
	incomingRoutes.DELETE("/foods", controllers.DeleteALLFoods())

}
