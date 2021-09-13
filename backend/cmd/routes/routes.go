package routes

import (
	"go-admin/cmd/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Authentication
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("logout", controllers.Logout)

}
