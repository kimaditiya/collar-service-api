package user

import (
	"gorm.io/gorm"
)

// buat interface untuk user
type Repository interface {
	Save(user User) (User, error)
	FindByNIK(nik string) (User, error)
	SaveUserRoles(userRoles UserRoles) (UserRoles, error)
	GetAllListUser() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllListUser() ([]User, error) {
	var listUser []User

	err := r.db.Table("users").Find(&listUser).Error
	if err != nil {
		return listUser, err
	}
	return listUser, nil
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

//Create User Roles
func (r *repository) SaveUserRoles(userRoles UserRoles) (UserRoles, error) {
	err := r.db.Create(&userRoles).Error

	if err != nil {
		return userRoles, err
	}
	return userRoles, nil
}
