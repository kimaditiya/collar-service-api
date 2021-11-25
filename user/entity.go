package user

import "time"

type User struct {
	ID         int
	FullName   string
	NIK        string
	Status     int
	Password   string
	Email      string
	Created_by string
	Created_at time.Time
	Update_by  string
	Update_at  time.Time
	CoreUser   string
}
type UserRoles struct {
	ID         int
	UserID     int
	RoleCode   int
	Created_by string
	Created_at time.Time
	Update_by  string
	Update_at  time.Time
}
