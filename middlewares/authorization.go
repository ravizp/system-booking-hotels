package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Authorization(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil role dari context
		role := c.Locals("role")

		// Periksa role
		if role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "You do not have permission to access this resource",
			})
		}

		return c.Next()
	}
}
