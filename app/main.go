package main

import (
	"order-management/config"
	"order-management/routes"
	"github.com/gofiber/fiber/v2"


)

func main() {
    app := fiber.New()

    config.Connectdb()

	routes.UserRoutes(app)

    app.Listen(":8000")
}