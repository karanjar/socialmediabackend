package posts

import (
	"socialmediabackend/controllers/posts"

	"github.com/gofiber/fiber/v2"
)

func Posts(r fiber.Router) {
	postRoute := r.Group("/posts")

	postRoute.Post("/", posts.AddPosts)

	postRoute.Get("/", posts.GetAllPosts)
	postRoute.Get("/:id", posts.GetPostById)

	postRoute.Put("/:id", posts.UpdatePostById)
	postRoute.Delete("/:id", posts.DeletePostById)

}
