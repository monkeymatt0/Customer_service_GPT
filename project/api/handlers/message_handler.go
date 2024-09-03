package handlers

import (
	"customer_service_gpt/models"
	"customer_service_gpt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	messageService *services.MessageService
	gptService     *services.GPTService
}

func NewMessageHandler(messageService *services.MessageService, gptService *services.GPTService) *MessageHandler {
	return &MessageHandler{
		messageService: messageService,
		gptService:     gptService,
	}
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from the authenticated context
	userID, _ := c.Get("user_id")
	message.UserID = userID.(uint)

	// Save the message to the database
	if err := h.messageService.CreateMessage(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create message"})
		return
	}

	// Get GPT response
	gptResponse, err := h.gptService.GetResponse(message.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get GPT response"})
		return
	}

	// Update the message with the GPT response
	message.Response = gptResponse
	if err := h.messageService.UpdateMessage(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update message with response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message.Message, "response": message.Response})
}
