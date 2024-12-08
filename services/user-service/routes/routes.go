package routes

import (
	"user-service/controllers"
	// "github.com/ravizp/system-booking-hotels/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/users")

	//public routes
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	//prorected routes
	// api.Use(middlewares.Authentication)
	api.Get("/", controllers.GetAllUsers) // Endpoint untuk mendapatkan seluruh user
	api.Get("/:id", controllers.GetUserByID) // Endpoint untuk mendapatkan user berdasarkan ID
}