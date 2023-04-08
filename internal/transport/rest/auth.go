package transport_rest

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	database "github.com/ofjangra/netlynk_server/internal/database"
	"github.com/ofjangra/netlynk_server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	newUser := new(models.User)

	validate := validator.New()

	if bodyParseErr := c.BodyParser(&newUser); bodyParseErr != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "something went wrong, try again later",
		})

	}

	validationErr := validate.Struct(newUser)

	if validationErr != nil {
		fmt.Println(validationErr)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	if newUser.Username == "" || newUser.Password == "" {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "please fill the required fields",
		})

	}

	userameIsValid := regexp.MustCompile(`^[A-Za-z0-9._-]{3,16}$`).MatchString(newUser.Username)

	if len(newUser.Username) < 3 || len(newUser.Username) > 16 {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "username length must be 3 to 16 characters",
		})

	} else if !userameIsValid {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "username can only contain alphabets, numericals, -, ., and _",
		})

	}

	passwordHasWhiteSpace := regexp.MustCompile(`\s`).MatchString(newUser.Password)

	if len(newUser.Password) < 8 {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "password must be longer than 8 characters",
		})

	} else if passwordHasWhiteSpace {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "password cannot contain space",
		})

	}

	passHash, hashErr := bcrypt.GenerateFromPassword([]byte(newUser.Password), 12)

	if hashErr != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "An error occured, try again later",
		})

	}

	newUser.Password = string(passHash)

	newUser.Email = strings.ToLower(newUser.Email)

	newUser.Username = strings.ToLower(newUser.Username)

	newUser.ID = primitive.NewObjectID()

	newUser.PhotoUrl = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png"

	insertionErr := database.CreateUser(newUser)

	if insertionErr != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": insertionErr.Error(),
		})

	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "account created successfully",
	})
}

func Signin(c *fiber.Ctx) error {

	type signinReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	JWTKEY := os.Getenv("JWTKEY")

	creds := new(signinReq)

	user := new(models.User)

	parseCredsErr := c.BodyParser(&creds)

	if parseCredsErr != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "failed to signin",
		})

	}

	if creds.Username == "" || creds.Password == "" {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "please provide required credentials",
		})

	}

	creds.Username = strings.ToLower(creds.Username)

	thisUser := database.GetuserByUsername(creds.Username)

	decodeErr := thisUser.Decode(&user)

	if decodeErr != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Credentials",
		})

	}

	if user.Username == "" {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Credentials",
		})

	}

	passMatchErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))

	if passMatchErr != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Credentials",
		})

	}

	links, linksFetchErr := database.GetAllLinks(user.ID)

	if linksFetchErr != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "something went wrong, try again later",
		})

	}

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, tokenErr := token.SignedString([]byte(JWTKEY))

	if tokenErr != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "something went wrong, try again later",
		})

	}

	cookie := new(fiber.Cookie)

	cookie.Name = "auth_token"
	cookie.Value = tokenString
	cookie.HTTPOnly = true
	// cookie.SameSite = "None"
	// cookie.Secure = true
	cookie.Expires = time.Now().Add(48 * time.Hour)

	c.Cookie(cookie)

	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"success": true,
		"message": "logged in successfully",
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

func Logout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)

	cookie.Name = "auth_token"
	cookie.Value = ""
	cookie.HTTPOnly = true
	// cookie.SameSite = "None"
	// cookie.Secure = true
	cookie.Expires = time.Now().Add(-10 * time.Second)
	c.Cookie(cookie)
	return c.SendStatus(200)
}
