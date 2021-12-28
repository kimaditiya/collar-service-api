package collectiongroup

import (
	"fmt"
	"strconv"
	"time"
)

type CollectionGroup struct {
	KodeCg                string                   `gorm:"column:kode_cg;primaryKey" json:"KodeCg"`
	CollectionGroupAbbr   string                   `gorm:"column:collection_group_abbr"json:"CollectionGroupAbbr"`
	PrivelegeCode         int                      `gorm:"column:privelege_code"json:"PrivelegeCode"`
	AreaID                int                      `gorm:"column:area_id"json:"AreaID"`
	BranchID              string                   `gorm:"column:branch_id"json:"BranchID"`
	PosID                 string                   `gorm:"column:pos_id"json:"PosID"`
	CreatedAt             time.Time                `gorm:"column:createdAt"json:"CreatedAt"`
	CreatedBy             string                   `gorm:"column:createdBy"json:"CreatedBy"`
	UpdatedAt             time.Time                `gorm:"column:updatedAt"json:"UpdatedAt"`
	UpdatedBy             string                   `gorm:"column:updateBy"json:"UpdatedBy"`
	CollectionGroupDetail []*CollectionGroupDetail `gorm:"foreignKey:kode_cg;references:kode_cg"json:"CollectionGroupDetail"`
}

type CollectionGroupDetail struct {
	KodeCg        string    `gorm:"column:kode_cg" json:"KodeCg"`
	PrivelegeCode int       `gorm:"column:privelege_code" json:"PrivelegeCode"`
	DataContain   string    `gorm:"column:data_contain" json:"DataContain"`
	CreatedAt     time.Time `gorm:"column:createdAt"json:"CreatedAt"`
	CreatedBy     string    `gorm:"column:createdBy"json:"CreatedBy"`
	UpdatedAt     time.Time `gorm:"column:updatedAt"json:"UpdatedAt"`
	UpdatedBy     string    `gorm:"column:updateBy"json:"UpdatedBy"`
}

type Privelege struct {
	PrivelegeCode string    `gorm:"column:privelege_code"json:"PrivelegeCode"`
	PrivelegeDesc string    `gorm:"column:privelege_desc"json:"PrivelegeDesc"`
	CreatedAt     time.Time `gorm:"column:createdAt" json:"CreatedAt"`
	CreatedBy     string    `gorm:"column:createdBy"json:"CreatedBy"`
	UpdatedAt     time.Time `gorm:"column:updatedAt"json:"UpdatedAt"`
	UpdateBy      string    `gorm:"column:updateBy"json:"UpdateBy"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides
func (CollectionGroupDetail) TableName() string {
	return "ms_collection_group_detail"
}

func (CollectionGroup) TableName() string {
	return "ms_collection_group"
}

func (u *CollectionGroup) Modify(value int64) {
	u.KodeCg = Numbering(value)
}

func Numbering(value int64) string {

	if value == 0 {
		value = 1
	} else {
		value = value + 1
	}

	return "CG" + strconv.Itoa(int(time.Now().Year())) + strconv.Itoa(int(time.Now().Month())) + fmt.Sprintf("%04d", value)
}
