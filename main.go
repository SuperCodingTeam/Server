package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"

	"github.com/SuperCodingTeam/controller"
	"github.com/SuperCodingTeam/database"
	model "github.com/SuperCodingTeam/model"
	"github.com/SuperCodingTeam/service"
	"github.com/SuperCodingTeam/utility"
)

func main() {
	fastergoding.Run()
	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	db := database.ConnectDatabase()

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Book{})

	app.Get("/:title", controller.GetBookByTitle)
	result := service.Login("admin", "admin")
	(utility.JWTDecode(result.Token))

	log.Fatal(app.Listen(":8080"))

}
