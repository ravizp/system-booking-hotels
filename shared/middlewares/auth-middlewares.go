package middlewares

import (
	"strings"

	"services/user-service/models"

	"shared/jwt"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token format"})
	}

	tokenString := tokenParts[1]
	parsedToken, err := jwt.ValidateToken(tokenString)
	if err != nil {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid token"})
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(403).JSON(fiber.Map{"error": "Invalid user ID in token"})
	}

	var user models.User
	if err := models.DB.First(&user, uint(userID)).Error; err != nil {
		return c.Status(403).JSON(fiber.Map{"error": "User not found"})
	}

	c.Locals("role", user.Role)
	c.Locals("userID", user.ID)
	return c.Next()
}

func AdminOnly(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "Admin access required"})
	}
	return c.Next()
}
