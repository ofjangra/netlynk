package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/netlynk_server/helpers"
	"github.com/ofjangra/netlynk_server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Profile(c *fiber.Ctx) error {

	userId := c.Locals("user_id").(string)

	result, resErr := helpers.GetUserById(userId)

	if resErr != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	user := new(models.User)

	err := result.Decode(&user)

	if err != nil {
		c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"id":        user.ID,
		"username":  user.Username,
		"photo_url": user.PhotoUrl,
		"email":     user.Email,
		"bio":       user.Bio,
	})
}

func GetUser(c *fiber.Ctx) error {

	username := c.Params("username")

	user := new(models.User)

	result := helpers.GetuserByUsername(username)

	err := result.Decode(&user)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"_id":       user.ID,
		"photo_url": user.PhotoUrl,
		"username":  user.Username,
		"bio":       user.Bio,
	})
}

func EditProfile(c *fiber.Ctx) error {

	userId := c.Locals("user_id").(string)

	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Something went wrong"})
	}

	update := new(models.User)

	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Something went wrong"})
	}

	updateErr := helpers.EditProfile(userId, update)

	if updateErr != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": updateErr.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "profile updated"})

}

func EditProfilePhoto(c *fiber.Ctx) error {

	userId := c.Locals("user_id").(string)

	update := bson.M{}

	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to parse body"})
	}
	if userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Something went wrong"})
	}

	updateErr := helpers.EditProfilePhoto(userId, update)

	if updateErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": updateErr.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "photo updated"})
}
