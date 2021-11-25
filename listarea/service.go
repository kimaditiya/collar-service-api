package listarea

type Service interface {
	GetAllListArea() ([]Area, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllListArea() ([]Area, error) {
	listArea, err := s.repository.GetAllListArea()
	if err != nil {
		return listArea, err
	}
	return listArea, nil
}
