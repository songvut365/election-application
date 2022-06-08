package handler

import (
	"election-service/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

var ElectionStatus bool

func httpRequest(url string, receiver interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}

	request, err := client.Get(url)
	if err != nil {
		return err
	}
	defer request.Body.Close()

	return json.NewDecoder(request.Body).Decode(receiver)
}

func Toggle(c *fiber.Ctx) error {
	// Parser
	var input model.ToggleInput
	c.BodyParser(&input)

	ElectionStatus = input.Enable

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
		"enable": ElectionStatus,
	})
}

func Count(c *fiber.Ctx) error {
	// Get all candidate
	var counts []model.Count
	httpRequest("http://localhost:5002/api/candidates", &counts)

	return c.Status(fiber.StatusOK).JSON(counts)
}

func Result(c *fiber.Ctx) error {
	// Get all candidate
	var candidates []model.ResultCandidate
	httpRequest("http://localhost:5002/api/candidates", &candidates)

	// Get all vote
	var count int
	httpRequest("http://localhost:5004/api/vote/count", &count)

	// Calculate percentage
	var resultCandidate []model.ResultCandidate

	for _, candidate := range candidates {
		var stringPercentage string

		if count == 0 {
			stringPercentage = "0%"
		} else {
			percentage := (float64(*candidate.VotedCount) / float64(count)) * 100
			stringPercentage = fmt.Sprintf("%.2f", percentage) + "%"
		}
		candidate.Percentage = &stringPercentage

		resultCandidate = append(resultCandidate, candidate)
	}

	return c.Status(fiber.StatusOK).JSON(resultCandidate)
}

func Export(c *fiber.Ctx) error {
	return c.Download(os.Getenv("CSV_FILE"), "export.csv")
}

func GetStatus(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"enable": ElectionStatus,
	})
}
