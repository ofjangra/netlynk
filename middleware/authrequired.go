package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Authrequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("netlynk_jwt")

		if cookie == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not authorized"})
		}

		token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
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
