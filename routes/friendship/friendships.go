package friendship

import (
	"socialmediabackend/controllers/friendship"

	"github.com/gofiber/fiber/v2"
)

func Friendships(r fiber.Router) {
	postRoute := r.Group("/friends")

	postRoute.Post("/", friendship.AddFriend)
	postRoute.Get("/:friendID", friendship.GetFriends)
	postRoute.Put("/", friendship.UpdateFriend)
	postRoute.Delete("/:friendID", friendship.DeleteFriend)
}
