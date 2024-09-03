package main

import (
	"customer_service_gpt/api"
	"customer_service_gpt/config"
	"customer_service_gpt/db"
	"customer_service_gpt/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db.InitDB(cfg)

	// Initialize services
	userService := &services.UserService{}
	messageService := &services.MessageService{}
	gptService := services.NewGPTService(cfg)

	// Set up Gin router
	r := gin.Default()

	// Setup routes
	api.SetupRoutes(r, userService, messageService, gptService)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
