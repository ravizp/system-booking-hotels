package routes

import (
	"booking-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/bookings")
	api.Post("/", controllers.CreateBooking)
	api.Get("/", controllers.GetBookings)
	api.Patch("/:id", controllers.UpdateBooking)
	api.Delete("/:id", controllers.DeleteBooking)

	api.Post("/payments/:bookingId", controllers.CreatePayment)
	api.Get("/payments", controllers.GetAllPayments)
	api.Post("/refunds", controllers.RefundBooking)
	api.Get("/refunds", controllers.GetAllRefunds)
	api.Patch("/refunds/:id", controllers.UpdateRefundStatus)
}
