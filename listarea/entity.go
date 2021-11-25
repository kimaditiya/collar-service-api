package listarea

import (
	"time"
)

type Area struct {
	ID           int    `json:"id"`
	AreaID       string `gorm:"column:areaID"json:"AreaID"`
	AreaFullName string `gorm:"column:areaFullName" json:"AreaFullName"`
	// IsActive     int
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
