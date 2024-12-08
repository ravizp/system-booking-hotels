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
	app.All("/hotels/*", handlers.ProxyHotelService)
	app.All("/bookings/*", handlers.ProxyBookingService)
	app.All("/notifications/*", handlers.ProxyNotificationService)
}
