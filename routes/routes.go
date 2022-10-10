package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/netlynk_server/controllers"
	"github.com/ofjangra/netlynk_server/middleware"
)

func Router(app *fiber.App) {

	authRequired := middleware.Authrequired()

	router := app

	// AUTH ROUTES

	router.Post("/signup", controllers.Signup)

	router.Post("/signin", controllers.Signin)

	router.Post("/profile/logout", authRequired, controllers.Logout)
	// LINK ROUTES

	router.Post("/createlink", authRequired, controllers.CreateOneLink)

	router.Put("/editlink/:id", authRequired, controllers.UpdateOneLink)

	router.Delete("/deletelink/:id", authRequired, controllers.DeleteOneLink)

	router.Get("/link/:id", controllers.GetALink)

	router.Get("/links/:id", controllers.GetAllLinks)

	//  PROFILE ROUTES

	router.Get("/profile", authRequired, controllers.Profile)

	router.Put("/editprofile", authRequired, controllers.EditProfile)

	router.Put("/editprofile/photo", authRequired, controllers.EditProfilePhoto)

	router.Get("/user/:username", controllers.GetUser)

	router.Get("/test", authRequired, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello FIber"})
	})
}
