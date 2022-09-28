package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserProps struct {
	ViewerId string `json:"viewer_id"`
	LoggedIn bool   `json:"logged_in"`
}

func Validateuser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user_props := new(UserProps)
		cookie := c.Cookies("netlynk_jwt")

		if cookie == "" {
			user_props.ViewerId = ""
			user_props.LoggedIn = false

			c.Locals("user_props", user_props)

			return c.Next()
		}

		token, tokenErr := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		})

		if tokenErr != nil {
			user_props.ViewerId = ""
			user_props.LoggedIn = false

			c.Locals("user_props", user_props)

			return c.Next()
		}

		user_id := token.Claims.(jwt.MapClaims)["id"]

		user_props.ViewerId = user_id.(string)
		user_props.LoggedIn = true

		c.Locals("user_props", user_props)

		return c.Next()
	}
}
