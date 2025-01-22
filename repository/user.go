package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/SuperCodingTeam/model"
)

func CreateUser(id string, password string, name string, email string, birthday *time.Time) {
	user := model.User{
		ID:       id,
		Password: password,
		Name:     name,
		Email:    email,
		Birthday: birthday,
	}

	result := db.Create(&user)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}
}

func ReadUser() []model.User {
	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}

	return users
}

func ReadUserByUUID(userUUID string) model.User {
	var user model.User
	result := db.Where("user_uuid = ?", userUUID).First(&user)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}

	return user
}

func ReadUserByID(userID string) model.User {
	var user model.User
	result := db.Where("id = ?", userID).First(&user)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}

	return user
}

func UpdateUserByUUID(userUUID string, user model.User) {
	var updateUser model.User
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
	var user model.User
	result := db.Delete(user)

	if result.Error != nil {
		log.Println(fmt.Sprintf("Error: %s", result.Error))
	}
}
