package repository

import (
	"github.com/SuperCodingTeam/database"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = database.ConnectDatabase()
}
