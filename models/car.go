package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	RegistrationNumber string
	CarModel           string
	Color              string
}
