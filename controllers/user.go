package controllers

import (
	"time"

	"github.com/SuperCodingTeam/models"
	"github.com/SuperCodingTeam/services"
	"github.com/SuperCodingTeam/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(userID string, password string) utils.Response {
	readUser := services.ReadUserByID(userID)
	if readUser == (models.User{}) {
		return utils.Response{StatusCode: 401, Message: "로그인에 실패하였습니다", Error: &models.BookPocketError{Code: "C005", Message: "Not Found"}}
	}

	err := bcrypt.CompareHashAndPassword([]byte(readUser.Password), []byte(password))
	if err != nil {
		return utils.Response{StatusCode: 401, Message: "로그인에 실패하였습니다", Error: &models.BookPocketError{Code: "C004", Message: "UnAuthorized"}}
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userUUID": readUser.UserUUID,
			"exp":      time.Now().Add(time.Minute * 30).Unix(),
		})

		secretKey := []byte("your-secret-key")
		signedToken, err := token.SignedString(secretKey)
		if err != nil {
			return utils.Response{
				StatusCode: 500,
				Message:    "토큰 생성 실패",
				Error:      &models.BookPocketError{Code: "C006", Message: "Token Creation Failed"},
			}
		}
		return utils.Response{StatusCode: 200, Message: "로그인에 성공하였습니다", Token: signedToken}
	}
}

func Register() {}

func Signout() {}

func GetProfile() {}

func UpdatePassword() {}

func UpdateEmail() {}

// admin admin
