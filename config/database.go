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
	// Create user_role_enum if not exists
	db.Exec(`
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role_enum') THEN
			CREATE TYPE user_role_enum AS ENUM ('admin', 'teller');
		END IF;
	END$$;
	`)

	// Create sales_order_status_enum if not exists
	db.Exec(`
	DO $$
	BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'sales_order_status_enum') THEN
			CREATE TYPE sales_order_status_enum AS ENUM ('pending', 'completed');
		END IF;
	END$$;
	`)

	// Auto-migrate models
	if err := db.AutoMigrate(&models.Category{}, &models.Product{}, &models.Supplier{}, &models.User{}, &models.Customer{}, &models.SalesOrder{}); err != nil {
		log.Fatal(" Failed to auto-migrate models:", err)
	}
}
