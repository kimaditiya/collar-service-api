package menu

import "gorm.io/gorm"

type Repository interface {
	Savemenus(menus Menu) (Menu, error)
	SaveParentMenus(menu Menu) (Menu, error)
	GetAllMenus() ([]Menu, error)
	GetAllMenusParent() ([]GetAllMenuParent, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Savemenus(menu Menu) (Menu, error) {
	err := r.db.Preload("MenuParent").Create(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (r *repository) SaveParentMenus(menu Menu) (Menu, error) {
	err := r.db.Table("menus").Create(&menu).Error
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (r *repository) GetAllMenus() ([]Menu, error) {
	var menus []Menu
	// err := r.db.Table("menus").Find(&menus).Error
	err := r.db.Debug().Preload("MenuChild").Where("parent_id = 0").Table("menus").Find(&menus).Error
	if err != nil {
		return menus, err
	}
	return menus, nil
}

func (r *repository) GetAllMenusParent() ([]GetAllMenuParent, error) {
	var menus []GetAllMenuParent
	// err := r.db.Table("menus").Find(&menus).Error
	err := r.db.Debug().Where("parent_id = 0").Table("menus").Find(&menus).Error
	if err != nil {
		return menus, err
	}
	return menus, nil
}
