package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type ReqHeader struct {
	Authorization string
}

func Authrequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := ReqHeader{}

		if err := c.ReqHeaderParser(&auth); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Something Went Wrong"})
		}

		if auth.Authorization == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not authorized"})
		}

		authToken := strings.Split(auth.Authorization, "Bearer ")

		token, err := jwt.Parse(authToken[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
		}

		payload := token.Claims.(jwt.MapClaims)

		c.Locals("user_id", payload["id"])

		return c.Next()
	}
}
