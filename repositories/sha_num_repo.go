package repositories

import (
	"numberniceic/models"

	"gorm.io/gorm"
)

// shaNumRepository is the unexported, concrete implementation.
type shaNumRepository struct {
	db *gorm.DB
}

// NewShaNumRepository is the exported constructor.
// It now returns the public interface type.
func NewShaNumRepository(db *gorm.DB) NumericRepository {
	return &shaNumRepository{db: db}
}

// GetByKey implements the NumericRepository interface.
func (s *shaNumRepository) GetByKey(key string) (models.NumericValue, error) {
	var shaNum models.ShaNum
	result := s.db.Where("char_key = ?", key).First(&shaNum)

	if result.Error != nil {
		return nil, result.Error
	}
	return shaNum, nil
}
