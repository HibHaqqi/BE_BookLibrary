package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	MemberId uint `gorm:"not null"`
	BookId   uint `gorm:"not null"`
	Rating   float64

	// Foreign key references
	Book   Book   `gorm:"foreignKey:BookId"`
	Member Member `gorm:"foreignKey:MemberId"`
}
