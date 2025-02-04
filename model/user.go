package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserUUID  uuid.UUID `gorm:"primaryKey;type:char(36);not null" json:"user_uuid"`
	ID        string    `gorm:"size:15;unique;not null" json:"id"`
	Password  string    `gorm:"size:64;not null" json:"-"`
	Name      string    `gorm:"size:10;not null" json:"name"`
	Email     string    `gorm:"size:40;unique;not null" json:"email"`
	Birthday  time.Time `gorm:"type:DATETIME" json:"birthday"`
	CreatedAt time.Time `gorm:"type:DATETIME;default:NOW()" json:"created_at"`
}

type LoginRequestType struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
}

type RegisterRequestType struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

type SignoutRequestType struct {
	Password string `json:"password"`
}

type UpdateRequestType struct {
	UserID   string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if user.UserUUID == uuid.Nil {
		user.UserUUID = uuid.New()
	}
	return
}
