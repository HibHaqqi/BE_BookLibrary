package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Author      string `gorm:"not null"`
	PublishYear int    `gorm:"not null"`
	Description string
}
