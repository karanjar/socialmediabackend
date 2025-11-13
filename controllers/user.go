package controllers

import (
	"socialmediabackend/internals/dto"
	_ "socialmediabackend/internals/dto"
	"socialmediabackend/internals/validate"
	"socialmediabackend/services"

	"github.com/gofiber/fiber/v2"
)

func Add(c *fiber.Ctx) error {

	ctx := c.UserContext()
	var user dto.Usercreate

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "incorrect input body",
		})
	}

	if err := validate.Users(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "incorrect input body",
		})
	}

	us := services.NewUserService()
	us.Name = user.Name
	us.Email = user.Email
	us.Password = user.Password

	us.Createuser(ctx)

	return nil

}
