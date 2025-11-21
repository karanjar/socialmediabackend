package users

import (
	users2 "socialmediabackend/controllers/users"

	"github.com/gofiber/fiber/v2"
)

func Users(r fiber.Router) {
	users := r.Group("/users")

	users.Post("/", users2.Add)
	users.Get("/", users2.GetAll)
	users.Get("/:id", users2.GetByID)
	users.Put("/:id", users2.Update)
	users.Delete("/:id", users2.Delete)
}
