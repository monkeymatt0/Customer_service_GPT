package services

import (
	"customer_service_gpt/models"
)

type UserServiceInterface interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	CreateSession(session *models.UserSession) error
}
