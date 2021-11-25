package listarea

import "gorm.io/gorm"

type Repository interface {
	GetAllListArea() ([]Area, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllListArea() ([]Area, error) {
	var listArea []Area
	err := r.db.Table("ms_area").Find(&listArea).Error
	if err != nil {
		return listArea, err
	}
	return listArea, nil
}
