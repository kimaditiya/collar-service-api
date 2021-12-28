package assign

type CreateAssign struct {
	Assign []Assign `json:"Assign" binding:"required"`
}
