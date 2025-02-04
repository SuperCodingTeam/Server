package service

import (
	"fmt"
	"time"

	"github.com/SuperCodingTeam/model"
	"github.com/SuperCodingTeam/repository"
	"github.com/SuperCodingTeam/utility"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(userID string, password string) (resp interface{}) {
	defer func() {
		if r := recover(); r != nil {
			resp = utility.FailResponse{Status: "서버 내부 오류", Code: 500, Message: "서버 내부 오류가 발생하였습니다.", Error: fmt.Sprintf("%v", r)}
		}
	}()

	if userID == "" || password == "" {
		log.Warn(fmt.Sprintf("id는 `%s`, password는 `%s`의 값이 전달됐습니다.\n", userID, password))
		return utility.FailResponse{Status: "처리할 수 없는 엔티티", Code: 422, Message: "데이터가 잘못 전송됐습니다.", Error: fmt.Sprintf("id는 `%s`, password는 `%s`의 값이 전달됐습니다.", userID, password)}
	}

	readUser := repository.ReadUser("id", userID)
	if readUser == (model.User{}) {
		return utility.FailResponse{Status: "실패", Code: 401, Message: "로그인에 실패하였습니다", Error: fmt.Sprintf("`%s`라는 id를 가진 유저를 찾을 수 없습니다.", userID)}
	}

	err := bcrypt.CompareHashAndPassword([]byte(readUser.Password), []byte(password))
	if err != nil {
		return utility.FailResponse{Status: "실패", Code: 401, Message: "로그인에 실패하였습니다", Error: "비밀번호가 일치하지 않습니다."}
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userUUID": readUser.UserUUID,
			"exp":      time.Now().Add(time.Minute * 30).Unix(),
		})

		secretKey := []byte("your-secret-key")
		signedToken, err := token.SignedString(secretKey)
		if err != nil {
			return utility.FailResponse{
				Status:  "서버 내부 오류",
				Code:    500,
				Message: "토큰 생성 실패",
				Error:   err.Error(),
			}
		}
		return utility.LoginResponse{Status: "성공", Code: 200, Message: "로그인에 성공하였습니다", Token: signedToken}
	}
}

func Register(userID string, password string, email string, name string, birthday time.Time) (resp interface{}) {
	defer func() {
		if r := recover(); r != nil {
			resp = utility.FailResponse{Status: "서버 내부 오류", Code: 500, Message: "서버 내부 오류가 발생하였습니다.", Error: fmt.Sprintf("%v", r)}
		}
	}()

	if userID == "" || password == "" || email == "" || name == "" {
		return utility.FailResponse{Status: "처리할 수 없는 엔티티", Code: 422, Message: "데이터가 잘못 전송됐습니다.", Error: fmt.Sprintf("id는 `%s`, password는 `%s`, name은 `%s`, email은 `%s`의 값이 전달됐습니다.", userID, password, name, email)}
	}
	if repository.CheckExistUser("email", email) {
		return utility.FailResponse{Status: "데이터 충돌", Code: 409, Message: "이미 사용중인 이메일입니다.", Error: fmt.Sprintf("`%s`라는 이메일을 가진 유저가 이미 존재합니다.", email)}
	}

	if repository.CheckExistUser("id", userID) {
		return utility.FailResponse{Status: "데이터 충돌", Code: 409, Message: "이미 사용중인 아이디입니다.", Error: fmt.Sprintf("`%s`라는 아이디를 가진 유저가 이미 존재합니다.", userID)}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return utility.FailResponse{Status: "서버 내부 오류", Code: 500, Message: "서버 내부 오류가 발생하였습니다.", Error: err.Error()}
	}

	repository.CreateUser(userID, string(hashedPassword), email, name, birthday)
	return utility.OKResponse{Status: "생성됨", Code: 201, Message: "회원가입에 성공하였습니다."}
}

func Signout(token string, password string) (resp interface{}) {
	defer func() {
		if r := recover(); r != nil {
			resp = utility.FailResponse{Status: "서버 내부 오류", Code: 500, Message: "서버 내부 오류가 발생하였습니다.", Error: fmt.Sprintf("%v", r)}
		}
	}()

	userUUID := utility.JWTDecode(token)
	readUser := repository.ReadUser("user_uuid", userUUID)
	if readUser == (model.User{}) {
		return utility.FailResponse{Status: "실패", Code: 401, Message: "본인확인에 실패하였습니다", Error: "해당 토큰을 사용하는 유저가 없습니다."}
	}

	err := bcrypt.CompareHashAndPassword([]byte(readUser.Password), []byte(password))
	if err != nil {
		return utility.FailResponse{Status: "실패", Code: 401, Message: "본인확인에 실패하였습니다", Error: "비밀번호가 일치하지 않습니다."}
	} else {
		repository.DeleteUserByUUID(userUUID)
		return utility.OKResponse{Status: "성공", Code: 200, Message: "성공적으로 회원을 탈퇴하였습니다."}
	}
}

func GetProfile(token string) (resp interface{}) {
	defer func() {
		if r := recover(); r != nil {
			resp = utility.FailResponse{Status: "서버 내부 오류", Code: 500, Message: "서버 내부 오류가 발생하였습니다.", Error: fmt.Sprintf("%v", r)}
		}
	}()

	userUUID := utility.JWTDecode(token)
	readUser := repository.ReadUser("user_uuid", userUUID)

	if token == "" {
		return utility.FailResponse{Status: "처리할 수 없는 엔티티", Code: 422, Message: "데이터가 잘못 전송됐습니다.", Error: fmt.Sprintf("토큰은 `%s`의 값이 전달됐습니다.", token)}
	}

	if readUser == (model.User{}) {
		return utility.FailResponse{Status: "찾을 수 없음", Code: 404, Message: "프로필 조회에 실패하였습니다.", Error: "해당 토큰을 사용하는 유저가 없습니다."}
	}

	return utility.ProfileResponse{Status: "성공", Code: 200, Message: "정보를 성공적으로 조회하였습니다.", User: readUser}
}

func UpdatePassword(userID string, email string, newPassword string) (resp interface{}) {
	defer func() {
		if r := recover(); r != nil {
			resp = utility.FailResponse{Status: "서버 내부 오류", Code: 500, Message: "서버 내부 오류가 발생하였습니다.", Error: fmt.Sprintf("%v", r)}
		}
	}()

	readUser := repository.ReadUser("id", userID)
	if readUser == (model.User{}) {
		return utility.FailResponse{Status: "찾을 수 없음", Code: 404, Message: "해당 유저를 찾을 수 없습니다.", Error: fmt.Sprintf("`%s`라는 이메일을 가진 유저가 이미 존재합니다.", email)}
	}

	if readUser.Email != email {
		return utility.FailResponse{Status: "금지됨", Code: 403, Message: "접근 권한이 없습니다.", Error: fmt.Sprintf("`%s`의 유저는 `%s`의 이메일을 갖고 있습니다. 입력된 `%s`의 이메일과 다릅니다", readUser.ID, readUser.Email, email)}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return utility.FailResponse{Status: "서버 내부 오류", Code: 500, Message: "서버 내부 오류가 발생하였습니다.", Error: err.Error()}
	}

	readUser.Password = string(hashedPassword)
	repository.UpdateUserByUUID(readUser.UserUUID.String(), readUser)

	return utility.OKResponse{Status: "성공", Code: 200, Message: "정보를 성공적으로 업데이트 하였습니다."}
}

func CheckValidateData(filter string, data string) (resp interface{}) {
	defer func() {
		if r := recover(); r != nil {
			resp = utility.FailResponse{Status: "서버 내부 오류", Code: 500, Message: "서버 내부 오류가 발생하였습니다.", Error: fmt.Sprintf("%v", r)}
		}
	}()

	if repository.CheckExistUser(filter, data) {
		return utility.FailResponse{Status: "데이터 충돌", Code: 409, Message: "이미 사용중인 데이터입니다.", Error: fmt.Sprintf("`%s` 값을 가진 `%s`는 중복되었습니다.", data, filter)}
	} else {
		return utility.OKResponse{Status: "성공", Code: 200, Message: "사용할 수 있는 데이터입니다."}
	}
}
