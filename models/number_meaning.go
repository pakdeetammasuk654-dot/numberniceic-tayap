package models

// NumberMeaning maps to the 'numbers' table in the database.
type NumberMeaning struct {
	PairNumberID  int    `gorm:"column:pairnumberid;primaryKey" json:"-"` // Hide from json output
	PairType      string `gorm:"column:pairtype" json:"pair_type"`
	PairNumber    string `gorm:"column:pairnumber" json:"pair_number"`
	PairPoint     int    `gorm:"column:pairpoint" json:"pair_point"`
	MiracleDetail string `gorm:"column:miracledetail" json:"miracle_detail"`
	MiracleDesc   string `gorm:"column:miracledesc" json:"miracle_desc"`
	DetailVip     string `gorm:"column:detail_vip" json:"detail_vip"`
}

// TableName explicitly sets the table name for GORM.
func (NumberMeaning) TableName() string {
	return "public.numbers"
}
