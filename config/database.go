package config

import (
	"fmt"
	"github.com/john-henry-alfeche/inventory-app-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Build DSN from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEZONE"),
	)

	// Open DB connection with SQL logging enabled
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // log SQL queries
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Store to global variable
	DB = db
	log.Println("Database connection established.")

	// Enable UUID extension (if not already done in pg)
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Auto-migrate models
	if err := db.AutoMigrate(&models.Category{}, &models.Product{}); err != nil {
		log.Fatal(" Failed to auto-migrate models:", err)
	}
}
