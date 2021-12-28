package wilayah

type CreateWilayahFormatter struct {
	Wilayah Wilayah `json:"Wilayah"`
}

func FormatCreateWilayah(wilayah Wilayah) CreateWilayahFormatter {
	formatter := CreateWilayahFormatter{
		Wilayah: wilayah,
	}
	return formatter
}
