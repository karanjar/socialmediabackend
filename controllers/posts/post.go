package posts

import (
	"socialmediabackend/internals/dto"
	"socialmediabackend/internals/validate"
	"socialmediabackend/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// create friends
func AddPosts(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var input dto.CreatePost
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validate.Users(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	ps := services.NewPostsService()
	newpost, err := ps.CreatePost(ctx, input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newpost)
}

// Get all Posts
func GetAllPosts(c *fiber.Ctx) error {
	ctx := c.UserContext()

	ps := services.NewPostsService()
	allposts, err := ps.GetAllPosts(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(allposts)
}

// Get Posts by Id
func GetPostById(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")
	uuID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	ps := services.NewPostsService()
	post, err := ps.GetPostByID(ctx, uuID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if post == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "post not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

// Update Posts
func UpdatePostById(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")

	uuID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var input dto.CreatePost
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := validate.Users(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	ps := services.NewPostsService()
	post, err := ps.UpdatePost(ctx, uuID, input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if post == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "post not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(post)
}

// Delete posts
func DeletePostById(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")

	ps := services.NewPostsService()
	if err := ps.DeletePost(ctx, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
