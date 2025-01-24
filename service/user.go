package service

import (
	"time"

	"github.com/SuperCodingTeam/model"
	"github.com/SuperCodingTeam/repository"
	"github.com/SuperCodingTeam/utility"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(userID string, password string) utility.Response {
	readUser := repository.ReadUser(userID, "id")
	if readUser == (model.User{}) {
		return utility.Response{StatusCode: 401, Message: "로그인에 실패하였습니다", Error: &model.BookPocketError{Code: "C005", Message: "Not Found"}}
	}

	err := bcrypt.CompareHashAndPassword([]byte(readUser.Password), []byte(password))
	if err != nil {
		return utility.Response{StatusCode: 401, Message: "로그인에 실패하였습니다", Error: &model.BookPocketError{Code: "C004", Message: "UnAuthorized"}}
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userUUID": readUser.UserUUID,
			"exp":      time.Now().Add(time.Minute * 30).Unix(),
		})

		secretKey := []byte("your-secret-key")
		signedToken, err := token.SignedString(secretKey)
		if err != nil {
			return utility.Response{
				StatusCode: 500,
				Message:    "토큰 생성 실패",
				Error:      &model.BookPocketError{Code: "C006", Message: "Token Creation Failed"},
			}
		}
		return utility.Response{StatusCode: 200, Message: "로그인에 성공하였습니다", Token: signedToken}
	}
}

func Register(name string, userID string, password string, birthday time.Time, email string) utility.Response {
	readUser := repository.ReadUser(email, "email")
	if readUser != (model.User{}) {
		return utility.Response{StatusCode: 409, Message: "이미 사용중인 이메일"}
	}

	readUser = repository.ReadUser(userID, "id")
	if readUser != (model.User{}) {
		return utility.Response{StatusCode: 409, Message: "이미 사용중인 아이디"}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return utility.Response{StatusCode: 500, Message: "비밀번호 암호화 실패", Error: &model.BookPocketError{Code: "asd", Message: err.Error()}}
	}

	repository.CreateUser(userID, string(hashedPassword), name, email, &birthday)
	return utility.Response{StatusCode: 201, Message: "회원가입 성공"}
}

func Signout(token string, password string) utility.Response {
	userUUID := utility.JWTDecode(token)
	readUser := repository.ReadUser(userUUID, "user_uuid")
	if readUser == (model.User{}) {
		return utility.Response{StatusCode: 401, Message: "본인확인에 실패하였습니다", Error: &model.BookPocketError{Code: "C005", Message: "Not Found"}}
	}

	err := bcrypt.CompareHashAndPassword([]byte(readUser.Password), []byte(password))
	if err != nil {
		return utility.Response{StatusCode: 401, Message: "본인확인에 실패하였습니다", Error: &model.BookPocketError{Code: "C004", Message: "UnAuthorized"}}
	} else {
		repository.DeleteUserByUUID(userUUID)
		return utility.Response{StatusCode: 200, Message: "회원 탈회 성공공"}
	}
}

func GetProfile(token string) utility.Response {
	userUUID := utility.JWTDecode(token)
	readUser := repository.ReadUser(userUUID, "user_uuid")

	if readUser == (model.User{}) {
		return utility.Response{StatusCode: 404, Message: "해당 유저를 찾을 수 없습니다.", Error: &model.BookPocketError{Code: "C005", Message: "Not Found"}}
	}

	return utility.Response{StatusCode: 200, Message: "정보를 성공적으로 조회하였습니다.", User: readUser}
}

func UpdatePassword(email string, oldPassword string, newPassword string) utility.Response {
	readUser := repository.ReadUser(email, "email")
	if readUser == (model.User{}) {
		return utility.Response{StatusCode: 404, Message: "해당 유저를 찾을 수 없습니다.", Error: &model.BookPocketError{Code: "C005", Message: "Not Found"}}
	}

	err := bcrypt.CompareHashAndPassword([]byte(readUser.Password), []byte(oldPassword))
	if err != nil {
		return utility.Response{StatusCode: 401, Message: "본인확인에 실패하였습니다", Error: &model.BookPocketError{Code: "C004", Message: "UnAuthorized"}}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	readUser.Password = string(hashedPassword)
	repository.UpdateUserByUUID(readUser.UserUUID.String(), readUser)

	return utility.Response{StatusCode: 200, Message: "정보를 성공적으로 업데이트 하였습니다."}
}
