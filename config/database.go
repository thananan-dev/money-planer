package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load configuration
	LoadConfig()

	// Connect to database using the configuration
	db, err := gorm.Open(postgres.Open(AppConfig.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Set some database configuration based on environment
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}

	if AppConfig.AppEnv.IsDevelopment() {
		// Development settings
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetMaxOpenConns(10)
	} else {
		// Production settings
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	}

	DB = db
}
