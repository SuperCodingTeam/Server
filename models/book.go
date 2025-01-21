package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	BookUUID  uuid.UUID `gorm:"primaryKey;type:char(36);not null"`
	BookName  string    `gorm:"not null"`
	Author    string    `gorm:"not null"`
	ImagePath string    `gorm:"null;default:null"`
}

func (book *Book) BeforeSave(tx *gorm.DB) (err error) {
	if book.BookUUID == uuid.Nil {
		book.BookUUID = uuid.New()
	}
	return nil
}
