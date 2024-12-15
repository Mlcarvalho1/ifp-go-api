package controllers

import (
	"log"

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

func GetCurrentUser(c *fiber.Ctx) error {
	userId := c.Locals("auth").(map[string]interface{})["userId"].(int)

	log.Printf("userId: %v\n", userId)

	user, err := services.GetCurrentUser(&userId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
