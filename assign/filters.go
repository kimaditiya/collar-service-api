package assign

type AssignFilter struct {
	KodeWilayah string `gorm:"column:kd_wilayah" json:"KodeWilayah" binding"required"`
	IDLogin     string `gorm:"column:id_login" json:"IDLogin"`
	UserCreate  string `gorm:"column:user_create;default:System" json:"UserCreate"`
}
