package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juniordeved/go-admin/cmd/database"
	"github.com/juniordeved/go-admin/cmd/routes"
	"log"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))
}
