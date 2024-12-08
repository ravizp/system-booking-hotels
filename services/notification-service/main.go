package main

import (
	"log"
	"notification-service/config"
	"notification-service/database"
	"notification-service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	database.Connect()

	// Create a new Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start the server
	port := config.GetEnv("NOTIFICATION_SERVICE_PORT", "8084")
	log.Printf("Notification service is running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
