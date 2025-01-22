package controller

import (
	"fmt"
	"net/url"

	"github.com/SuperCodingTeam/utility"
	"github.com/gofiber/fiber/v2"
)

func GetBookByTitle(c *fiber.Ctx) error {
	title := c.Params("title")
	decodeTitle, _ := url.QueryUnescape(title)
	fmt.Println(decodeTitle)

	return c.JSON(fiber.Map{"data": utility.GetBookByKeyword(decodeTitle)})
}
