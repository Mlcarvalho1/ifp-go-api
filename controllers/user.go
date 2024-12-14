package controllers

import (
	"github.com/gofiber/fiber/v2"
	"ifp-analysis.com/services"
)

func ListAllUsers(c *fiber.Ctx) error {
	user, err := services.ListAllUsers()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
