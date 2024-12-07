package main

import (
	"hotel-service/config"
	"hotel-service/database"
	"hotel-service/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()

	database.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	port := config.GetEnv("PORT", "8082")
	log.Printf("Hotel service running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}