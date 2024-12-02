package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Environment represents the application environment
type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

// Config holds all configuration
type Config struct {
	AppEnv      Environment
	DatabaseURL string
	DBHost      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
}

var AppConfig Config

// LoadConfig loads the configuration based on the environment
func LoadConfig() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development" // Default to development if not specified
	}

	// Load the appropriate .env file
	envFile := ".env"
	if env != "production" {
		envFile = fmt.Sprintf(".env.%s", env)
	}

	// Try to load environment-specific file
	err := godotenv.Load(envFile)
	if err != nil {
		// If environment-specific file doesn't exist, try loading default .env
		err = godotenv.Load()
		if err != nil {
			log.Printf("Warning: Error loading %s file", envFile)
		}
	}

	AppConfig = Config{
		AppEnv:      Environment(env),
		DBHost:      getEnvWithDefault("DB_HOST", "localhost"),
		DBUser:      getEnvWithDefault("DB_USER", "postgres"),
		DBPassword:  getEnvWithDefault("DB_PASSWORD", ""),
		DBName:      getEnvWithDefault("DB_NAME", "money_planer"),
		DBPort:      getEnvWithDefault("DB_PORT", "5432"),
	}

	// Construct Database URL
	AppConfig.DatabaseURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		AppConfig.DBHost,
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBName,
		AppConfig.DBPort,
	)
}

// getEnvWithDefault gets an environment variable with a default value
func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// IsDevelopment checks if the current environment is development
func (e Environment) IsDevelopment() bool {
	return e == Development
}

// IsProduction checks if the current environment is production
func (e Environment) IsProduction() bool {
	return e == Production
}
