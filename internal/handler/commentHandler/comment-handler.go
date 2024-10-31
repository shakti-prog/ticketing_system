package commenthandler

import (
	Comment "ticketing_system_backend/internal/repository/comment"
	"ticketing_system_backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetCommentHandler(c *fiber.Ctx) error {
	var commentId string = c.Params("ticketId")
	if commentId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid comment Id"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": "Comment is present"})
}

func CreateCommentHandler(c *fiber.Ctx) error {
	var commentParams utils.CommentParams
	err := c.BodyParser(&commentParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Required fields are missing for comment creation"})
	}
	err = Comment.CreateComment(commentParams.TicketId, commentParams.UserId, commentParams.Description)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"response": "comment added successfully"})

}
