package collectiongroup

import (
	"time"
)

type CollectionGroup struct {
	// gorm.Model
	CollectionGrupAbbr string    `gorm:"column:collection_group_abbr"json:"CollectionGrupAbbr"`
	PrivelegeCode      int       `gorm:"column:privelege_code"json:"PrivelegeCode"`
	AreaID             int       `gorm:"column:area_id"json:"AreaID"`
	BranchID           string    `gorm:"column:branch_id"json:"BranchID"`
	PosID              string    `gorm:"column:pos_id"json:"PosID"`
	CreatedAt          time.Time `gorm:"column:createdAt"json:"CreatedAt"`
	CreatedBy          string    `gorm:"column:createdBy"json:"CreatedBy"`
	UpdatedAt          time.Time `gorm:"column:updatedAt"json:"UpdatedAt"`
	UpdatedBy          string    `gorm:"column:updateBy"json:"UpdatedBy"`
	// Area               []listarea.Area `gorm:"foreignkey:AreaID"json:"Area"`
	// PrivelegeDesc      string    `gorm:"column:privelege_desc" json:"PrivelegeDesc"`
}

type CollectionGroupDetail struct {
	CollectionGroupID int       `gorm:"column:collection_group_id"json:"CollectionGroupID"`
	PrivelegeCode     int       `gorm:"column:privelege_code"json:"PrivelegeCode"`
	DataContain       string    `gorm:"column:data_contain" json:"DataContain"`
	CreatedAt         time.Time `gorm:"column:created_at"json:"CreatedAt"`
	CreatedBy         string    `gorm:"column:created_by"json:"CreatedBy"`
	UpdatedAt         time.Time `gorm:"column:updated_at"json:"UpdatedAt"`
	UpdatedBy         string    `gorm:"column:updated_by"json:"UpdatedBy"`
}

type Privelege struct {
	PrivelegeCode string    `gorm:"column:privelege_code"json:"PrivelegeCode"`
	PrivelegeDesc string    `gorm:"column:privelege_desc"json:"PrivelegeDesc"`
	CreatedAt     time.Time `gorm:"column:createdAt" json:"CreatedAt"`
	CreatedBy     string    `gorm:"column:createdBy"json:"CreatedBy"`
	UpdatedAt     time.Time `gorm:"column:updatedAt"json:"UpdatedAt"`
	UpdateBy      string    `gorm:"column:updateBy"json:"UpdateBy"`
}
