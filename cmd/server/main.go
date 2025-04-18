package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Инициализация Gin
	r := gin.Default()

	// Простая проверка
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "MediGo API is running",
		})
	})

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // по умолчанию порт 8080
	}

	r.Run(":" + port)
}
