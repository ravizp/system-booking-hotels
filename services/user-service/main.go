package main

import (
	"log"
	"os"
	"user-service/database"
	"user-service/models"
	"user-service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi database
	models.InitDatabase()

	// Jalankan migrasi
	database.RunMigrations()

	// Inisialisasi server
	app := fiber.New()
	routes.SetupRoutes(app)
	
	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}
	
	log.Printf("User service running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
