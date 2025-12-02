package main

import (
	"log"
	"numberniceic/database"
	"numberniceic/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDb()

	// Initialize template engine
	engine := html.New("./views", ".gohtml")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
