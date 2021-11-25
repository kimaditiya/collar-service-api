package listbranch

type BranchByArea struct {
	ID                 int     `gorm:"column:id" json:"ID"`
	BranchID           string  `gorm:"column:branchID" json:"BranchID"`
	BranchFullName     string  `gorm:"column:branchFullName" json:"BranchFullName"`
	AreaID             string  `gorm:"column:areaID" json:"AreaID"`
	BranchInitialName  string  `gorm:"column:branchInitialName" json:"BranchInitialName"`
	BranchAddress      string  `gorm:"column:branchAddress" json:"BranchAddress"`
	BranchDisctrict    string  `gorm:"column:branchDisctrict" json:"BranchDisctrict"`
	BranchSubDisctrict string  `gorm:"column:branchSubDisctrict" json:"BranchSubDisctrict"`
	BranchCity         string  `gorm:"column:branchCity" json:"BranchCity"`
	BranchZipCode      string  `gorm:"column:branchZipCode" json:"BranchZipCode"`
	BranchLong         float32 `gorm:"column:branchLong" json:"BranchLong"`
	BranchLat          float32 `gorm:"column:branchLat" json:"BranchLat"`
}

func FormatBranchByArea(branch Branch) BranchByArea {
	formatter := BranchByArea{
		ID:                 branch.ID,
		BranchID:           branch.BranchID,
		BranchFullName:     branch.BranchFullName,
		AreaID:             branch.AreaID,
		BranchInitialName:  branch.BranchInitialName,
		BranchAddress:      branch.BranchAddress,
		BranchDisctrict:    branch.BranchDisctrict,
		BranchSubDisctrict: branch.BranchSubDisctrict,
		BranchCity:         branch.BranchCity,
		BranchZipCode:      branch.BranchZipCode,
		BranchLong:         branch.BranchLong,
		BranchLat:          branch.BranchLat,
	}
	return formatter
}

type BranchByPosFormatter struct {
	ID             int    `json:"id"`
	PosID          string `json:"posID"`
	BranchID       string `json:"branch_id"`
	PosName        string `json:"posName"`
	PosDescription string `json:"posDescription"`
	PCode          string `json:"pCode"`
	IsPos          int    `json:"isPos"`
	// Branch         []Branch `json:"Branches" gorm:"foreignkey:BranchID"`
}

func FormatBranchByPos(branchPos BranchPos) BranchByPosFormatter {
	formatter := BranchByPosFormatter{
		ID:             branchPos.ID,
		BranchID:       branchPos.BranchID,
		PosID:          branchPos.PosID,
		PosName:        branchPos.PosName,
		PosDescription: branchPos.PosDescription,
		PCode:          branchPos.PCode,
		IsPos:          branchPos.IsPos,
		// Branch:         []Branch{},
	}
	return formatter
}
