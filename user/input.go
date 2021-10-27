package user

type RegisterUserInput struct {
	FullName string `json:"fullname" binding:"required"`
	NIK      string `json:"nik" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}
