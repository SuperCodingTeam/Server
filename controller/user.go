package controller

import (
	"fmt"
	"time"

	"github.com/SuperCodingTeam/model"
	"github.com/SuperCodingTeam/service"
	"github.com/SuperCodingTeam/utility"
	"github.com/gofiber/fiber/v2"
)

// @Summary		로그인
// @Description	로그인 후 토큰을 반환합니다.
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			request body		model.LoginRequestType	true	"로그인을 위해 필요한 유저 아이디 및 암호화 되지 않은 비밀번호"
// @Success		200		{object}	utility.LoginResponse
// @Failure		401		{object}	utility.FailResponse
// @Failure		422		{object}	utility.FailResponse
// @Failure		500		{object}	utility.FailResponse
// @Router			/login [post]
func Login(c *fiber.Ctx) error {
	var request model.LoginRequestType
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(utility.FailResponse{Status: "처리할 수 없는 엔티티", Code: 422, Message: "데이터가 잘못 전송됐습니다.", Error: fmt.Sprintf("id는 `%s`, password는 `%s`의 값이 전달되었습니다.", request.UserID, request.Password)})
	}

	result := service.Login(request.UserID, request.Password)
	fmt.Println(result)

	return c.JSON(service.Login(request.UserID, request.Password))
}

// @Summary		회원가입
// @Description	회원가입을 통해 데이터베이스에 유저 객체를 삽입합니다.
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			request body		model.RegisterRequestType	true	"회원가입을 위해 필요한 유저 아이디, 암호화 되지 않은 비밀번호, 이메일, 이름 및 생일"
// @Success		201		{object}	utility.OKResponse
// @Failure		401		{object}	utility.FailResponse
// @Failure		409		{object}	utility.FailResponse
// @Failure		422		{object}	utility.FailResponse
// @Failure		500		{object}	utility.FailResponse
// @Router			/user [post]
func Register(c *fiber.Ctx) error {
	var request model.RegisterRequestType
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(utility.FailResponse{Status: "처리할 수 없는 엔티티", Code: 422, Message: "데이터가 잘못 전송됐습니다.", Error: fmt.Sprintf("id는 `%s`, password는 `%s`, name은 `%s`, email은 `%s`의 값이 전달됐습니다.", request.UserID, request.Password, request.Name, request.Email)})
	}
	day, err := time.Parse("2006-01-02", request.Birthday)
	if err != nil {
		return c.JSON(utility.FailResponse{Status: "서버 내부 오류", Code: 500, Message: "서버 내부 오류가 발생하였습니다", Error: err.Error()})
	}
	return c.JSON(service.Register(request.UserID, request.Password, request.Email, request.Name, day))
}

// @Summary		프로필 조회
// @Description	유저의 프로필 정보를 token을 통해 조회합니다.
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			token	query		string	true	"token을 통해 유저 조회"
// @Success		200		{object}	utility.ProfileResponse
// @Failure		401		{object}	utility.FailResponse
// @Failure		422		{object}	utility.FailResponse
// @Failure		500		{object}	utility.FailResponse
// @Router			/user [get]
func GetProfileController(c *fiber.Ctx) error {
	token := c.Queries()["token"]
	return c.JSON(service.GetProfile(token))
}

// @Summary		회원 탈퇴
// @Description	유저 정보를 제거 합니다, 토큰과 비밀번호를 입력받습니다.
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			token	query		string	true	"token을 통해 유저 조회"
// @Param			request body		model.SignoutRequestType 	true 	"비밀번호를 통해 본인 확인"
// @Success		200		{object}	utility.OKResponse
// @Failure		401		{object}	utility.FailResponse
// @Failure		422		{object}	utility.FailResponse
// @Failure		500		{object}	utility.FailResponse
// @Router			/user [delete]
func SignoutController(c *fiber.Ctx) error {
	var request model.SignoutRequestType
	token := c.Queries()["token"]
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(utility.FailResponse{Status: "처리할 수 없는 엔티티", Code: 422, Message: "데이터가 잘못 전송됐습니다.", Error: fmt.Sprintf("password는 `%s`의 값이 전달되었습니다.", request.Password)})
	}
	return c.JSON(service.Signout(token, request.Password))
}

// @Summary		비밀번호 변경
// @Description	비밀번호를 잊었을 경우 변경합니다.
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			request body		model.UpdateRequestType 	true 	"비밀번호를 통해 본인 확인"
// @Success		200		{object}	utility.OKResponse
// @Failure		401		{object}	utility.FailResponse
// @Failure		422		{object}	utility.FailResponse
// @Failure		500		{object}	utility.FailResponse
// @Router			/user/forgot/password [patch]
func UpdateUserController(c *fiber.Ctx) error {
	var request model.UpdateRequestType
	if err := c.BodyParser(&request); err != nil {
		return c.JSON(utility.FailResponse{Status: "처리할 수 없는 엔티티", Code: 422, Message: "데이터가 잘못 전송됐습니다.", Error: fmt.Sprintf("password는 `%s`의 값이 전달되었습니다.", request.Password)})
	}
	return c.JSON(service.UpdatePassword(request.UserID, request.Email, request.Password))
}

// @Summary		데이터 중복 검사
// @Description	id 또는 email중 하나만 검사하시길 바랍니다.
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			filter	query		string	true	"id 또는 email중 하나"
// @Param			data	query		string	true	"중복 검사할 데이터"
// @Success		200		{object}	utility.OKResponse
// @Failure		409		{object}	utility.FailResponse
// @Failure		422		{object}	utility.FailResponse
// @Failure		500		{object}	utility.FailResponse
// @Router			/user/check/validate [post]
func CheckValidate(c *fiber.Ctx) error {
	m := c.Queries()
	filter := m["filter"]
	data := m["data"]

	return c.JSON(service.CheckValidateData(filter, data))
}
