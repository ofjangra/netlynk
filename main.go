package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/ofjangra/netlynk_server/routes"
)

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))
	routes.Router(app)

	err := app.Listen(port)
	if err != nil {
		log.Fatal(err)
	}

}
