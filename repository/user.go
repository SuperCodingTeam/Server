package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/SuperCodingTeam/model"
)

func checkValidFilter(filter string) bool {
	for _, data := range [3]string{"user_uuid", "id", "email"} {
		if data == filter {
			return true
		}
	}

	return false
}

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

func ReadUser(data string, filter string) model.User {
	var user model.User
	var query string

	if !checkValidFilter(filter) {
		return model.User{}
	}

	result := db.Where(fmt.Sprintf("%s = ?", query), data).First(&user)
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
		updateUser.Password = user.Password
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
