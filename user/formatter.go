package user

type UserFormatter struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	NIK      string `json:"nik"`
	Token    string `json:"token"`
	Email    string `json:"email"`
	CoreUser string `json:"coreuser"`
}

func FormatUser(user User, token string) UserFormatter {
	formater := UserFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		NIK:      user.NIK,
		Token:    token,
		CoreUser: user.CoreUser,
		Email:    user.Email,
	}
	return formater
}
