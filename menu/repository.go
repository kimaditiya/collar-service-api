package menu

import "gorm.io/gorm"

type Repository interface {
	Savemenus(menus Menu) (Menu, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Savemenus(menu Menu) (Menu, error) {
	err := r.db.Create(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}
