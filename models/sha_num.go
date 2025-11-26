package models

type ShaNum struct {
	// char_key เป็น Primary Key
	CharKey string `gorm:"column:char_key;primaryKey;type:text" json:"char_key"`

	// sha_value เปลี่ยนชื่อ field ให้ตรงกับตารางใหม่
	ShaValue int `gorm:"column:sha_value;not null" json:"sha_value"`
}

// ระบุชื่อตารางให้ชัดเจน
func (ShaNum) TableName() string {
	return "public.sha_nums"
}
