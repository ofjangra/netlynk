package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	db_config "github.com/ofjangra/netlynk_server/internal/config/db"
	"github.com/ofjangra/netlynk_server/internal/routes"
)

// func init() {
// 	envLoadErr := godotenv.Load(".env")

// 	if envLoadErr != nil {
// 		log.Fatal("Failed to load environment variables 1", envLoadErr)
// 	}

// }

func App() {

	db_config.GetCollections()

	port := os.Getenv("PORT")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	app.Static("/", "../web")

	routes.Router(app)

	indexPath, pathErr := filepath.Abs("../web/index.html")

	if pathErr != nil {
		fmt.Println(pathErr)
	}
	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendFile(indexPath)
	})

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}

}
