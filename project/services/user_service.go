package services

import (
	"customer_service_gpt/db"
	"customer_service_gpt/models"
)

type UserService struct{}

func (s *UserService) CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (s *UserService) CreateSession(session *models.UserSession) error {
	return db.DB.Create(session).Error
}
