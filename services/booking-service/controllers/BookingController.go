package controllers

import (
	"booking-service/database"
	"booking-service/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateBooking(c *fiber.Ctx) error {
	var input struct {
		UserID       uint   `json:"userId"`
		RoomID       uint   `json:"roomId"`
		CheckInDate  string `json:"check_in_date"`  // Input string dari JSON
		CheckOutDate string `json:"check_out_date"` // Input string dari JSON
	}

	// Parse JSON body
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input format"})
	}

	// Validate required fields
	if input.UserID == 0 || input.RoomID == 0 || input.CheckInDate == "" || input.CheckOutDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	// Parse and validate dates
	checkInDate, err := time.Parse("2006-01-02", input.CheckInDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid check_in_date format"})
	}

	checkOutDate, err := time.Parse("2006-01-02", input.CheckOutDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid check_out_date format"})
	}

	if !checkInDate.Before(checkOutDate) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Check-in date must be before check-out date"})
	}

	// Create booking
	booking := models.Booking{
		UserID:       input.UserID,
		RoomID:       input.RoomID,
		CheckInDate:  checkInDate,
		CheckOutDate: checkOutDate,
		Status:       "pending",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := database.DB.Create(&booking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create booking"})
	}

	// Return created booking
	return c.Status(fiber.StatusCreated).JSON(booking)
}

func GetBookings(c *fiber.Ctx) error {
	var bookings []models.Booking
	if err := database.DB.Find(&bookings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch bookings"})
	}
	return c.JSON(bookings)
}

func UpdateBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	var booking models.Booking
	if err := database.DB.First(&booking, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Booking not found"})
	}

	type UpdateRequest struct {
		Status string `json:"status"`
	}

	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	booking.Status = req.Status
	booking.UpdatedAt = time.Now()

	if err := database.DB.Save(&booking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update booking"})
	}

	return c.JSON(fiber.Map{"message": "Booking updated successfully", "booking": booking})
}

func DeleteBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Booking{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete booking"})
	}
	return c.JSON(fiber.Map{"message": "Booking deleted successfully"})
}
