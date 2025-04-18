package models

import (
	"github.com/jinzhu/gorm"
)

// User модель для регистрации и авторизации пользователей
type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"` // 'patient', 'doctor', 'admin'
}

// Doctor модель для врачей
type Doctor struct {
	gorm.Model
	UserID         uint   `gorm:"not null"`
	Specialization string `gorm:"not null"`
	Hospital       string `gorm:"not null"`
}

// Appointment модель для записей на приём
type Appointment struct {
	gorm.Model
	PatientID uint   `gorm:"not null"`
	DoctorID  uint   `gorm:"not null"`
	Date      string `gorm:"not null"`
	Symptoms  string
	Status    string `gorm:"not null"` // 'scheduled', 'completed', 'canceled'
}
