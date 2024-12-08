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

	return proxy.Do(c, targetURL)
}

func ProxyHotelService(c *fiber.Ctx) error {
	hotelServiceURL := os.Getenv("HOTEL_SERVICE_URL")
	if hotelServiceURL == "" {
		hotelServiceURL = "http://hotel-service:8082"
	}

	targetURL := hotelServiceURL + c.OriginalURL()
	log.Printf("Redirecting to: %s", targetURL)

	return proxy.Do(c, targetURL)
}

func ProxyBookingService(c *fiber.Ctx) error {
	bookingServiceURL := os.Getenv("BOOKING_SERVICE_URL")
	if bookingServiceURL == "" {
		bookingServiceURL = "http://booking-service:8083"
	}

	targetURL := bookingServiceURL + c.OriginalURL()
	log.Printf("Redirecting to: %s", targetURL)

	return proxy.Do(c, targetURL)
}

func ProxyNotificationService(c *fiber.Ctx) error {
	notificationServiceURL := os.Getenv("NOTIFICATION_SERVICE_URL")
	if notificationServiceURL == "" {
		notificationServiceURL = "http://notification-service:8084"
	}

	targetURL := notificationServiceURL + c.OriginalURL()
	log.Printf("Redirecting to: %s", targetURL)

	return proxy.Do(c, targetURL)
}
