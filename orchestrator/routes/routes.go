package routes

import (
	"orchestrator/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.All("/users/*", handlers.ProxyUserService)
	app.All("/hotels/*", handlers.ProxyHotelService)
	app.All("/bookings/*", handlers.ProxyBookingService)
}