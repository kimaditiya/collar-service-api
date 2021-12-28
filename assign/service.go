package assign

import "fmt"

type Service interface {
	GetAllListAssign(filter *AssignFilter) ([]Assign, error)
	CreateAssign(input CreateAssign) ([]Assign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllListAssign(filter *AssignFilter) ([]Assign, error) {
	listAssign, err := s.repository.GetAllListAssign(filter)
	if err != nil {
		return listAssign, err
	}
	return listAssign, nil
}

/**
	*@author wawan
	*Create Assign
**/
func (s *service) CreateAssign(input CreateAssign) ([]Assign, error) {

	fmt.Printf("%+v\n", input.Assign)
	newAssign, err := s.repository.CreateAssign(input.Assign)
	if err != nil {
		return newAssign, err
	}

	return newAssign, nil
}
