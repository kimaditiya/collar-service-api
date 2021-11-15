package customers

import "time"

type FindCustomer struct {
	AgreementNo  string `json:"agreementno"`
	CustomerName string `json:"cust_name"`
}

type FilterCustomerByProduct struct {
	Product string `json:"Product"`
}

type FilterListCustomerByOvd struct {
	OvdDays int `json:"OvdDays"`
}

type FilterListCustomerByDueDate struct {
	DueDate time.Time `json:"DueDate"`
}
