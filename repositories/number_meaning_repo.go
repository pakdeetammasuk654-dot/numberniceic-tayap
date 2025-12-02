package repositories

import (
	"numberniceic/models"

	"gorm.io/gorm"
)

// NumberMeaningRepository defines the interface for fetching number meanings.
// The function is renamed to reflect that it only uses the pair number.
type NumberMeaningRepository interface {
	FindByPairNumber(pairNumber string) (*models.NumberMeaning, error)
}

// numberMeaningRepository is the concrete implementation.
type numberMeaningRepository struct {
	db *gorm.DB
}

// NewNumberMeaningRepository creates a new repository for number meanings.
func NewNumberMeaningRepository(db *gorm.DB) NumberMeaningRepository {
	return &numberMeaningRepository{db: db}
}

// FindByPairNumber now searches for a number meaning only by the 2-digit number,
// as per the new requirement. The pairtype column is ignored.
func (r *numberMeaningRepository) FindByPairNumber(pairNumber string) (*models.NumberMeaning, error) {
	var meaning models.NumberMeaning
	// The query now ONLY uses pairnumber.
	result := r.db.Where("pairnumber = ?", pairNumber).First(&meaning)
	if result.Error != nil {
		return nil, result.Error
	}
	return &meaning, nil
}
