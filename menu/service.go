package menu

type Service interface {
	MenuInput(input MenuInput) (Menu, error)
	GetAllMenu() ([]Menu, error)
	GetAllMenusParent() ([]GetAllMenuParent, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) MenuInput(input MenuInput) (Menu, error) {
	menu := Menu{}
	menu.MenuCode = input.MenuCode
	menu.MenuTitle = input.MenuTitle
	menu.Url = input.Url
	menu.MenuIcons = input.MenuIcons

	if menu.ParentID != menu.ID {
		menu.MenuIcons = input.MenuIcons

		newMenu, err := s.repository.Savemenus(menu)
		return newMenu, err
	} else {
		menu.ParentID = input.ParentID
		newMenu, err := s.repository.SaveParentMenus(menu)
		return newMenu, err
	}
	// newMenu, err := s.repository.Savemenus(menu)
	// if err != nil {
	// 	return newMenu, err
	// }

	// return newMenu, nil

}

func (s *service) GetAllMenu() ([]Menu, error) {
	menus, err := s.repository.GetAllMenus()
	if err != nil {
		return menus, err
	}
	return menus, nil
}

func (s *service) GetAllMenusParent() ([]GetAllMenuParent, error) {
	menus, err := s.repository.GetAllMenusParent()
	if err != nil {
		return menus, err
	}
	return menus, nil
}
