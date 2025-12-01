package notifications

import (
	"bufio"
	"fmt"
	"socialmediabackend/internals/notifications"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func StreamNotifications(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
	}

	// 1. Register the user to ensure their channel exists
	notifications.Register(userID)

	// 2. Set headers for Server-Sent Events (SSE)
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	// 3. Start the stream writer
	c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		// Get the user's channel
		ch, ok := notifications.GetUserChannel(userID)
		if !ok {
			return
		}

		// Notify that connection is open
		fmt.Fprintf(w, "data: Connected to notification stream\n\n")
		w.Flush()

		// Listen for new messages
		for {
			select {
			case msg := <-ch:
				// Write the message in SSE format: "data: <message>\n\n"
				fmt.Fprintf(w, "data: %s\n\n", msg)
				w.Flush()
			}
		}
	})

	return nil
}
