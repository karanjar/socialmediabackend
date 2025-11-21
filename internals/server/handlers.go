package server

import (
	"socialmediabackend/routes/friendship"
	"socialmediabackend/routes/users"

	"github.com/gofiber/fiber/v2"
)

//Default error handler

func ErrorHandler(c *fiber.Ctx, err error) error {
	msg := err.Error()
	return c.Status(fiber.StatusInternalServerError).JSON(msg)
}

// Not found handler
//var NotFoundHandler = func(ctx *fiber.Ctx) error {
//	return ctx.Status(fiber.StatusNotFound).JSON("requested resource not found")
//}

func Addroutes(app *fiber.App) {
	baseRoute := app.Group("/socio")
	users.Users(baseRoute)
	friendship.Friendships(baseRoute)
}
