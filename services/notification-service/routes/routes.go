package routes

import (
	"notification-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/notifications", controllers.SendNotification)
	app.Get("/notifications", controllers.GetAllNotifications)
	app.Delete("/notifications/:id", controllers.DeleteNotification)
}
