package dkh

type CreateDkhFormatter struct {
	Dkh []Dkh `json:"Dkh"`
}

func FormatCreateDkh(dkh []Dkh) CreateDkhFormatter {
	formatter := CreateDkhFormatter{
		Dkh: dkh,
	}
	return formatter
}
