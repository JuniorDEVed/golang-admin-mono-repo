package routes

import (
	"go-admin/cmd/controllers"
	"go-admin/cmd/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	// Auth middleware ---> routes below must be authenticated
	app.Use(middlewares.IsAuthenticated)

	app.Get("/user", controllers.User)
	app.Post("logout", controllers.Logout)

}
