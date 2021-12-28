package dkh

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAllListDkh() ([]Dkh, error)
	CreateDkh(dkh []Dkh) ([]Dkh, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllListDkh() ([]Dkh, error) {
	var listdkh []Dkh

	err := r.db.Table("tr_dkh").Preload("Customers").Find(&listdkh).Error
	if err != nil {
		return listdkh, err
	}
	return listdkh, nil
}

func (r *repository) CreateDkh(dkh []Dkh) ([]Dkh, error) {

	for i := range dkh {
		dkh[i].ModifyData()
	}

	err := r.db.Debug().Table("tr_dkh").Create(&dkh).Error

	if err != nil {
		return dkh, err
	}

	return dkh, nil
}
