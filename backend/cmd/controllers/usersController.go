package controllers

import (
	"go-admin/cmd/database"
	"go-admin/cmd/models"

	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}
