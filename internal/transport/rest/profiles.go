package transport_rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	database "github.com/ofjangra/netlynk_server/internal/database"
	"github.com/ofjangra/netlynk_server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Profile(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)

	result, resErr := database.GetUserById(userId)

	if resErr != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": resErr.Error(),
		})
	}

	user := new(models.User)

	err := result.Decode(&user)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "user not found",
		})
	}

	links, linksFetchErr := database.GetAllLinks(user.ID)

	if linksFetchErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "something went wrong, try again later",
		})
	}

	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"success": true,
		"profile": fiber.Map{
			"id":        user.ID,
			"username":  user.Username,
			"name":      user.Name,
			"photo_url": user.PhotoUrl,
			"bio":       user.Bio,
			"email":     user.Email,
			"links":     links,
		},
	})
}

func GetUser(c *fiber.Ctx) error {

	username := c.Params("username")

	user := new(models.User)

	result := database.GetuserByUsername(username)

	err := result.Decode(&user)

	if err != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}
	links, linksFetchErr := database.GetAllLinks(user.ID)

	if linksFetchErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "something went wrong",
		})
	}

	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"success": true,
		"profile": fiber.Map{
			"id":        user.ID,
			"photo_url": user.PhotoUrl,
			"username":  user.Username,
			"name":      user.Name,
			"bio":       user.Bio,
			"links":     links,
		},
	})
}

func EditProfile(c *fiber.Ctx) error {

	userId := c.Locals("user_id").(string)

	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
		})
	}

	body := bson.M{}

	if err := c.BodyParser(&body); err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong struct",
		})
	}

	fmt.Println(" body : ", body)

	updateErr := database.EditProfile(userId, body)

	if updateErr != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"success": false,
			"message": updateErr.Error()})
	}

	newProfileData, profileErr := database.GetUserById(userId)

	if profileErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": true,
			"message": "something went wrong",
		})
	}

	user := new(models.User)

	decodeErr := newProfileData.Decode(&user)

	if decodeErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": true,
			"message": "something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "profile updated",
		"profile": fiber.Map{
			"photo_url": user.PhotoUrl,
			"username":  user.Username,
			"name":      user.Name,
			"bio":       user.Bio,
			"email":     user.Email,
		},
	})

}
