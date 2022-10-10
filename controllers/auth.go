package controllers

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ofjangra/netlynk_server/helpers"
	"github.com/ofjangra/netlynk_server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SigninReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var JWTKEY string

func init() {
	// EnvErr := godotenv.Load(".env")

	// if EnvErr != nil {
	// 	log.Fatal(EnvErr)
	// }

	JWTKEY = os.Getenv("JWTKEY")
}

func Signup(c *fiber.Ctx) error {

	newUser := new(models.User)

	err := c.BodyParser(&newUser)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to create account"})
	}

	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "please provide required credentials"})
	} else if len(newUser.Username) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "username too short"})
	} else if len(newUser.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "password too short"})
	}

	passHash, hashErr := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)

	if hashErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "An error occured, try again later"})
	}

	newUser.Password = string(passHash)
	newUser.Email = strings.ToLower(newUser.Email)
	newUser.Username = strings.ToLower(newUser.Username)

	newUser.ID = primitive.NewObjectID()

	newUser.PhotoUrl = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png"

	insertionErr := helpers.CreateUser(newUser)

	if insertionErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": insertionErr.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Account created successfully"})

}

func Signin(c *fiber.Ctx) error {

	creds := new(SigninReq)

	user := new(models.User)

	parseCredsErr := c.BodyParser(&creds)

	if parseCredsErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to signin"})
	}

	if creds.Username == "" || creds.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Please provide required credentials"})
	}

	creds.Username = strings.ToLower(creds.Username)

	thisUser := helpers.GetuserByUsername(creds.Username)

	thisUser.Decode(&user)

	if user.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Credentials"})
	}

	passMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))

	if passMatchErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Credentials"})
	}

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 168).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, tokenErr := token.SignedString([]byte(JWTKEY))

	if tokenErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to login"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user, "token": tokenString})

}

func Logout(c *fiber.Ctx) error {
	c.ClearCookie("netlynk_jwt")
	return c.SendStatus(200)
}
