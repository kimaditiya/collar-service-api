package user

type UserFormatter struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	NIK      string `json:"nik"`
	Token    string `json:"token"`
	Email    string `json:"email"`
	CoreUser string `json:"core_user"`
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

type UserRolesFormatter struct {
	// ID int `json:"id"`
	// UserID   int `json:"userId"`
	UserID   int `json:"user_id"`
	RoleCode int `json:"rolecode"`
}

func FormatUserRoles(userRoles UserRoles) UserRolesFormatter {
	formatter := UserRolesFormatter{
		// ID:       userRoles.ID,
		UserID:   userRoles.UserID,
		RoleCode: userRoles.RoleCode,
	}
	return formatter
}
