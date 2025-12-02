package models

// SatNum is the struct for the sat_nums table.
type SatNum struct {
	CharKey  string `gorm:"column:char_key;primaryKey;type:text" json:"char_key"`
	SatValue int    `gorm:"column:sat_value;not null" json:"sat_value"`
}

// TableName explicitly sets the table name for GORM.
func (SatNum) TableName() string {
	return "public.sat_nums"
}

// GetValue implements the NumericValue interface, returning the integer value for SatNum.
func (s SatNum) GetValue() int {
	return s.SatValue
}
