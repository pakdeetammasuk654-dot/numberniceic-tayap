package routes

import (
	"numberniceic/database"
	"numberniceic/handlers"
	"numberniceic/repositories"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := database.DB

	// --- ส่วนเดิม (SAT) ---
	satNumRepo := repositories.NewSatNumRepository(db)
	satNumHandler := handlers.NewSatNumHandler(satNumRepo)

	// --- ส่วนใหม่ (SHA) เพิ่มตรงนี้ ---
	shaNumRepo := repositories.NewShaNumRepository(db)
	shaNumHandler := handlers.NewShaNumHandler(shaNumRepo)

	api := app.Group("/api")

	// Routes สำหรับ SAT
	api.Get("/satnums/:key", satNumHandler.GetSatNum)

	// Routes สำหรับ SHA (เพิ่มใหม่)
	api.Get("/shanums/:key", shaNumHandler.GetShaNum)

}
