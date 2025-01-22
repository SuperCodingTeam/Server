package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserUUID  uuid.UUID  `gorm:"primaryKey;type:char(36);not null"`
	ID        string     `gorm:"size:15;unique;not null"`
	Password  string     `form:"size:64;not null"`
	Name      string     `gorm:"size:10;not null"`
	Email     string     `gorm:"size:40;unique;not null"`
	Birthday  *time.Time `gorm:"type:DATETIME"`
	CreatedAt time.Time  `gorm:"type:DATETIME;default:NOW()"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if user.UserUUID == uuid.Nil {
		user.UserUUID = uuid.New()
	}
	return
}
