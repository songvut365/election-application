package main

import (
	"os"
	"vote-service/config"
	"vote-service/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup
	config.SetupEnv()
	config.SetupDatabase()

	// Router
	vote := app.Group("/api/vote")
	vote.Post("/status", handler.CheckVoteStatus)
	vote.Post("/", handler.Vote)

	vote.Get("/count", handler.GetAllVoteCount)

	app.Listen(":" + os.Getenv("PORT"))
}
