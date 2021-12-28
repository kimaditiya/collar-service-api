package reasons

import "gorm.io/gorm"

type Repository interface {
	GetAllListReason() ([]Reasons, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllListReason() ([]Reasons, error) {
	var reasons []Reasons
	err := r.db.Table("ms_reasons").Find(&reasons).Error
	if err != nil {
		return reasons, err
	}
	return reasons, nil
}
