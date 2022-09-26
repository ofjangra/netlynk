package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/netlynk_server/helpers"
	"github.com/ofjangra/netlynk_server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type link struct {
	title string
	url   string
}

func CreateOneLink(c *fiber.Ctx) error {
	var newLink models.Links

	if err := c.BodyParser(&newLink); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to create the link"})
	}
	createrId, idErr := primitive.ObjectIDFromHex(c.Locals("user_id").(string))

	if idErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to create the link"})
	}
	newLink.Created_by = createrId
	insertionErr := helpers.CreateALink(newLink)

	if insertionErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create the link"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Link Created Successfully"})
}

func UpdateOneLink(c *fiber.Ctx) error {

	linkId := c.Params("id")

	if linkId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to edit the link"})
	}

	linkBody := bson.M{}

	if err := c.BodyParser(&linkBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to edit the link"})
	}

	updateErr := helpers.EditALink(linkId, linkBody)

	if updateErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to edit the link"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Link updated"})
}

func DeleteOneLink(c *fiber.Ctx) error {
	linkId := c.Params("id")

	deleteErr := helpers.DeleteALink(linkId)

	if deleteErr != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete link"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Link Deleted Successfully"})
}

func GetALink(c *fiber.Ctx) error {
	linkId := c.Params("id")

	link := new(models.Links)

	result, respErr := helpers.GetALink(linkId)

	if respErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Link not found"})
	}

	result.Decode(&link)

	return c.Status(fiber.StatusFound).JSON(link)
}

func GetAllLinks(c *fiber.Ctx) error {
	createrId := c.Params("id")

	links, err := helpers.GetAllLinks(createrId)

	if err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Failed to fetch links"})
	}

	return c.Status(fiber.StatusFound).JSON(links)
}
