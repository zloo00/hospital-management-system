package models

import (
	"gorm.io/gorm"
	"time"
)

type Appointment struct {
	gorm.Model
	PatientID       uint       `gorm:"not null"`
	Patient         Patient    `gorm:"foreignKey:PatientID"`
	DoctorID        uint       `gorm:"not null"`
	Doctor          Doctor     `gorm:"foreignKey:DoctorID"`
	DepartmentID    uint       `gorm:"not null"`
	Department      Department `gorm:"foreignKey:DepartmentID"`
	AppointmentDate time.Time  `gorm:"not null"`
	Status          string     `gorm:"not null"` // scheduled, completed, canceled
	Diagnosis       string
	Prescription    string
	Notes           string
}
