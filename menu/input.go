package menu

type MenuInput struct {
	MenuCode  string `json:"menu_code" binding:"required"`
	MenuTitle string `json:"menu_title" binding:"required"`
	Url       string `json:"url" binding:"required"`
	MenuIcons string `json:"menu_icons"`
	ParentID  int    `json:"parent_id"`
}
