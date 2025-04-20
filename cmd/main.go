package main

import (
	"hospital-app/internal/config"
	"hospital-app/internal/models" // Подключаем модели
	"hospital-app/internal/routes"
	"log"
)

func main() {
	// Initialize configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations to ensure that the required tables exist
	err = db.AutoMigrate(&models.User{}) // Миграция для модели User
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize router
	router := routes.SetupRouter(db, cfg)

	// Run the server
	log.Printf("Server is running on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
