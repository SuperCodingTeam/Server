package repository

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SuperCodingTeam/model"
	"gorm.io/gorm"
)

func checkValidFilter(filter string) bool {
	for _, data := range [3]string{"user_uuid", "id", "email"} {
		if data == filter {
			return true
		}
	}

	return false
}

func CreateUser(id string, password string, email string, name string, birthday time.Time) model.User {
	user := model.User{
		ID:       id,
		Password: password,
		Name:     name,
		Email:    email,
		Birthday: birthday,
	}

	result := db.Create(&user)

	if result.Error != nil {
		log.Printf("Error: %s\n", result.Error)
		return model.User{}
	}

	return user
}

func ReadUser(filter string, data string) model.User {
	var user model.User
	var count int64

	if !checkValidFilter(filter) {
		panic(fmt.Sprintf("%s는 유요한 필터가 아닙니다. 필터는 'user_uuid', 'id', 'email'중 하나입니다.\n", filter))
	}

	result := db.Where(fmt.Sprintf("%s = ?", filter), data).First(&user).Count(&count)
	if result.Error != nil {
		panic(result.Error)
	}

	return user
}

func UpdateUserByUUID(userUUID string, user model.User) model.User {
	updateUser := ReadUser("user_uuid", userUUID)
	updateUser.Name = user.Name
	updateUser.Email = user.Email
	updateUser.Birthday = user.Birthday
	updateUser.Password = user.Password

	db.Save(&updateUser)
	return updateUser
}

func DeleteUserByUUID(userUUID string) {
	user := ReadUser("user_uuid", userUUID)
	result := db.Delete(user)

	if result.Error != nil {
		panic(result.Error)
	}
}

func CheckExistUser(filter string, data string) bool {
	var user model.User
	var count int64
	if checkValidFilter(filter) {
		result := db.Where(fmt.Sprintf("%s = ?", filter), data).First(&user).Count(&count)
		return !errors.Is(result.Error, gorm.ErrRecordNotFound)
	}

	panic(fmt.Sprintf("%s는 유요한 필터가 아닙니다. 필터는 'user_uuid', 'id', 'email'중 하나입니다.\n", filter))
}
