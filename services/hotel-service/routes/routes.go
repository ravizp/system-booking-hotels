package routes

import (
	"hotel-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/hotels", controllers.GetHotels)
	app.Get("/hotels/:id", controllers.GetHotel)
	app.Post("/hotels", controllers.CreateHotel)
	app.Post("/rooms", controllers.CreateRoom)
}
