package customers

import "time"

type Customers struct {
	RYM                  string
	BranchID             string
	ApplicationID        string
	AgreementNo          string
	FullName             string
	BranchFullName       string
	Periode              int
	Product              string
	ProductID            string
	Model                string
	DueDate              time.Time
	InstallmentAmount    float64
	LateChargeAmount     float64
	AmountTobePaid       float64
	OvdDays              int
	CollectorCode        string
	IDNumber             string
	IsSyariah            int
	Tenor                int
	OutstandingPrincipal float64
	Bucket               string
	ContractStatus       string
	NoPolicy             string
	NoRangka             string
	NoMesin              string
	DebtorAddress        string
	DebtorDisctrict      string
	DebtorSubDistrict    string
	DebtorCity           string
	DebtorPostalCode     string
	CreatedAt            time.Time
	CreatedBy            string
	UpdatedAt            time.Time
	UpdatedBy            string
}
