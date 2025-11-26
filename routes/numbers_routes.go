package routes

import (
	"numberniceic/database"
	"numberniceic/handlers"
	"numberniceic/repositories"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// 1. เตรียม DB (สมมติว่า connect แล้วจาก database.DB)
	db := database.DB

	// 2. สร้าง Repository (Layer Database)
	satNumRepo := repositories.NewSatNumRepository(db)

	// 3. สร้าง Handler โดยส่ง Repo เข้าไป (Layer Logic/HTTP)
	satNumHandler := handlers.NewSatNumHandler(satNumRepo)

	// 4. กำหนด Route
	api := app.Group("/api")
	api.Get("/satnums/:key", satNumHandler.GetSatNum)
}
