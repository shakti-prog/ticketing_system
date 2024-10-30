package loginhandler

import (
	"ticketing_system_backend/internal/repository/login"
	"ticketing_system_backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	var credentials utils.Credentials
	err := c.BodyParser(&credentials)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	loginError := login.Login(credentials.Useremail, credentials.Password)
	if loginError != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Incorrect email or password"})
	}
	token, err := GenerateJWTToken(credentials.Useremail)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error while generating token pls try again"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})

}
