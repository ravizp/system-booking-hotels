package handlers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func ProxyUserService(c *fiber.Ctx) error {
    userServiceURL := os.Getenv("USER_SERVICE_URL")
    if userServiceURL == "" {
        userServiceURL = "http://user-service:8081"
    }

    targetURL := userServiceURL + c.OriginalURL()
    log.Printf("Redirecting to: %s", targetURL)

    // Menggunakan middleware proxy untuk meneruskan semua permintaan
    return proxy.Do(c, targetURL)
}


// func ProxyHotelService(c *fiber.Ctx) error {
// 	return c.Redirect("http://hotel-service:8082"+c.OriginalURL(), fiber.StatusTemporaryRedirect)
// }

// func ProxyBookingService(c *fiber.Ctx) error {
// 	return c.Redirect("http://booking-service:8083"+c.OriginalURL(), fiber.StatusTemporaryRedirect)
// }