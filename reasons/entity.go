package reasons

type Reasons struct {
	ID         string `gorm:"column:id" json:"ID"`
	Keterangan string `gorm:"column:keterangan" json:"Keterangan"`
	IsActive   int    `gorm:"column:is_active" json:"IsActive"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used `Reasons`
func (Reasons) TableName() string {
	return "ms_reasons"
}
