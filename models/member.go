package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
	Phone string
}
