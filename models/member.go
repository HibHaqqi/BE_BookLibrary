package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Phone    string
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
}
