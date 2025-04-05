package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
}
