package collectiongroup

type Service interface {
	ListPrivelege() ([]Privelege, error)
	ListCG() ([]CollectionGroup, error)
	CreateCG(input CreateCGInput) (CollectionGroup, error)
	UpdateCG(input CreateCGInput) (CollectionGroup, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) ListPrivelege() ([]Privelege, error) {
	listPrivelege, err := s.repository.ListPrivelege()

	if err != nil {
		return listPrivelege, err
	}

	return listPrivelege, nil
}

/**
	*@author wawan
	*Create Collectioun Group
**/
func (s *service) CreateCG(input CreateCGInput) (CollectionGroup, error) {

	newCollectionGroup, err := s.repository.CreateCollectionGroup(input.CollectionGroup)
	if err != nil {
		return newCollectionGroup, err
	}

	return newCollectionGroup, nil
}

/**
	*@author wawan
	*List Collectioun Group
**/
func (s *service) ListCG() ([]CollectionGroup, error) {
	listCollection, err := s.repository.ListCollectionGroup()

	if err != nil {
		return listCollection, err
	}

	return listCollection, nil
}

/**
	*@author wawan
	*Update Collectioun Group
**/
func (s *service) UpdateCG(input CreateCGInput) (CollectionGroup, error) {

	updateCollectionGroup, err := s.repository.UpdateCollectionGroup(input.CollectionGroup)
	if err != nil {
		return updateCollectionGroup, err
	}

	return updateCollectionGroup, nil
}
