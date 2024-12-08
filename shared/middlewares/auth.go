package middlewares

import (
	"shared/jwt"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized, token required",
		})
	}

	claims, err := jwt.ParseToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Simpan informasi user ke context
	c.Locals("userId", claims.UserID)
	c.Locals("name", claims.Name)
	c.Locals("role", claims.Role)

	return c.Next()
}

func Authorization(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		if role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden, insufficient permissions",
			})
		}
		return c.Next()
	}
}