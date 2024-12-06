package main

import (
	"log"

	"github.com/raviz/system-booking-hotels/orchestrator/config"
	"github.com/raviz/system-booking-hotels/orchestrator/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()


	// Load env variable
	config.LoadConfig()


	//inisiasi routes
	routes.SetupRoutes(app)

	//start server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Orchestrator running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}