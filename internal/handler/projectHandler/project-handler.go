package projecthandler

import (
	"ticketing_system_backend/internal/repository/project"
	"ticketing_system_backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetProjectHandler(c *fiber.Ctx) error {
	projectNames, err := project.GetProjects()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": projectNames})
}

func CreateProjectHandler(c *fiber.Ctx) error {
	var projectParameters utils.ProjectParams
	err := c.BodyParser(&projectParameters)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Required fields are missing for project creation"})
	}
	err = project.CreateProject(projectParameters.ProjectName, projectParameters.Created_By)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": "Project created successfully"})
}
