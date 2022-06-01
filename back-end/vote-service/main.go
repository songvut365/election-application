package main

import (
	"os"
	"vote-service/config"
	"vote-service/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Setup
	config.SetupEnv()
	config.SetupDatabase()

	// Router
	vote := app.Group("/api/vote")
	vote.Post("/status", handler.CheckVoteStatus)
	vote.Post("/", handler.Vote)

	app.Listen(":" + os.Getenv("PORT"))
}
