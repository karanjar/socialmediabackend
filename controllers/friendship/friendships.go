package friendship

import (
	"socialmediabackend/internals/dto"
	"socialmediabackend/internals/validate"
	"socialmediabackend/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	fs := services.NewFriendshipService()
	_, err := fs.SendFriendrequest(ctx, input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
}

// GETTING ALL FRIENDS
func GetAllFriends(c *fiber.Ctx) error {
	ctx := c.UserContext()
	idStr := c.Params("userID")
	userId, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "userId is not a valid uuid"})
	}
	fs := services.NewFriendshipService()
	friend, err := fs.Getfriends(ctx, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if len(friend) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "friend not found",
			"data":  []string{},
		})
	}
	for i := range friend {
		friend[i].User.Password = ""
		friend[i].Friendship.Password = ""
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": friend})

}

//Getting By Id friends

func GetFriends(c *fiber.Ctx) error {
	ctx := c.UserContext()
	idstring := c.Params("friendID")

	iduuid, err := uuid.Parse(idstring)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid uuid format",
		})
	}

	fs := services.NewFriendshipService()
	friends, err := fs.GetfriendById(ctx, iduuid)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if friends == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no friends found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(friends)
}

func UpdateFriend(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var friend dto.FriendsUpdate
	if err := c.BodyParser(&friend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := validate.Users(friend); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	fs := services.NewFriendshipService()
	friends, err := fs.Updatefriends(ctx, friend.UserID, friend.FriendID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if friends == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no friends found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(friends)

}

func DeleteFriend(c *fiber.Ctx) error {
	ctx := c.UserContext()
	idstring := c.Params("friendID")
	iduuid, err := uuid.Parse(idstring)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid uuid format",
		})
	}

	userId := c.Query("user_id")
	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "userId is required",
		})
	}
	useriduuid, err := uuid.Parse(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid uuid format",
		})
	}

	fs := services.NewFriendshipService()
	if err := fs.DeleteFriendship(ctx, useriduuid, iduuid); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
