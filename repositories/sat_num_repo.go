package repositories

import (
	"numberniceic/models"

	"gorm.io/gorm"
)

// satNumRepository is the unexported, concrete implementation.
type satNumRepository struct {
	db *gorm.DB
}

// NewSatNumRepository is the exported constructor.
// It now returns the public interface type, hiding the implementation details.
func NewSatNumRepository(db *gorm.DB) NumericRepository {
	return &satNumRepository{db: db}
}

// GetByKey implements the NumericRepository interface.
func (s *satNumRepository) GetByKey(key string) (models.NumericValue, error) {
	var satNum models.SatNum
	result := s.db.Where("char_key = ?", key).First(&satNum)

	if result.Error != nil {
		return nil, result.Error
	}
	return satNum, nil
}
