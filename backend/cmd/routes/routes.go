package routes

import (
	"go-admin/cmd/controllers"
	"go-admin/cmd/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// Auth middleware ---> routes below must be authenticated
	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users/", controllers.AllUsers)
}
