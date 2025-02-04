package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"

	"github.com/SuperCodingTeam/controller"
	"github.com/SuperCodingTeam/database"
	model "github.com/SuperCodingTeam/model"

	_ "github.com/SuperCodingTeam/docs"
)

func main() {
	fastergoding.Run()
	app := fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})
	db := database.ConnectDatabase()

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Book{})

	app.Get("/", controller.GetBookController)

	app.Post("/login", controller.Login)
	app.Get("/user", controller.GetProfileController)
	app.Post("/user", controller.Register)
	app.Patch("/user", controller.UpdateUserController)
	app.Delete("/user", controller.SignoutController)
	app.Post("/user/check/validate", controller.CheckValidate)

	app.Use(swagger.New(swagger.Config{
		BasePath: "/api/v1",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		CacheAge: 0,
	}))

	log.Fatal(app.Listen(":8080"))

}
