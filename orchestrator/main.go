package main

import (
	"log"

	"orchestrator/config"
	"orchestrator/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Load environment variables
	config.LoadEnv()

	// inisiasi routes
	routes.SetupRoutes(app)

	// Start server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Orchestrator running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
