package app

import (
	"log"
	"socialmediabackend/internals/database"
	"socialmediabackend/internals/server"
)

func Setup() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	server.Setup()
	app := server.New()

	if err := app.Listen(":3015"); err != nil {
		log.Fatalf("error starting the server: %v\n", err)
	}

}
