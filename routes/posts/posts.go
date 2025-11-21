package posts

import "github.com/gofiber/fiber/v2"

func Posts(r fiber.Router) {
	postRoute := r.Group("/posts")

	postRoute.Get("/", nil)
	postRoute.Get("/", nil)

	postRoute.Put("/:id", nil)
	postRoute.Delete("/:id", nil)

}
