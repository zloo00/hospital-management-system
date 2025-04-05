package config

import (
	"github.com/aluazholdykan/hospital-management-system/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("hospital.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Patient{})
	db.AutoMigrate(&models.Patient{}, &models.Doctor{})

	DB = db
}
