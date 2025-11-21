package users

import (
	"socialmediabackend/internals/dto"
	"socialmediabackend/internals/validate"
	"socialmediabackend/services"

	"github.com/gofiber/fiber/v2"
)

// Add USER
func Add(c *fiber.Ctx) error {
	ctx := c.UserContext()
	var input dto.Usercreate

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := validate.Users(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	us := services.NewUserService()
	createdUser, err := us.CreateUser(ctx, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	createdUser.Password = "*****"
	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

// GetByID users
func GetByID(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")

	us := services.NewUserService()
	user, err := us.GetUserByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	user.Password = "*****"
	return c.JSON(user)
}

// Get All user
func GetAll(c *fiber.Ctx) error {
	ctx := c.UserContext()
	us := services.NewUserService()
	users, err := us.GetAllUsers(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	for i := range users {
		users[i].Password = "*****"
	}
	return c.JSON(users)
}

// Update USER
func Update(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")

	var input dto.Userupdate
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := validate.Users(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	us := services.NewUserService()
	updatedUser, err := us.UpdateUser(ctx, id, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if updatedUser == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	updatedUser.Password = "*****"
	return c.JSON(updatedUser)
}

// Delete USER
func Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")

	us := services.NewUserService()
	if err := us.DeleteUser(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent) // 204 No Content
}
