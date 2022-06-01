package main

import (
	"candidate-service/config"
	"candidate-service/handler"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

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

	app.Listen(":" + os.Getenv("PORT"))
}
