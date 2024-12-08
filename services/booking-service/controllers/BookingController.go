package controllers

import (
	"booking-service/database"
	"booking-service/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateBooking(c *fiber.Ctx) error {
    booking := new(models.Booking)
    if err := c.BodyParser(booking); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
    }

    // Parse check_in_date and check_out_date from the request body
    checkInDate, err := time.Parse("2006-01-02", booking.CheckInDate.Format("2006-01-02"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid check_in_date format"})
    }
    checkOutDate, err := time.Parse("2006-01-02", booking.CheckOutDate.Format("2006-01-02"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid check_out_date format"})
    }

    booking.CheckInDate = checkInDate
    booking.CheckOutDate = checkOutDate
    booking.Status = "pending"
    booking.CreatedAt = time.Now()
    booking.UpdatedAt = time.Now()

    if err := database.DB.Create(&booking).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create booking"})
    }

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