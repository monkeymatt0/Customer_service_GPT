package services

import (
	"customer_service_gpt/db"
	"customer_service_gpt/models"
)

type UserService struct{}

func (s *UserService) CreateUser(user *models.User) (uint, error) {
	tx := db.DB.Create(user)
	return user.ID, tx.Error
}
func (s *UserService) UpdateUser(user *models.User) error {
	return db.DB.Model(&user).Updates(user).Error
}
func (s *UserService) DeleteUser(id uint) error {
	var user models.User
	return db.DB.Delete(&user, id).Error
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (s *UserService) CreateSession(session *models.UserSession) error {
	return db.DB.Create(session).Error
}

func (s *UserService) GetSession(id uint) (*models.UserSession, error) {
	var session *models.UserSession
	if err := db.DB.Where("user_id = ?", id).First(&session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (s *UserService) DeleteSession(id uint) error {
	var userSession models.UserSession
	return db.DB.Delete(&userSession, id).Error
}
