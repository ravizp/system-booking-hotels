package routes

import (
	"user-service/controllers"

	// "user-service/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/users")

	// Public routes
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	// api yang kena middleware belum jalan masih kena authorization
	// selebihnya aman ga saya lanjutin authenttcation untuk endpoint lainnya

	// api.Use(middlewares.Authentication)
	api.Post("/logout", controllers.Logout)
	api.Get("/", controllers.GetAllUsers)    // Endpoint untuk mendapatkan seluruh user
	api.Get("/:id", controllers.GetUserByID) // Endpoint untuk mendapatkan user berdasarkan ID
}
