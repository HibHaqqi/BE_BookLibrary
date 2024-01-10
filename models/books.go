package models

type Book struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"type:varchar(300)" json:"title"`
	Author      string `gorm:"type:varchar(300)" json:"author"`
	PublishYear string `gorm:"type:varchar(300)" json:"publish_year"`
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`
}
