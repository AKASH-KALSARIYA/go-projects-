package routes

import (
	controllers "restaurant-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/menu", controllers.Getmenus())
	incomingRoutes.GET("/menu/:menu_id", controllers.Getmenu())
	incomingRoutes.POST("/menu", controllers.Createmenu())
	incomingRoutes.PATCH("/menu/:menu_id", controllers.Updatemenu())

	incomingRoutes.DELETE("/menu/:menu_id", controllers.Deletemenu())
	incomingRoutes.DELETE("/menu", controllers.DeleteALLmenu())

}
