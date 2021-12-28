package assign

import (
	"collar-service-api/user"
	"collar-service-api/wilayah"
)

type Assign struct {
	KodeWilayah string          `gorm:"column:kd_wilayah" json:"KodeWilayah" binding"required"`
	IDLogin     int             `gorm:"column:id_login" json:"IDLogin"`
	UserCreate  string          `gorm:"column:user_create;default:System" json:"UserCreate"`
	User        user.User       `gorm:"foreignkey:id_login;references:id"`
	Wilayah     wilayah.Wilayah `gorm:"foreignkey:kd_wilayah;references:kd_wilayah"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func (Assign) TableName() string {
	return "tr_assign_col"
}
