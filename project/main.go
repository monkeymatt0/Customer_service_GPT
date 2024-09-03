package main

import (
	"customer_service_gpt/api/handlers"
	"customer_service_gpt/api/middlewares"
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

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	messageHandler := handlers.NewMessageHandler(messageService, gptService)

	// Set up Gin router
	r := gin.Default()

	// Public routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/messages", messageHandler.CreateMessage)
	}

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
