package services

import (
	"customer_service_gpt/models"
)

type UserServiceInterface interface {
	// User related
	CreateUser(user *models.User) (uint, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint) (*models.User, error)
	DeleteUser(id uint) error
	GetUserByEmail(email string) (*models.User, error)

	// Session related
	CreateSession(session *models.UserSession) error
	DeleteSession(id uint) error
	GetSession(id uint) (*models.UserSession, error)
}
