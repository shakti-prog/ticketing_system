package main

import (
	"log"
	"ticketing_system_backend/internal/config"
	"ticketing_system_backend/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.LoadEnv()
	router.Router(app)
	log.Fatal(app.Listen(":9001"))
}
