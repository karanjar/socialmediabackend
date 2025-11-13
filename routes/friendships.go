package routes

import "github.com/gofiber/fiber/v2"

func Friendships(r fiber.Router) {
	postRoute := r.Group("/posts")

	postRoute.Post("/", nil)
	postRoute.Get("/", nil)
	postRoute.Delete("/", nil)
}
