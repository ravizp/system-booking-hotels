package controllers

import (
	"net/http"
	"user-service/models"

	"github.com/gofiber/fiber/v2"
	"github.com/ravizp/system-booking-hotels/shared/jwt"
	"github.com/ravizp/system-booking-hotels/shared/utils"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	//hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot hash password"})
	}
	user.Password = hashedPassword

	//save ke database
	if err := models.DB.Create(user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
	data := new(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	})
	if err := c.BodyParser(data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var user models.User
	if err := models.DB.Where("email = ?", data.Email).First(&user).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if err := utils.ComparePassword(user.Password, data.Password); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := jwt.GenerateToken(user.ID, user.Name, user.Role)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"message": "Login successful", "token": token})
}

func GetAllUsers(c *fiber.Ctx) error {
	// Verify JWT token
	userID, err := jwt.ValidateToken(c.Get("Authorization"))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.JSON(fiber.Map{"data": users, "user_id": userID})
}

func GetUserByID(c *fiber.Ctx) error {
	// Verify JWT token
	userID, err := jwt.ValidateToken(c.Get("Authorization"))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")

	var user models.User
	if err := models.DB.First(&user, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{"data": user, "user_id": userID})
}

func Logout(c *fiber.Ctx) error {
	// Verify JWT token
	_, err := jwt.ValidateToken(c.Get("Authorization"))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Dummy implementation
	// Usually, token management is handled on the frontend or with a blacklist
	return c.JSON(fiber.Map{"message": "Logout successful"})
}
