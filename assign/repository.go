package assign

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllListAssign(filter *AssignFilter) ([]Assign, error)
	CreateAssign(assign []Assign) ([]Assign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllListAssign(filter *AssignFilter) ([]Assign, error) {

	var listwilayah []Assign
	err := r.db.Table("tr_assign_col").Where(filter).Preload("Wilayah.WilayahDetail.Customers").Preload("User").Find(&listwilayah).Error
	if err != nil {
		return listwilayah, err
	}
	return listwilayah, nil
}

func (r *repository) CreateAssign(assign []Assign) ([]Assign, error) {

	err := r.db.Debug().Table("tr_assign_col").Create(&assign).Error

	if err != nil {
		return assign, err
	}

	return assign, nil
}
