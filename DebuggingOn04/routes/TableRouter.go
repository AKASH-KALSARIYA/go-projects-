package routes

import (
	controllers "restaurant-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/Table", controllers.GetTables())
	incomingRoutes.GET("/Table/:Order_id", controllers.GetTable())
	incomingRoutes.POST("/Table", controllers.CreateTable())
	incomingRoutes.PATCH("/Table/:Order_id", controllers.UpdateTable())

	incomingRoutes.DELETE("/Table/:Order_id", controllers.DeleteTable())
	incomingRoutes.DELETE("/Table", controllers.DeleteALLTable())

}
