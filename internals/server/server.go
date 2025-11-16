package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: ErrorHandler,
		BodyLimit:    16 * 1024 * 1024,
	})

	// Middleware
	Middleware(app)

	// Recovery and 404
	app.Use(recover.New())
	//app.Use(NotFoundHandler)

	return app
}
