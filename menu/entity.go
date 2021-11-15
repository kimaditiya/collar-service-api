package menu

import "time"

type Menu struct {
	ID         int
	MenuCode   string
	MenuTitle  string
	Url        string
	Created_by string
	Created_at time.Time
	Update_by  string
	Update_at  time.Time
	MenuIcons  string
}
