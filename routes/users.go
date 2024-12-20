package routes

import (
	"github.com/gofiber/fiber/v2"
	"ifp-analysis.com/controllers"
	middleware "ifp-analysis.com/middlewares"
)

func UserRoutes(api fiber.Router) {
	user := api.Group("/user")

	api.Use(middleware.AuthMiddleware)

	user.Get("/", controllers.ListAllUsers)
	user.Get("/info", controllers.GetCurrentUser)
}
