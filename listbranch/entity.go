package listbranch

import "time"

type Branch struct {
	ID                 int       `gorm:"column:id" json:"ID"`
	BranchID           string    `gorm:"column:branchID" json:"BranchID"`
	BranchFullName     string    `gorm:"column:branchFullName" json:"BranchFullName"`
	AreaID             string    `gorm:"column:areaID" json:"AreaID"`
	BranchInitialName  string    `gorm:"column:branchInitialName" json:"BranchInitialName"`
	BranchAddress      string    `gorm:"column:branchAddress" json:"BranchAddress"`
	BranchDisctrict    string    `gorm:"column:branchDisctrict" json:"BranchDisctrict"`
	BranchSubDisctrict string    `gorm:"column:branchSubDisctrict" json:"BranchSubDisctrict"`
	BranchCity         string    `gorm:"column:branchCity" json:"BranchCity"`
	BranchZipCode      string    `gorm:"column:branchZipCode" json:"BranchZipCode"`
	BranchLong         float32   `gorm:"column:branchLong" json:"BranchLong"`
	BranchLat          float32   `gorm:"column:branchLat" json:"BranchLat"`
	CreatedAt          time.Time `gorm:"column:cratedAt" json:"CreatedAt"`
	CreatedBy          string    `gorm:"column:cratedBy" json:"CreatedBy"`
	UpdatedAt          time.Time `gorm:"column:updatedAt" json:"UpdatedAt"`
	UpdatedBy          string    `gorm:"column:updatedBy" json:"UpdatedBy"`
}
type BranchPos struct {
	ID             int    `gorm:"column:id" json:"ID"`
	PosID          string `gorm:"column:posID" json:"PosID"`
	BranchID       string `gorm:"column:"branch_id" json:"BranchID"`
	PosName        string `gorm:"column:posName" json:"PosName"`
	PosDescription string `gorm:"column:posDescription" json:"PosDescription"`
	PCode          string `gorm:"column:pCode" json:"PCode"`
	IsPos          int    `gorm:"column:isPos" json:"IsPos"`
	// Branch         []Branch `gorm:"foreignkey:BranchID" `
	// `gorm:"many2many:branch_id"`
	// `gorm:"column:branch_id" json:"BranchID"`
}
