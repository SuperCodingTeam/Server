package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/SuperCodingTeam/database"
	"github.com/SuperCodingTeam/models"
)

func main() {
	app := fiber.New()
	db := database.ConnectDatabase()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})
	log.Fatal(app.Listen(":8080"))
}
