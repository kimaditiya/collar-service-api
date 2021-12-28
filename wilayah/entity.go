package wilayah

import (
	"collar-service-api/collectiongroup"
	"collar-service-api/customers"
	"fmt"
	"strconv"
	"time"
)

type VwListWilayah struct {
	KodePos         string                          `gorm:"column:kodepos" json:"KodePos"`
	City            string                          `gorm:"column:city" json:"City"`
	District        string                          `gorm:"column:disctrict" json:"District"`
	SubDistrict     string                          `gorm:"column:subdisctrict" json:"SubDistrict"`
	R2              string                          `gorm:"column:R2" json:"R2"`
	R4              string                          `gorm:"column:R4" json:"R4"`
	KodeSubPos      string                          `gorm:"column:kd_subpos" json:"KodeSubPos"`
	BranchID        string                          `gorm:"column:BranchID" json:"BranchID"`
	WilayahDetail   WilayahDetail                   `gorm:"foreignkey:kd_subpos;references:kd_subpos"`
	CollectionGroup collectiongroup.CollectionGroup `gorm:"foreignkey:BranchID;references:branch_id"`
}

type Wilayah struct {
	KodeCg        string           `gorm:"column:kd_cg" json:"KodeCg"`
	KdWilayah     string           `gorm:"column:kd_wilayah" json:"KdWilayah"`
	AreaTagih     string           `gorm:"column:area_tagih" json:"AreaTagih"`
	UserCreate    string           `gorm:"column:user_create;default:System" json:"UserCreate"`
	BranchID      string           `gorm:"column:branch_id" json:"BranchID"`
	WilayahDetail []*WilayahDetail `gorm:"foreignKey:kd_wilayah;references:kd_wilayah"json:"WilayahDetail"`
}

type WilayahDetail struct {
	KdWilayah  string              `gorm:"column:kd_wilayah" json:"KdWilayah"`
	KodeSubPos string              `gorm:"column:kd_subpos" json:"KodeSubPos"`
	Customers  customers.Customers `gorm:"foreignkey:kd_subpos;references:kd_subpos"`
}

type Tabler interface {
	TableName() string
}

func (WilayahDetail) TableName() string {
	return "tr_wilayah_detail"
}

// TableName overrides the table name used by User to `profiles`
func (Wilayah) TableName() string {
	return "tr_wilayah_assign"
}

func (w *Wilayah) ModifyData(value int64) {
	w.KdWilayah = Numbering(value, w.BranchID)
}

func Numbering(value int64, brancid string) string {

	if value == 0 {
		value = 1
	} else {
		value = value + 1
	}

	return "KW" + brancid + strconv.Itoa(int(time.Now().Year())) + strconv.Itoa(int(time.Now().Month())) + fmt.Sprintf("%04d", value)
}
