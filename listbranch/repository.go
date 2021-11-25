package listbranch

import "gorm.io/gorm"

type Repository interface {
	GetAllListBranch() ([]Branch, error)
	GetAllListBranchPos() ([]BranchPos, error)
	BranchByArea(AreaID string) ([]Branch, error)
	FindPosByBranch(PosID string) ([]BranchPos, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllListBranch() ([]Branch, error) {
	var listBranch []Branch
	err := r.db.Table("ms_branch").Find(&listBranch).Error
	if err != nil {
		return listBranch, err
	}
	return listBranch, nil
}
func (r *repository) BranchByArea(AreaID string) ([]Branch, error) {
	var branchByArea []Branch

	err := r.db.Where("areaID = ?", AreaID).Table("ms_branch").Find(&branchByArea).Error
	if err != nil {
		return branchByArea, err
	}

	return branchByArea, nil
}
func (r *repository) GetAllListBranchPos() ([]BranchPos, error) {
	var listBranchPos []BranchPos
	err := r.db.Where("posID").Table("ms_branch_pos").Find(&listBranchPos).Error
	if err != nil {
		return listBranchPos, err
	}
	return listBranchPos, nil
}
func (r *repository) FindPosByBranch(PosID string) ([]BranchPos, error) {
	var posByBranch []BranchPos

	err := r.db.Where("posID = ?", PosID).Table("ms_branch_pos").Find(&posByBranch).Error
	if err != nil {
		return posByBranch, err
	}

	return posByBranch, nil
}
