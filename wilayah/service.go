package wilayah

import (
	"fmt"
)

type Service interface {
	GetAllListWilayah() ([]VwListWilayah, error)
	CreateWilayah(input CreateWilayah) (Wilayah, error)
	GetAllListWilayahAssign() ([]Wilayah, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllListWilayah() ([]VwListWilayah, error) {
	listWilayah, err := s.repository.GetAllListWilayah()
	if err != nil {
		return listWilayah, err
	}
	return listWilayah, nil
}

func (s *service) GetAllListWilayahAssign() ([]Wilayah, error) {
	listWilayah, err := s.repository.GetAllListWilayahAssign()
	if err != nil {
		return listWilayah, err
	}
	return listWilayah, nil
}

/**
	*@author wawan
	*Create Wilayah
**/
func (s *service) CreateWilayah(input CreateWilayah) (Wilayah, error) {

	fmt.Printf("%+v\n", input.Wilayah)

	newWilayah, err := s.repository.CreateWilayah(input.Wilayah)
	if err != nil {
		return newWilayah, err
	}

	return newWilayah, nil
}
