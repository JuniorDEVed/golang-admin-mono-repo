package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/juniordeved/go-admin/cmd/database"
	"github.com/juniordeved/go-admin/cmd/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	log.Fatal(app.Listen(":8000"))
}
