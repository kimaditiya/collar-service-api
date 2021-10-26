package user

type UserFormatter struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	NIK      string `json:"nik"`
	Token    string `json:"token"`
	Email    string `json:"email"`
}

func FormatUser(user User, token string) UserFormatter {
	formater := UserFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		NIK:      user.NIK,
		Token:    token,
		Email:    user.Email,
	}
	return formater
}
