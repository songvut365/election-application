package handler

import (
	"election-service/model"
	"os"

	"github.com/gofiber/fiber/v2"
)

var ElectionStatus bool

func Toggle(c *fiber.Ctx) error {
	var input model.ToggleInput
	c.BodyParser(&input)

	ElectionStatus = input.Enable

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"enable": ElectionStatus,
	})
}

func Count(c *fiber.Ctx) error {
	return c.SendString("Election Count")
}

func Result(c *fiber.Ctx) error {
	return c.SendString("Election Result")
}

func Export(c *fiber.Ctx) error {
	return c.Download(os.Getenv("CSV_FILE"), "export.csv")
}

func GetStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(ElectionStatus)
}
