package dkh

import (
	"fmt"
)

type Service interface {
	GetAllListDkh() ([]Dkh, error)
	CreateDkh(input CreateDkhInput) ([]Dkh, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllListDkh() ([]Dkh, error) {
	listdkh, err := s.repository.GetAllListDkh()
	if err != nil {
		return listdkh, err
	}
	return listdkh, nil
}

/**
	*@author wawan
	*Create DKH
**/
func (s *service) CreateDkh(input CreateDkhInput) ([]Dkh, error) {

	fmt.Printf("%+v\n", input.Dkh)

	newDkh, err := s.repository.CreateDkh(input.Dkh)
	if err != nil {
		return newDkh, err
	}

	return newDkh, nil
}
