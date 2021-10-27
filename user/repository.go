package user

import (
	"gorm.io/gorm"
)

// buat interface untuk user
type Repository interface {
	Save(user User) (User, error)
	FindByNIK(nik string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// function Save
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

//Find By NIK
func (r *repository) FindByNIK(nik string) (User, error) {
	var user User

	err := r.db.Where("nik = ?", nik).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
