package main

import (
	"candidate-service/config"
	"candidate-service/handler"
	"candidate-service/rabbitmq"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Setup
	config.SetupEnv()
	config.SetupDatabase()

	// Router
	candidate := app.Group("/api/candidates")
	candidate.Get("/", handler.GetAllCandidate)
	candidate.Post("/", handler.CreateCandidate)
	candidate.Get("/:id", handler.GetCandidate)
	candidate.Put("/:id", handler.UpdateCandidate)
	candidate.Delete("/:id", handler.DeleteCandidate)

	// RabbitMQ
	go rabbitmq.ReceiveVote()

	app.Listen(":" + os.Getenv("PORT"))
}
