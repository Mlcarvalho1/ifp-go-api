package middleware

import "github.com/gofiber/fiber/v2"

func Permissions(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authData := c.Locals("auth")

		if authData == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Não autorizado.",
			})
		}

		return c.Next()
	}
}
