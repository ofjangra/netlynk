package routes

import (
	"github.com/gofiber/fiber/v2"
	middleware "github.com/ofjangra/netlynk_server/internal/middleware"
	transport_rest "github.com/ofjangra/netlynk_server/internal/transport/rest"
)

func Router(app *fiber.App) {

	authRequired := middleware.Authrequired()

	router := app

	// AUTH ROUTES

	router.Post("/api/signup", transport_rest.Signup)

	router.Post("/api/signin", transport_rest.Signin)

	router.Post("/api/profile/logout", authRequired, transport_rest.Logout)

	// LINK ROUTES

	router.Post("/api/createlink", authRequired, transport_rest.CreateOneLink)

	router.Put("/api/editlink/:id", authRequired, transport_rest.UpdateOneLink)

	router.Delete("/api/deletelink/:id", authRequired, transport_rest.DeleteOneLink)

	router.Get("/api/link/:id", transport_rest.GetALink)

	router.Get("/api/user/:id/links", transport_rest.GetAllLinks)

	//  PROFILE ROUTES

	router.Get("/api/profile", authRequired, transport_rest.Profile)

	router.Put("/api/editprofile", authRequired, transport_rest.EditProfile)

	router.Get("/api/user/:username", transport_rest.GetUser)

	router.Get("/api/test", authRequired, func(c *fiber.Ctx) error {

		return c.JSON(fiber.Map{"message": "Hello FIber"})

	})
}
