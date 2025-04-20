package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	Floor       int `gorm:"not null"`
}
