package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID   uint   `gorm:"not null"`
	User     User   `gorm:"foreignKey:UserID"`
	Message  string `gorm:"not null"`
	Response string
}
