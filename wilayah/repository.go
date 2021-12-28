package wilayah

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	GetAllListWilayah() ([]VwListWilayah, error)
	GetAllListWilayahAssign() ([]Wilayah, error)
	CreateWilayah(wilayah Wilayah) (Wilayah, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllListWilayah() ([]VwListWilayah, error) {
	var listwilayah []VwListWilayah

	err := r.db.Table("vw_listwilayah").Preload("WilayahDetail").Preload("CollectionGroup").Find(&listwilayah).Error
	if err != nil {
		return listwilayah, err
	}
	return listwilayah, nil
}

func (r *repository) GetAllListWilayahAssign() ([]Wilayah, error) {
	var listwilayah []Wilayah

	err := r.db.Table("tr_wilayah_assign").Preload("WilayahDetail").Find(&listwilayah).Error
	if err != nil {
		return listwilayah, err
	}
	return listwilayah, nil
}

func (r *repository) CreateWilayah(wilayah Wilayah) (Wilayah, error) {

	var result int64

	err2 := r.db.Table("tr_wilayah_detail").Count(&result).Error
	if err2 != nil {
		log.Fatal(err2)
	}

	wilayah.ModifyData(result)

	err := r.db.Debug().Table("tr_wilayah_assign").Create(&wilayah).Error

	if err != nil {
		return wilayah, err
	}

	return wilayah, nil
}
