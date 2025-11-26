package handlers

import (
	"numberniceic/repositories"

	"github.com/gofiber/fiber/v2"
)

type SatNumHandler struct {
	repo repositories.SatNumRepository
}

func NewSatNumHandler(repo repositories.SatNumRepository) *SatNumHandler {
	return &SatNumHandler{repo: repo}
}

func (h *SatNumHandler) GetSatNum(c *fiber.Ctx) error {
	key := c.Params("key")

	// เรียกใช้ Repository แทนการเรียก DB ตรงๆ
	satNum, err := h.repo.GetByKey(key)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Data not found"})
	}

	return c.JSON(satNum)
}
