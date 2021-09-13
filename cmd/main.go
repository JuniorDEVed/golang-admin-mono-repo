package main

import (
	"log"

	"go-admin/cmd/database"
	"go-admin/cmd/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
