package collectiongroup

type Service interface {
	ListPrivelege() ([]Privelege, error)
	CreateCG(input CreateCGInput) (CollectionGroup, error)
	// CreateCG() (CollectionGroup, error)
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
func (s *service) CreateCG(input CreateCGInput) (CollectionGroup, error) {
	collectionGroup := CollectionGroup{}
	collectionGroup.CollectionGrupAbbr = input.CollectionGrupAbbr
	collectionGroup.PrivelegeCode = input.PrivelegeCode
	collectionGroup.AreaID = input.AreaID
	// collectionGroup.BranchID = input.BranchID
	// collectionGroup.PosID = input.PosID
	newCollectionGroup, err := s.repository.CreateCollectionGroup(collectionGroup)
	if err != nil {
		return newCollectionGroup, err
	}
	return newCollectionGroup, nil
}
