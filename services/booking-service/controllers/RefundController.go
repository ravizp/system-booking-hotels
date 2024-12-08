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

	// Cari booking di database
	var booking models.Booking
	if err := database.DB.First(&booking, req.BookingID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Booking not found"})
	}

	// Buat refund baru
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

func GetAllRefunds(c *fiber.Ctx) error {
	var refunds []models.Refund
	if err := database.DB.Find(&refunds).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch refunds"})
	}

	return c.JSON(refunds)
}

func UpdateRefundStatus(c *fiber.Ctx) error {
	type UpdateStatusRequest struct {
		Status string `json:"status"` // approved, rejected
	}

	refundID := c.Params("id")
	var req UpdateStatusRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var refund models.Refund
	if err := database.DB.First(&refund, refundID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Refund not found"})
	}

	// Update status dan waktu approval jika disetujui
	refund.Status = req.Status
	if req.Status == "approved" {
		now := time.Now()
		refund.ApprovedDate = &now
	}

	if err := database.DB.Save(&refund).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update refund status"})
	}

	return c.JSON(fiber.Map{"message": "Refund status updated", "refund": refund})
}
