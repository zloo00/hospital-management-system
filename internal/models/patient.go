package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	UserID      uint   `gorm:"unique;not null"`
	User        User   `gorm:"foreignKey:UserID"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	DateOfBirth string `gorm:"not null"`
	Gender      string `gorm:"not null"`
	Address     string
	PhoneNumber string
}
