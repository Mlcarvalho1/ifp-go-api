package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Setup middlewares

	api := app.Group("/", logger.New())

	UserRoutes(api)
}
