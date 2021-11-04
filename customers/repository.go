package customers

import (
	"gorm.io/gorm"
)

// buat interface untuk user
type Repository interface {
	ListOfCustomer(agreementno string) ([]Customers, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ListOfCustomer(AgreementNo string) ([]Customers, error) {

	//var rstcustomers Customers
	var rstcustomers []Customers

	err := r.db.Debug().Table("tr_coll_loc").Select("RYM, BranchID, ApplicationID, AgreementNo, FullName, Periode, Product, ProductID, Model, DueDate, InstallmentAmount, LateChargeAmount, MinimalPayment, AmountTobePaid, OvdDays, CollectorCode, IDNumber, isSyariah, Tenor, OutstandingPrincipal, Bucket, ContractStatus, NoPolicy, NoRangka, NoMesin, DebtorAddress, DebtorDisctrict, DebtorSubDistrict, DebtorCity, DebtorPostalCode, CreatedAt, CreatedBy, UpdatedAt, UpdatedBy").
		Find(&rstcustomers).Error

	if err != nil {
		return rstcustomers, err
	}

	return rstcustomers, nil

}
