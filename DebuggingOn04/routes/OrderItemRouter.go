package routes

import (
	controllers "restaurant-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/OrderItem", controllers.GetOrderItems())
	incomingRoutes.GET("/OrderItem/:OrderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/OrderItem-order/:OrderItem_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/OrderItem", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/OrderItem/:OrderItem_id", controllers.UpdateOrderItem())

	incomingRoutes.DELETE("/OrderItem/:OrderItem_id", controllers.DeleteOrderItem())
	incomingRoutes.DELETE("/OrderItem", controllers.DeleteALLOrderItem())

}
