package collectiongroup

type CreateCGInput struct {
	CollectionGrupAbbr string `json:"CollectionGroupAbbr"`
	PrivelegeCode      int    `json:"PrivelegeCode"`
	AreaID             int    `json:"AreaID"`
	BranchID           string `json:"BranchID"`
	PosID              string `json:"PosID"`

	CollectionGroupDetail []CollectionGroupDetail `json:"CollectionGroupDetail"`
	// CollectionGroupID     int                     `json:"CollectionGroupID"`
	// DataContain           string                  `json:"DataContain"`
}

type CreateCGDetailInput struct {
	CollectionGroupID int    `json:"CollectionGroupID"`
	PrivelegeCode     int    `json:"PrivelegeCode"`
	DataContain       string `json:"DataContain"`
}
