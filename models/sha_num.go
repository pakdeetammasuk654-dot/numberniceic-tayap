package models

// ShaNum is the struct for the sha_nums table.
type ShaNum struct {
	CharKey  string `gorm:"column:char_key;primaryKey;type:text" json:"char_key"`
	ShaValue int    `gorm:"column:sha_value;not null" json:"sha_value"`
}

// TableName explicitly sets the table name for GORM.
func (ShaNum) TableName() string {
	return "public.sha_nums"
}

// GetValue implements the NumericValue interface, returning the integer value for ShaNum.
func (s ShaNum) GetValue() int {
	return s.ShaValue
}
