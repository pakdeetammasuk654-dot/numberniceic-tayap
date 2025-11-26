package models

// SatNum คือ struct ที่ map กับตาราง sat_nums
type SatNum struct {
	CharKey string `gorm:"column:char_key;primaryKey;type:text" json:"char_key"`

	SatValue int `gorm:"column:sat_value;not null" json:"sat_value"`
}
