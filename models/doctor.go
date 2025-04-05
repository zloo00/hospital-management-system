package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
