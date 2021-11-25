package user

type RegisterUserInput struct {
	FullName string `json:"fullname" binding:"required"`
	NIK      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	CoreUser string `json:"coreuser" binding:"required"`
}

type LoginInput struct {
	NIK      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRolesInput struct {
	// ID       int `json:"id" binding:"required"`
	UserID   int `json:"user_id" binding:"required"`
	RoleCode int `json:"rolecode" binding:"required"`
}
