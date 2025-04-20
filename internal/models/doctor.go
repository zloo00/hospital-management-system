package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	UserID         uint   `gorm:"unique;not null"`
	User           User   `gorm:"foreignKey:UserID"`
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	Specialization string `gorm:"not null"`
	RoomNumber     string `gorm:"not null"`
	PhoneNumber    string
	DepartmentID   uint
	Department     Department `gorm:"foreignKey:DepartmentID"`
}
