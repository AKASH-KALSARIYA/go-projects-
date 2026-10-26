package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"restaurant-management-system/database"
	"restaurant-management-system/routes"

	"go.mongodb.org/mongo-driver/mongo"
)

var FoodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "1111"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
