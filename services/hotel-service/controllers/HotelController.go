package controllers

import (
	"hotel-service/database"
	"hotel-service/models"

	"github.com/gofiber/fiber/v2"
)

func GetHotels(c *fiber.Ctx) error {
	var hotels []models.Hotel
	database.DB.Preload("Rooms").Find(&hotels)
	return c.JSON(hotels)
}

func GetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	var hotel models.Hotel
	result := database.DB.Preload("Rooms").First(&hotel, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Hotel not found"})
	}
	return c.JSON(hotel)
}

func CreateHotel(c *fiber.Ctx) error {
	var hotel models.Hotel
	if err := c.BodyParser(&hotel); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Create(&hotel)
	return c.JSON(hotel)
}

func CreateRoom(c *fiber.Ctx) error {
	var room models.Room
	if err := c.BodyParser(&room); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	database.DB.Create(&room)
	return c.JSON(room)
}