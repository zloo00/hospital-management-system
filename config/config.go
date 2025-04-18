package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Подключение к PostgreSQL
)

// Структура для хранения конфигурации
type Config struct {
	DB_URL     string
	JWT_SECRET string
	DB         *gorm.DB
}

// Функция для загрузки конфигурации из переменных окружения
func LoadConfig() *Config {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is required but not set in .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required but not set in .env file")
	}

	// Подключение к базе данных
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	return &Config{
		DB_URL:     dbURL,
		JWT_SECRET: jwtSecret,
		DB:         db,
	}
}
