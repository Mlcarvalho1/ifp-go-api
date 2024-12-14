package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"ifp-analysis.com/config"
)

// Define a struct to hold the authentication configuration.
type AuthConfig struct {
	SecretKey string
}

// Middleware function for Fiber
func AuthMiddleware() fiber.Handler {
	secretKey := config.GetEnv("SECRET_KEY")

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token não existe.",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Formato do token inválido.",
			})
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido.",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("auth", map[string]interface{}{
				"currentUserId": claims["id"],
				"type":          claims["type"],
				"permissions":   claims["permissions"],
			})
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido.",
			})
		}

		return c.Next()
	}
}
