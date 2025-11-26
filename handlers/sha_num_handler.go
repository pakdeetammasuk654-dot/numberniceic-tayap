package handlers

import (
	"numberniceic/repositories"

	"github.com/gofiber/fiber/v2"
)

type ShaNumHandler struct {
	repo repositories.ShaNumRepository
}

func NewShaNumHandler(repo repositories.ShaNumRepository) *ShaNumHandler {
	return &ShaNumHandler{repo: repo}
}

// GET /api/shanums/:key
func (h *ShaNumHandler) GetShaNum(c *fiber.Ctx) error {
	key := c.Params("key")
	shaNum, err := h.repo.GetByKey(key)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Data not found"})
	}
	return c.JSON(shaNum)
}
