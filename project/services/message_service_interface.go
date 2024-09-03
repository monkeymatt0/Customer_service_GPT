package services

import (
	"customer_service_gpt/models"
)

type MessageServiceInterface interface {
	CreateMessage(message *models.Message) error
	UpdateMessage(message *models.Message) error
}
