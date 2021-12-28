package customers

import (
	"time"

	"gorm.io/gorm"
)

// buat interface untuk user
type Repository interface {
	// ListOfCustomer(agreementno string) ([]Customers, error)
	ListOfCustomer() ([]Customers, error)
	FindLisOfCustomerByProduct(Product string) ([]Customers, error)
	FindListOfCustomerByOvd(OvdDays int) ([]Customers, error)
	FindListOfCustomerByDueDate(DueDate time.Time) ([]Customers, error)

	// Service Mobile
	ListOfCustomerColl() ([]Customers, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) ListOfCustomer(AgreementNo string) ([]Customers, error) {
func (r *repository) ListOfCustomer() ([]Customers, error) {

	//var rstcustomers Customers
	var rstcustomers []Customers

	err := r.db.Debug().Table("vw_listloc").Find(&rstcustomers).Error

	if err != nil {
		return rstcustomers, err
	}

	return rstcustomers, nil

}

func (r *repository) FindLisOfCustomerByProduct(product string) ([]Customers, error) {
	var listOfCustomerByProduct []Customers

	err := r.db.Where("Product = ?", product).Table("vw_listloc").Find(&listOfCustomerByProduct).Error
	if err != nil {
		return listOfCustomerByProduct, err
	}

	return listOfCustomerByProduct, nil
}

func (r *repository) FindListOfCustomerByOvd(OvdDays int) ([]Customers, error) {
	var listOfCustomerByOvd []Customers

	err := r.db.Where("OvdDays = ?", OvdDays).Table("vw_listloc").Find(&listOfCustomerByOvd).Error
	if err != nil {
		return listOfCustomerByOvd, err
	}

	return listOfCustomerByOvd, nil
}

func (r *repository) FindListOfCustomerByDueDate(DueDate time.Time) ([]Customers, error) {
	var listOfCustomerByDueDate []Customers

	err := r.db.Where("DueDate = ?", DueDate).Table("vw_listloc").Find(&listOfCustomerByDueDate).Error
	if err != nil {
		return listOfCustomerByDueDate, err
	}

	return listOfCustomerByDueDate, nil
}

func (r *repository) ListOfCustomerColl() ([]Customers, error) {

	var rstcustomers []Customers

	err := r.db.Debug().Table("vw_listloc").Find(&rstcustomers).Error

	if err != nil {
		return rstcustomers, err
	}

	return rstcustomers, nil

}
