package middleware

import (
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type ReqHeader struct {
	Authorization string
}

func Authrequired() fiber.Handler {
	return func(c *fiber.Ctx) error {

		if c.Cookies("auth_token") == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not authorized"})
		}

		// fmt.Println("cookie:", c.Cookies("auth_token"))

		token, err := jwt.Parse(c.Cookies("auth_token"), func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTKEY")), nil
		})
		// fmt.Println(token)
		if err != nil {
			// fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not verify token"})
		}

		payload := token.Claims.(jwt.MapClaims)

		fmt.Println("payload  :  ", payload["id"])
		// c.Locals("employee_id", payload["id"])
		c.Locals("user_id", payload["id"])
		fmt.Println(c.Locals("user_id"))
		return c.Next()
		// auth := ReqHeader{}

		// if err := c.ReqHeaderParser(&auth); err != nil {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Something Went Wrong"})
		// }

		// if auth.Authorization == "" {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not authorized"})
		// }

		// authToken := strings.Split(auth.Authorization, "Bearer ")

		// token, err := jwt.Parse(authToken[1], func(token *jwt.Token) (interface{}, error) {
		// 	return []byte(os.Getenv("JWTKEY")), nil
		// })

		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
		// }

		// payload := token.Claims.(jwt.MapClaims)

		// c.Locals("user_id", payload["id"])

		// return c.Next()
	}
}
