package wilayah

type CreateWilayah struct {
	Wilayah Wilayah `json:"Wilayah" binding:"required"`
}
