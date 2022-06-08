package main

import (
	"election-service/config"
	"election-service/handler"
	"election-service/rabbitmq"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
)

// Instance
var ElectionStatus = true

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Setup
	config.SetupEnv()

	// Inject instance
	handler.ElectionStatus = ElectionStatus

	// Router
	election := app.Group("/api/election")
	election.Post("/toggle", handler.Toggle)
	election.Post("/count", handler.Count)
	election.Post("/result", handler.Result)
	election.Post("/export", handler.Export)

	election.Get("/status", handler.GetStatus)

	// Web Socket
	ws := app.Group("/ws/candidates", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	ws.Get("/", websocket.New(handler.AllCandidateVoteStream))
	ws.Get("/:id", websocket.New(handler.CandidateVoteStream))

	// RabbitMQ
	go rabbitmq.ReceiveVote()

	app.Listen(":" + os.Getenv("PORT"))
}
