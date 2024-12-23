package middlewares

import (
	"log"

	"github.com/ravizp/system-booking-hotels/shared/jwt"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	log.Printf("Authorization Header: %s", authHeader) // Log header untuk verifikasi

	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format",
		})
	}

	tokenString := authHeader[7:]
	claims, err := jwt.GetClaims(tokenString)
	if err != nil {
		log.Printf("Token validation error: %v", err) // Log error jika parsing token gagal
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Simpan klaim ke context
	c.Locals("userId", claims["userId"])
	c.Locals("name", claims["name"])
	c.Locals("role", claims["role"])

	return c.Next()
}
