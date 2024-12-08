package main

import (
	"booking-service/config"
	"booking-service/database"
	"booking-service/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Database connection
	database.ConnectDB(
		config.GetEnv("DB_HOST", "localhost"),
		config.GetEnv("DB_USER", "postgres"),
		config.GetEnv("DB_PASSWORD", "postgres"),
		config.GetEnv("DB_NAME", "SystemHotelDB"),
		config.GetEnv("DB_PORT", "5432"),
	)

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	port := config.GetEnv("PORT", "8083")
	log.Printf("Booking service running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}