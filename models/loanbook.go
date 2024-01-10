package models

type LoanBook struct {
	Id            int64  `gorm:"primaryKey" json:"id"`
	MemberId      string `gorm:"type:varchar(300)" json:"member_id"`
	BookId        string `gorm:"type:varchar(300)" json:"book_id"`
	LenghtofLoan  string `gorm:"type:varchar(300)" json:"length_of_loan"`
	DueDateofLoan string `gorm:"type:date" json:"due_date_of_loan"`
}
