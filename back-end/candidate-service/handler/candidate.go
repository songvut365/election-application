package handler

import "github.com/gofiber/fiber/v2"

func GetAllCandidate(c *fiber.Ctx) error {
	return c.SendString("Get all candidate")
}

func CreateCandidate(c *fiber.Ctx) error {
	return c.SendString("Create Candidate")
}

func GetCandidate(c *fiber.Ctx) error {
	return c.SendString("Get Candidate")
}

func UpdateCandidate(c *fiber.Ctx) error {
	return c.SendString("Update Candidate")
}

func DeleteCandidate(c *fiber.Ctx) error {
	return c.SendString("Delete Candidate")
}
