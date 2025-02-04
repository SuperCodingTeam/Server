package controller

import (
	"github.com/SuperCodingTeam/service"
	"github.com/gofiber/fiber/v2"
)

// 책목록 조회
//
//	@Summary		책목록을 조회합니다.
//	@Description	책 목록을 query와 target을 통해 불러옵니다. target은 title, isbn, publisher, person이 있으며, query는 검색어입니다.
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			query	query		string	true	"query를 통해 검색어 입력"
//	@Param			target	query		string	true	"target을 통해 검색 타입 지정"
//	@Success		200		{object}	utility.OKResponse
//	@Failure		401		{object}	utility.FailResponse
//	@Failure		500		{object}	utility.FailResponse
//	@Router			/ [get]
func GetBookController(c *fiber.Ctx) error {
	m := c.Queries()
	query := m["query"]
	target := m["target"]
	if query == "" || target == "" {
		return c.JSON(fiber.Map{"message": "query와 target 필드는 필수입니다!", "code": 404})
	}

	return c.JSON(fiber.Map{"message": "데이터를 성공적으로 불러왔습니다!", "code": 200, "data": service.GetBookByKeyword(query, target)})
}
