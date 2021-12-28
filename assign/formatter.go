package assign

type CreateAssignInputFormatter struct {
	Assign []Assign `json:"Assign"`
}

func FormatCreateAssign(assign []Assign) CreateAssignInputFormatter {
	formatter := CreateAssignInputFormatter{
		Assign: assign,
	}
	return formatter
}
