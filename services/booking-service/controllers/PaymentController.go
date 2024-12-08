package controllers

import (
	"booking-service/database"
	"booking-service/models"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func CreatePayment(c *fiber.Ctx) error {
	// Ambil bookingId dari parameter
	bookingID := c.Params("bookingId")

	// Cari booking di database
	var booking models.Booking
	if err := database.DB.First(&booking, "id = ?", bookingID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Booking not found"})
	}

	data := struct {
		Amount int    `json:"amount"` // Jumlah dalam IDR (rupiah penuh)
		Method string `json:"method"` // Metode pembayaran (e.g., card)
	}{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input format"})
	}

	// Validasi input
	if data.Amount <= 0 || data.Method == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Amount and method are required"})
	}

	// Konfigurasi Stripe
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	if stripe.Key == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Stripe API key not configured"})
	}

	// Konversi cent ke rupiah
	amountInCents := data.Amount * 100

	// Buat PaymentIntent
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(amountInCents)),
		Currency: stripe.String("idr"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}
	intent, err := paymentintent.New(params)
	if err != nil {
		log.Printf("Stripe Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create payment intent"})
	}

	// Simpan pembayaran ke database
	payment := models.Payment{
		BookingID:     booking.ID,
		Amount:        data.Amount,
		PaymentMethod: data.Method,
		Status:        "pending",
		PaymentDate:   time.Now(),
	}
	if err := database.DB.Create(&payment).Error; err != nil {
		log.Printf("Database Error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save payment"})
	}

	// Kirim client secret ke frontend untuk menyelesaikan pembayaran
	return c.JSON(fiber.Map{
		"message":       "Payment initiated",
		"client_secret": intent.ClientSecret,
	})
}

// GetAllPayments - Mengambil semua data pembayaran
func GetAllPayments(c *fiber.Ctx) error {
	var payments []models.Payment
	if err := database.DB.Find(&payments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch payments"})
	}

	return c.JSON(payments)
}
