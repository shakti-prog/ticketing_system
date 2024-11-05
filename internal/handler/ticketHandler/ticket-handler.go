package tickethandler

import (
	"strconv"
	"ticketing_system_backend/internal/repository/ticket"
	"ticketing_system_backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetTicketsHandler(c *fiber.Ctx) error {
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
	var ticket_no, err = strconv.Atoi(c.Params("ticketId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ticket Id"})
	}
	var udpateTicketParams utils.UpdateTicketParams
	err = c.BodyParser(&udpateTicketParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Body params"})
	}
	err = ticket.UpdateTicket(udpateTicketParams.Field, udpateTicketParams.Value, int64(ticket_no))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": "Ticket Updated successfully"})
}

func GetSpecificTicketHandler(c *fiber.Ctx) error {
	var ticketId, err = strconv.Atoi(c.Params("ticketId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ticket Id"})
	}
	ticketDetails, err := ticket.GetSpecificTicket((int64(ticketId)))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": ticketDetails})
}
