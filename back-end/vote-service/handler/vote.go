package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"vote-service/config"
	"vote-service/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func httpRequest(url string, receiver interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}

	request, err := client.Get(url)
	if err != nil {
		return err
	}
	defer request.Body.Close()

	return json.NewDecoder(request.Body).Decode(receiver)
}

func CheckVoteStatus(c *fiber.Ctx) error {
	db := config.MI.DB.Collection(os.Getenv("MONGODB_COLLECTION"))

	// Parser Input
	var input model.Vote
	c.BodyParser(&input)

	// Find by national id
	vote := &model.Vote{}

	query := bson.M{"nationalid": input.NationalID}

	err := db.FindOne(c.Context(), query).Decode(vote)
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": true,
	})
}

func Vote(c *fiber.Ctx) error {
	db := config.MI.DB.Collection(os.Getenv("MONGODB_COLLECTION"))

	// Parser Input
	var input model.Vote
	c.BodyParser(&input)

	// Check election is closed (API)
	var status model.ElectionStatus
	httpRequest("http://localhost:5003/api/election/status", &status)

	if !status.Enable {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "error",
			"message": "Election is closed",
		})
	}

	// Find by national id
	vote := &model.Vote{}

	query := bson.M{"nationalid": input.NationalID}

	err := db.FindOne(c.Context(), query).Decode(vote)
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "error",
			"message": "Already voted",
		})
	}

	// Voting
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	db.InsertOne(c.Context(), input)

	// Update candidate voted count (RabbitMQ)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}
