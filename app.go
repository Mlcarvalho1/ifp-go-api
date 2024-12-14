package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"ifp-analysis.com/database"
	"ifp-analysis.com/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	app := fiber.New()

	database.ConnectDb()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
