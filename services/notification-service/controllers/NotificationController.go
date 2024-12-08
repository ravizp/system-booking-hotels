package controllers

import (
	"notification-service/database"
	"notification-service/models"

	"github.com/gofiber/fiber/v2"
)

// SendNotification handles creating a new notification
func SendNotification(c *fiber.Ctx) error {
	var payload models.Notification
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := database.DB.Create(&payload).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create notification"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Notification created successfully", "data": payload})
}

// GetAllNotifications retrieves all notifications for a user
func GetAllNotifications(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User ID is required"})
	}

	var notifications []models.Notification
	if err := database.DB.Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch notifications"})
	}

	return c.JSON(fiber.Map{"data": notifications})
}

// DeleteNotification deletes a notification by ID
func DeleteNotification(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Notification{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete notification"})
	}

	return c.JSON(fiber.Map{"message": "Notification deleted successfully"})
}
