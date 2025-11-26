package main

import (
	"log"
	"numberniceic/database"
	"numberniceic/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDb()
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
