package models

import "gorm.io/gorm"

type Role string

const (
	AdminRole   Role = "admin"
	DoctorRole  Role = "doctor"
	PatientRole Role = "patient"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     Role   `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"unique;not null"`
}
