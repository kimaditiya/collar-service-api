package dkh

import (
	"collar-service-api/customers"
	"strconv"
	"time"
)

type Dkh struct {
	TransNo   string              `gorm:"column:trans_no" json:"TransNo"`
	BranchId  string              `gorm:"column:branch_id" json:"BranchId"`
	Periode   int                 `gorm:"column:periode" json:"Periode"`
	NoKontrak string              `gorm:"column:no_kontrak" json:"NoKontrak"`
	Customers customers.Customers `gorm:"foreignkey:no_kontrak,periode;references:AgreementNo,Periode"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func (Dkh) TableName() string {
	return "tr_dkh"
}

func (d *Dkh) ModifyData() {
	d.TransNo = "TR" + strconv.Itoa(d.Periode) + d.NoKontrak + "-" + strconv.Itoa(int(time.Now().Year())) + strconv.Itoa(int(time.Now().Month()))
}
