package api

import (
	"customer_service_gpt/api/handlers"
	"customer_service_gpt/api/middlewares"
	"customer_service_gpt/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userService *services.UserService, messageService *services.MessageService, gptService *services.GPTService) {
	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	messageHandler := handlers.NewMessageHandler(messageService, gptService)

	// Public routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// Protected routes
	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		api.POST("/messages", messageHandler.CreateMessage)
		api.DELETE("/logout", userHandler.Logout)
		api.PATCH("/user/:id", userHandler.UpdateUser)
		// Add other protected routes here
	}
}
