package app

import (
	"log"
	"socialmediabackend/internals/database"
	"socialmediabackend/internals/server"
)

func Setup() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a new Fiber app
	app := server.New() // returns *fiber.App

	// Register /users routes
	server.Addroutes(app)

	// Start the server
	if err := app.Listen(":3015"); err != nil {
		log.Fatalf("error starting the server: %v\n", err)
	}
}
