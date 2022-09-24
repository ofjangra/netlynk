package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ofjangra/netlynk_server/controllers"
)

func Router(app *fiber.App) {

	router := app

	router.Post("/signup", controllers.Signup)

	router.Post("/signin", controllers.Signin)

	router.Post("/createlink", controllers.CreateOneLink)

}
