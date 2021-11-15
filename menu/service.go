package menu

type Service interface {
	MenuInput(input MenuInput) (Menu, error)
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

	newMenu, err := s.repository.Savemenus(menu)
	if err != nil {
		return newMenu, err
	}

	return newMenu, nil

}
