package handler

import (
	"candidate-service/config"
	"candidate-service/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllCandidate(c *fiber.Ctx) error {
	db := config.DB

	var candidates = []model.Candidate{}
	db.Model(&model.Candidate{}).Find(&candidates)

	return c.Status(fiber.StatusOK).JSON(candidates)
}

func CreateCandidate(c *fiber.Ctx) error {
	db := config.DB

	var candidate model.Candidate
	c.BodyParser(&candidate)

	db.Create(&candidate)

	return c.Status(fiber.StatusOK).JSON(candidate)
}

func GetCandidate(c *fiber.Ctx) error {
	db := config.DB

	id := c.Params("id")
	var candidate model.Candidate

	err := db.Model(&model.Candidate{}).Where("id = ?", id).First(&candidate).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "candidate not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(candidate)
}

func UpdateCandidate(c *fiber.Ctx) error {
	db := config.DB

	var updateCandidate model.Candidate
	c.BodyParser(&updateCandidate)

	id := c.Params("id")

	db.Model(&model.Candidate{}).Where("id = ?", id).Updates(&updateCandidate)

	var candidate model.Candidate
	db.Model(&model.Candidate{}).Where("id = ?", id).First(&candidate)

	return c.Status(fiber.StatusOK).JSON(candidate)
}

func DeleteCandidate(c *fiber.Ctx) error {
	db := config.DB

	id := c.Params("id")
	db.Delete(&model.Candidate{}, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "candidate deleted",
	})
}
