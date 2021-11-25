package collectiongroup

type CreateCGInputFormatter struct {
	CollectionGrupAbbr string `json:"CollectionGroupAbbr"`
	// PrivelegeDesc      string `gorm:"foreignkey:PrivelegeCode"json:"PrivelegeDesc"`
	PrivelegeCode int `json:"PrivelegeCode"`
	// PrivelegeDesc Privelege `gorm:"foreignkey:PrivelegeCode"json:"PrivelegeDesc"`
	AreaID int `json:"AreaID"`
	// Area     []listarea.Area
	BranchID string `json:"BranchID"`
	PosID    string `json:"PosID"`
}
type CreateCGDetailInputFormatter struct {
	CollectionGroupID int    `gorm:"column:collection_group_id"json:"CollectionGroupID"`
	PrivelegeCode     int    `gorm:"column:privelege_code"json:"PrivelegeCode"`
	DataContain       string `gorm:"column:data_contain" json:"DataContain"`
}

func FormatCreateCG(collectionGroup CollectionGroup) CreateCGInputFormatter {
	formatter := CreateCGInputFormatter{
		CollectionGrupAbbr: collectionGroup.CollectionGrupAbbr,
		PrivelegeCode:      collectionGroup.PrivelegeCode,
		// PrivelegeDesc:      collectionGroup.PrivelegeDesc,
		AreaID: collectionGroup.AreaID,
		// Area:     collectionGroup.Area,
		BranchID: collectionGroup.BranchID,
		PosID:    collectionGroup.PosID,
	}
	return formatter
}
func FormatCreateCGDetail(collectionGroupDetail CollectionGroupDetail) CreateCGDetailInputFormatter {
	formatter := CreateCGDetailInputFormatter{
		CollectionGroupID: collectionGroupDetail.CollectionGroupID,
		PrivelegeCode:     collectionGroupDetail.PrivelegeCode,
		DataContain:       collectionGroupDetail.DataContain,
	}
	return formatter
}
