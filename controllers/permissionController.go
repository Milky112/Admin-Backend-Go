package controllers

import (
	"fibr/database"
	"fibr/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermisions(c *fiber.Ctx) error {
	var permissions []models.Permission

	database.DB.Find(&permissions)

	return c.JSON(permissions)
}
