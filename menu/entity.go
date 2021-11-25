package menu

import (
	"time"
)

type Menu struct {
	// gorm.Model
	ID        int       `gorm:"column:id" json:"ID"`
	MenuCode  string    `gorm:"column:menu_code" json:"MenuCode"`
	MenuTitle string    `gorm:"column:menu_title" json:"MenuTitle"`
	Url       string    `gorm:"column:url" json:"Url"`
	CreatedBy string    `gorm:"column:created_by" json:"CreatedBy"`
	CreatedAt time.Time `gorm:"column:created_at" json:"CreatedAt"`
	UpdateBy  string    `gorm:"column:updated_by" json:"UpdateBy"`
	// UpdateAt  time.Time `gorm:"column:updated_at" json:"UpdateAt"`
	MenuIcons string `gorm:"column:menu_icons" json:"MenuIcons"`
	ParentID  int    `column:"parent_id"`
	MenuChild []Menu `gorm:"foreignKey:ParentID;" json:"MenuChild"`
	// DeletedAt time.Time
	// `gorm:"many2many:menu_parent;"`
}

type GetAllMenuParent struct {
	ID        int
	MenuTitle string
}
