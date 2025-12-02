package routes

import (
	"numberniceic/database"
	"numberniceic/handlers"
	"numberniceic/repositories"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := database.DB

	// --- Repositories ---
	satNumRepo := repositories.NewSatNumRepository(db)
	shaNumRepo := repositories.NewShaNumRepository(db)
	// Create the new repository for number meanings
	meaningRepo := repositories.NewNumberMeaningRepository(db)

	// --- Handlers ---
	// Update the DecodeHandler to accept the new repository
	decodeHandler := handlers.NewDecodeHandler(satNumRepo, shaNumRepo, meaningRepo)

	// --- View Routes ---
	app.Get("/", handlers.APIList)
	app.Get("/search", handlers.SearchPage)

	// --- API Routes ---
	api := app.Group("/api")

	// The single, powerful endpoint for decoding and summing.
	api.Get("/decode/:name", decodeHandler.DecodeName)

	// The individual routes for debugging.
	api.Get("/satnums/:key", handlers.NewGenericNumHandler(satNumRepo))
	api.Get("/shanums/:key", handlers.NewGenericNumHandler(shaNumRepo))
}
