package config

import (
	"fmt"
	"log"
	"os"

	"go-auth-app/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.AutoMigrate(&models.User{}, &models.ProductWarranty{}, &models.RepairWarranty{}, &models.RepairDetail{}, &models.RepairData{})
	var count int64

	db.Model(&models.User{}).Where("role = ?", "admin").Count(&count)
	if count == 0 {
		adminUsername := os.Getenv("ADMIN_USERNAME")
		adminPassword := os.Getenv("ADMIN_PASSWORD")
		adminPhone := os.Getenv("ADMIN_PHONE")
		adminGender := os.Getenv("ADMIN_GENDER")
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
		admin := models.User{
			Username: adminUsername,
			Phone:    adminPhone,
			Password: string(hashedPassword),
			Gender:   adminGender,
			Role:     "admin",
		}
		db.Create(&admin)
	}

	log.Println("Connected to database successfully!")
	DB = db
}
