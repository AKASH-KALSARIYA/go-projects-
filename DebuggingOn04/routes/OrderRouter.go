package routes

import (
	controllers "restaurant-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/Order", controllers.GetOrders())
	incomingRoutes.GET("/Order/:Order_id", controllers.GetOrder())
	incomingRoutes.POST("/Order", controllers.CreateOrder())
	incomingRoutes.PATCH("/Order/:Order_id", controllers.UpdateOrder())

	incomingRoutes.DELETE("/Order/:Order_id", controllers.DeleteOrder())
	incomingRoutes.DELETE("/Order", controllers.DeleteALLOrder())

}
