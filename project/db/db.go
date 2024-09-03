package db

import (
	"customer_service_gpt/config"
	"customer_service_gpt/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migrate the models
	err = DB.AutoMigrate(&models.User{}, &models.Message{}, &models.UserSession{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	log.Println("Database connected and migrated successfully")
}
