package routes

import (
	controllers "restaurant-management-system/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/invoice", controllers.Getinvoices())
	incomingRoutes.GET("/invoice/:invoice_id", controllers.Getinvoice())
	incomingRoutes.POST("/invoice", controllers.Createinvoice())
	incomingRoutes.PATCH("/invoice/:invoice_id", controllers.Updateinvoice())

	incomingRoutes.DELETE("/invoice/:invoice_id", controllers.Deleteinvoice())
	incomingRoutes.DELETE("/invoice", controllers.DeleteALLinvoice())

}
