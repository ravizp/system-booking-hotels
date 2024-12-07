package routes

import (
	"orchestrator/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "Orchestrator running!"})
	})

	app.All("/users/*", handlers.ProxyUserService)
	app.All("/hotels/*", func(c *fiber.Ctx) error {
		targetURL := "http://hotel-service:8082" + c.OriginalURL()
		return c.Redirect(targetURL, fiber.StatusTemporaryRedirect)
	})
	
}