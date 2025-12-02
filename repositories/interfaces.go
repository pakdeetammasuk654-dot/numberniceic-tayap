package repositories

import "numberniceic/models"

// NumericRepository defines a standard interface for repositories that fetch
// a numeric value model by a string key. This abstraction allows handlers
// to be independent of the concrete repository implementation.
type NumericRepository interface {
	GetByKey(key string) (models.NumericValue, error)
}
