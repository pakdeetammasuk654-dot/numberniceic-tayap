package repositories

import (
	"numberniceic/models"

	"gorm.io/gorm"
)

type SatNumRepository interface {
	GetByKey(key string) (*models.SatNum, error)
}

type satNumRepository struct {
	db *gorm.DB
}

func (s *satNumRepository) GetByKey(key string) (*models.SatNum, error) {
	var satNum models.SatNum
	result := s.db.Where("char_key = ?", key).First(&satNum)

	if result.Error != nil {
		return nil, result.Error
	}
	return &satNum, nil
}

func NewSatNumRepository(db *gorm.DB) SatNumRepository {
	return &satNumRepository{db: db}
}
