package services

import (
	"customer_service_gpt/db"
	"customer_service_gpt/models"
)

type MessageService struct{}

func (s *MessageService) CreateMessage(message *models.Message) error {
	return db.DB.Create(message).Error
}

func (s *MessageService) UpdateMessage(message *models.Message) error {
	return db.DB.Save(message).Error
}
