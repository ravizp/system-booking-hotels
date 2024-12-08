package controllers

import (
	"booking-service/database"
	"booking-service/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func CreatePayment(c *fiber.Ctx) error {
	bookingID := c.Params("bookingId")

	// Cari booking di database
	var booking models.Booking
	if err := database.DB.First(&booking, "id = ?", bookingID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Booking not found"})
	}

	// Parse request body
	data := struct {
		Amount  int    `json:"amount"`
		Method  string `json:"method"`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Konfigurasi Stripe
	stripe.Key = "your-stripe-secret-key"

	// Buat PaymentIntent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(data.Amount)),
		Currency: stripe.String("idr"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}
	intent, err := paymentintent.New(params)
	if err != nil {
		log.Println("Stripe Error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create payment"})
	}

	// Simpan payment ke database
	payment := models.Payment{
		BookingID:    booking.ID,
		Amount:       data.Amount,
		PaymentMethod: data.Method,
		Status:       "pending",
		PaymentDate:  time.Now(),
	}
	if err := database.DB.Create(&payment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save payment"})
	}

	// Response dengan client secret untuk Stripe
	return c.JSON(fiber.Map{
		"message":       "Payment initiated",
		"client_secret": intent.ClientSecret,
	})
}
