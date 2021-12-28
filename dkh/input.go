package dkh

type CreateDkhInput struct {
	Dkh []Dkh `json:"Dkh" binding:"required"`
}
