package controllers

import (
	"socialmediabackend/internals/dto"
	"socialmediabackend/internals/validate"

	"github.com/gofiber/fiber/v2"
)

// Create friends
func AddFriend(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var input dto.FriendsCreate

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Users(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
}
