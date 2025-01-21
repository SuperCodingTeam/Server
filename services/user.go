package services

import (
	"fmt"
	"log"
	"time"

	models "github.com/SuperCodingTeam/models"
)

func CreateUser(id string, name string, email string, password string, birthday *time.Time) {
	user := models.User{
		ID:       id,
		Name:     name,
		Password: password,
		Email:    email,
		Birthday: birthday,
	}

	result := db.Create(&user)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}
}

func ReadUser() {
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	} else {
		println(users)
	}
}

func ReadUserByUUID(userUUID string) {
	var user models.User
	result := db.Where("user_uuid = ?", userUUID).First(&user)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}
}

func ReadUserByID(userID string) models.User {
	var user models.User
	result := db.Where("id = ?", userID).First(&user)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}

	return user
}

func UpdateUserByUUID(userUUID string, user models.User) {
	var updateUser models.User
	result := db.Where("user_uuid = ?", userUUID).First(&updateUser)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	} else {
		updateUser.Name = user.Name
		updateUser.Email = user.Email
		updateUser.Birthday = user.Birthday
	}
	db.Save(&user)
}

func DeleteUserByUUID(userUUID string) {
	var user models.User
	result := db.Delete(user)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}
}
