package signuphandler

import (
	signup "ticketing_system_backend/internal/repository/signUp"
	"ticketing_system_backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func SignUpHandler(c *fiber.Ctx) error {
	var credentials utils.Credentials
	err := c.BodyParser(&credentials)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email or password missing"})
	}
	err = signup.SignUp(credentials.Username, credentials.Useremail, credentials.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"response": "User created successfully"})
}
