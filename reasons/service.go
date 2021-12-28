package reasons

type Service interface {
	GetAllListReason() ([]Reasons, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllListReason() ([]Reasons, error) {
	listReasons, err := s.repository.GetAllListReason()
	if err != nil {
		return listReasons, err
	}
	return listReasons, nil
}
