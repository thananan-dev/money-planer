package main

import (
	"money-planer/config"
	"money-planer/models"
	"money-planer/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	config.ConnectDatabase()

	// Auto migrate the schema
	config.DB.AutoMigrate(&models.Transaction{})

	// Create Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Run the server
	r.Run(":8080")
}
