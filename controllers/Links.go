package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/netlynk_server/helpers"
	"github.com/ofjangra/netlynk_server/models"
)

func CreateOneLink(c *fiber.Ctx) error {
	var newLink models.Links

	if err := c.BodyParser(&newLink); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to create the link"})
	}

	insertionErr := helpers.CreateALink(newLink)

	if insertionErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create the link"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Link Created Successfully"})
}
