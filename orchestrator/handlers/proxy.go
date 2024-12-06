package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ProxyUserService(c *fiber.Ctx) error {
	return c.Redirect("http://user-service:8081"+c.OriginalURL(), fiber.StatusTemporaryRedirect)
}

func ProxyHotelService(c *fiber.Ctx) error {
	return c.Redirect("http://hotel-service:8082"+c.OriginalURL(), fiber.StatusTemporaryRedirect)
}

func ProxyBookingService(c *fiber.Ctx) error {
	return c.Redirect("http://booking-service:8083"+c.OriginalURL(), fiber.StatusTemporaryRedirect)
}