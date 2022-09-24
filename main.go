package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/ofjangra/netlynk_server/routes"
)

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")

	app := fiber.New()

	routes.Router(app)

	err := app.Listen(port)
	if err != nil {
		log.Fatal(err)
	}

}
