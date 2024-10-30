package tickethandler

import (
	"fmt"
	"ticketing_system_backend/internal/repository/ticket"
	"ticketing_system_backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetTicketHandler(c *fiber.Ctx) error {
	var ticketId string = c.Params("projectId")
	if ticketId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project Id"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": "Ticket is present"})
}

func CreateTicketHandler(c *fiber.Ctx) error {
	var ticketParams utils.TicketParams
	err := c.BodyParser(&ticketParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Required fields are missing for creating a ticket"})
	}
	err = ticket.CreateTicket(ticketParams.AssigneeId, ticketParams.ReporterId, ticketParams.ProjectId, ticketParams.Description, ticketParams.Status, ticketParams.Priority)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": "Ticket created successfully"})
}

func UpdateTicketHandler(c *fiber.Ctx) error {
	var ticketId string = c.Params("ticketId")
	var field string = c.Params("field")
	var value string = c.Params("value")
	if ticketId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ticket Id"})
	}
	fmt.Println(field)
	fmt.Println(value)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": "Ticket is present"})
}
