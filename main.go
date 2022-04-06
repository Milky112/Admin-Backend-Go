package main

import (
	"fibr/database"
	"fibr/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	

	app.Listen("127.0.0.1:3000")
}
