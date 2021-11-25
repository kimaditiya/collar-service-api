package menu

import "time"

type MenuFormatter struct {
	MenuID    int       `json:"id"`
	MenuCode  string    `json:"menu_code"`
	MenuTitle string    `json:"menu_title"`
	Url       string    `json:"url"`
	MenuIcons string    `json:"menu_icons"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
	CreatedBy string    `json:"created_by"`
	// UpdateAt  time.Time `gorm:"default:CURRENT_TIMESTAMP;column:update_at" json:"update_at"`
	UpdateBy string `json:"update_by"`
	ParentID int    `json:"parent_id"`
}

type GetAllMenuParentFormatter struct {
	MenuParent []GetAllMenuParent
}

func FormatMenu(menu Menu) MenuFormatter {
	formater := MenuFormatter{
		MenuID:    menu.ID,
		MenuCode:  menu.MenuCode,
		MenuTitle: menu.MenuTitle,
		Url:       menu.Url,
		MenuIcons: menu.MenuIcons,
		CreatedAt: menu.CreatedAt,
		CreatedBy: menu.CreatedBy,
		// UpdateAt:  menu.UpdateAt,
		UpdateBy: menu.UpdateBy,
		ParentID: menu.ParentID,
	}

	return formater
}

func FormatGetAllMenuParent(getAllMenuParent GetAllMenuParent) GetAllMenuParentFormatter {
	formatter := GetAllMenuParentFormatter{
		MenuParent: []GetAllMenuParent{},
	}
	return formatter
}
