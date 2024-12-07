package routes

import (
	"user-service/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/ravizp/system-booking-hotels/shared/middlewares"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/users")

	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	api.Get("/", middlewares.AdminOnly, controllers.GetAllUsers)   // Only admin can access
	api.Get("/:id", middlewares.AdminOnly, controllers.GetUserByID) // Endpoint untuk mendapatkan user berdasarkan ID
}