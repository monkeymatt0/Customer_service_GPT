package models

import (
	"gorm.io/gorm"
)

type UserSession struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`
	Token  string `gorm:"not null"`
}
