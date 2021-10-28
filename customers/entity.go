package customers

import "time"

type Customers struct {
	Dd                     int
	BranchID               string
	ReferenceAct           string
	ApplicationID          string
	AgreementNo            string
	FullName               string
	BranchFullName         string
	UserName               string
	PtpDate                string
	AssetType              string
	DueDate                time.Time
	Ovd                    string
	Tenor                  string
	InstallmentOf          string
	InstallmentAmount      string
	OutstandingPrincipal   string
	Late_cash              string
	Asset_mode             string
	Asset_year             int
	Process_date           time.Time
	PaidDate               string
	ContractStatus         string
	RRDDate                string
	DebtorPostalCode       string
	DebtorResidenceAddress string
	AssignmentCollGroup    string
	Kode_post              int
	CollectorId            int
	CollectorType          string
	City                   string
	Kecamatan              string
	Kelurahan              string
	Created_by             string
	Created_at             time.Time
	Update_by              string
	Update_at              time.Time
}
