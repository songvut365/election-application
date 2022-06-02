package main

import (
	"os"
	"vote-service/config"
	"vote-service/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Setup
	config.SetupEnv()
	config.SetupDatabase()

	// Router
	vote := app.Group("/api/vote")
	vote.Post("/status", handler.CheckVoteStatus)
	vote.Post("/", handler.Vote)

	app.Listen(":" + os.Getenv("PORT"))
}
