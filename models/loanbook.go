package models

import "gorm.io/gorm"

type LoanBook struct {
	gorm.Model
	BookID       uint   `gorm:"not null"`
	MemberID     uint   `gorm:"not null"`
	LengthOfLoan int    `gorm:"not null"`
	DueDate      string `gorm:"not null"`

	// Foreign key references
	Book   Book   `gorm:"foreignKey:BookID"`
	Member Member `gorm:"foreignKey:MemberID"`
}
