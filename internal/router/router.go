package router

import (
	commenthandler "ticketing_system_backend/internal/handler/commentHandler"
	loginhandler "ticketing_system_backend/internal/handler/loginHandler"
	projecthandler "ticketing_system_backend/internal/handler/projectHandler"
	signuphandler "ticketing_system_backend/internal/handler/signUpHandler"
	tickethandler "ticketing_system_backend/internal/handler/ticketHandler"
	"ticketing_system_backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	app.Post("/signUp", signuphandler.SignUpHandler)

	app.Post("/login", loginhandler.LoginHandler)

	app.Get("/project", middleware.JwtMiddleWare, projecthandler.GetProjectHandler)

	app.Post("/project", middleware.JwtMiddleWare, projecthandler.CreateProjectHandler)

	app.Get("/tickets/:projectId", middleware.JwtMiddleWare, tickethandler.GetTicketsHandler)

	app.Get("/ticket/:ticketId", middleware.JwtMiddleWare, tickethandler.GetSpecificTicketHandler)

	app.Post("/ticket", middleware.JwtMiddleWare, tickethandler.CreateTicketHandler)

	app.Patch("/ticket/:ticketId", middleware.JwtMiddleWare, tickethandler.UpdateTicketHandler)

	app.Get("/comment/:ticketId", middleware.JwtMiddleWare, commenthandler.GetCommentHandler)

	app.Post("/comment", middleware.JwtMiddleWare, commenthandler.CreateCommentHandler)

}
