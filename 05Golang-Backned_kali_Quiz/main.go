package main

import (
	"quiz-backend/database"
	"quiz-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize Database (Only Questions table)
	database.ConnectDB()

	// 2. Initialize Gin
	r := gin.Default()

	// 3. Register Routes
	api := r.Group("/api")
	routes.RegisterRoutes(api)

	// 4. Start Server
	r.Run(":6060")
}
