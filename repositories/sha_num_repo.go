package repositories

import (
	"numberniceic/models"

	"gorm.io/gorm"
)

// 1. Interface
type ShaNumRepository interface {
	GetByKey(key string) (*models.ShaNum, error)
}

// 2. Implementation
type shaNumRepository struct {
	db *gorm.DB
}

func NewShaNumRepository(db *gorm.DB) ShaNumRepository {
	return &shaNumRepository{db: db}
}

func (r *shaNumRepository) GetByKey(key string) (*models.ShaNum, error) {
	var shaNum models.ShaNum
	// ค้นหาจากตาราง sha_nums
	result := r.db.Where("char_key = ?", key).First(&shaNum)
	if result.Error != nil {
		return nil, result.Error
	}
	return &shaNum, nil
}
