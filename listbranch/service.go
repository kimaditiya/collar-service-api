package listbranch

type Service interface {
	GetAllListBranch() ([]Branch, error)
	GetAllListBranchPos() ([]BranchPos, error)
	BranchByArea(input BranchByAreaInput) ([]Branch, error)
	FindPosByBranch(input BranchByPosInput) ([]BranchPos, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllListBranch() ([]Branch, error) {
	listBranch, err := s.repository.GetAllListBranch()
	if err != nil {
		return listBranch, err
	}
	return listBranch, nil
}
func (s *service) BranchByArea(input BranchByAreaInput) ([]Branch, error) {
	branchByArea := input.AreaID

	branchFindByAreaID, err := s.repository.BranchByArea(branchByArea)

	if err != nil {
		return branchFindByAreaID, err
	}

	return branchFindByAreaID, nil
}
func (s *service) GetAllListBranchPos() ([]BranchPos, error) {
	listBranchPos, err := s.repository.GetAllListBranchPos()
	if err != nil {
		return listBranchPos, err
	}
	return listBranchPos, nil
}
func (s *service) FindPosByBranch(input BranchByPosInput) ([]BranchPos, error) {
	posByBranch := input.PosID

	posFindByBranch, err := s.repository.FindPosByBranch(posByBranch)

	if err != nil {
		return posFindByBranch, err
	}

	return posFindByBranch, nil
}
