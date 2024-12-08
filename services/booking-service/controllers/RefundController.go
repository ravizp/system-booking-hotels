package controllers

import (
	"booking-service/database"
	"booking-service/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RefundBooking(c *fiber.Ctx) error {
	type RefundRequest struct {
		BookingID uint   `json:"bookingId"`
		Reason    string `json:"reason"`
	}

	var req RefundRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var booking models.Booking
	if err := database.DB.First(&booking, req.BookingID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Booking not found"})
	}

	//contoh pengajuan refund dari total pembayaran
	refund := models.Refund{
		BookingID:   req.BookingID,
		Amount:      100, 
		Reason:      req.Reason,
		Status:      "pending",
		RequestDate: time.Now(),
	}

	if err := database.DB.Create(&refund).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create refund"})
	}

	return c.JSON(fiber.Map{"message": "Refund requested", "refund": refund})
}
