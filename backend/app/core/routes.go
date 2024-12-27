package core

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", healthCheck)
}

func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "alive"})
}
