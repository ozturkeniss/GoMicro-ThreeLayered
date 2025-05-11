package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/ozturkeniss/gomicro-app/order-service/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database")

	// AutoMigrate the Order model
	err = DB.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("Failed to auto migrate Order model: %v", err)
	}
	log.Println("Order table auto migrated successfully")
}

// TestDatabaseConnection tries to connect to the database and logs the result
func TestDatabaseConnection() {
	if DB == nil {
		log.Println("[ERROR] Database connection is not initialized.")
		return
	}
	db, err := DB.DB()
	if err != nil {
		log.Printf("[ERROR] Failed to get generic database object: %v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Printf("[ERROR] Database ping failed: %v", err)
		return
	}
	log.Println("[INFO] Database connection is healthy.")
} 