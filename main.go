package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ofjangra/netlynk_server/routes"
)

func main() {

	// godotenv.Load(".env")

	port := os.Getenv("PORT")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	app.Static("/", "./dist")

	routes.Router(app)

	indexPath, pathErr := filepath.Abs("./dist/index.html")

	if pathErr != nil {
		fmt.Println(pathErr)
	}
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile(indexPath)
	})

	err := app.Listen(port)
	if err != nil {
		log.Fatal(err)
	}

}
