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
func AuthMiddleware(c *fiber.Ctx) error {
	secretKey := config.GetEnv("SECRET_KEY")
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
		// Validação do userId
		userIdFloat, ok := claims["id"].(float64) // JWT decodifica números como float64
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Formato inválido para o User ID.",
			})
		}
		userId := int(userIdFloat)

		// Validação do type
		userType, ok := claims["type"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Formato inválido para o tipo de usuário.",
			})
		}

		// Validação das permissões
		permissionsInterface, ok := claims["permissions"].([]interface{})
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Formato inválido para as permissões.",
			})
		}

		permissions := []string{}
		for _, permission := range permissionsInterface {
			if permissionStr, ok := permission.(string); ok {
				permissions = append(permissions, permissionStr)
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Permissões devem ser um array de strings.",
				})
			}
		}

		// Armazena os dados no contexto
		c.Locals("auth", map[string]interface{}{
			"userId":      userId,
			"type":        userType,
			"permissions": permissions,
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token inválido.",
		})
	}

	return c.Next()
}
