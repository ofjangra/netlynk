package transport_rest

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	database "github.com/ofjangra/netlynk_server/internal/database"
	"github.com/ofjangra/netlynk_server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOneLink(c *fiber.Ctx) error {

	var newLink models.Links

	if err := c.BodyParser(&newLink); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "failed to create the link",
		})
	}
	creatorId, idErr := primitive.ObjectIDFromHex(c.Locals("user_id").(string))

	if idErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "failed to create the link",
		})
	}
	newLink.ID = primitive.NewObjectID()
	newLink.Created_by = creatorId
	newLink.CreatedOn = primitive.NewDateTimeFromTime(time.Now())
	insertionErr := database.CreateALink(newLink)

	if insertionErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": insertionErr.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Link Created Successfully",
		"_id":     newLink.ID,
	})
}

func UpdateOneLink(c *fiber.Ctx) error {
	fmt.Println("Update link")
	linkId := c.Params("id")

	if linkId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to edit the link",
		})
	}

	linkBody := bson.M{}

	if err := c.BodyParser(&linkBody); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update the link",
		})
	}

	thisLink, linkFindErr := database.GetALink(linkId)

	if linkFindErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update the link",
		})
	}
	link := new(models.Links)

	decodeErr := thisLink.Decode(&link)

	if decodeErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update the link",
		})
	}

	if link.Created_by.Hex() == c.Locals("user_id").(string) {

		linkBody["updated_on"] = primitive.NewDateTimeFromTime(time.Now())
		updateErr := database.EditALink(linkId, linkBody)

		if updateErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   updateErr.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"success": true,
			"message": "Link updated",
		})

	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": "failed to update link",
	})

}

func DeleteOneLink(c *fiber.Ctx) error {
	linkId := c.Params("id")

	thisLink, linkFindErr := database.GetALink(linkId)

	if linkFindErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete the link",
		})
	}
	link := new(models.Links)

	decodeErr := thisLink.Decode(&link)

	if decodeErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete the link",
		})
	}

	if link.Created_by.Hex() == c.Locals("user_id").(string) {

		deleteErr := database.DeleteALink(linkId)

		if deleteErr != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"message": deleteErr.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Link Deleted Successfully",
		})

	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"message": "Failed to delete the link",
	})

}

func GetALink(c *fiber.Ctx) error {

	linkId := c.Params("id")

	link := new(models.Links)

	result, respErr := database.GetALink(linkId)

	if respErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": respErr.Error(),
		})
	}

	decodeErr := result.Decode(&link)

	if decodeErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "failed to fetch link",
		})
	}

	return c.Status(fiber.StatusFound).JSON(link)
}

func GetAllLinks(c *fiber.Ctx) error {

	createrId := c.Params("id")

	hexId, idErr := primitive.ObjectIDFromHex(createrId)

	if idErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "failed to fetch link",
		})
	}

	links, err := database.GetAllLinks(hexId)

	if err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusFound).JSON(links)
}
