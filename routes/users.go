package routes

import (
	"socialmediabackend/controllers"

	"github.com/gofiber/fiber/v2"
)

func Users(r fiber.Router) {
	users := r.Group("/users")

	users.Post("/", controllers.Add)
	users.Get("/:id", controllers.GetByID)
	users.Put("/:id", controllers.Update)
	users.Delete("/:id", controllers.Delete)
}
