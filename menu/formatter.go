package menu

type MenuFormatter struct {
	ID             int    `json:"id"`
	MenuCode       string `json:"menu_code"`
	MenuTitle      string `json:"menu_title"`
	Url            string `json:"url"`
	MenuIcons      string `json:"menu_icons"`
	MenuTitleChild string `json:"menu_title_child"`
	UrlChid        string `json:"url_child"`
}

func FormatMenu(menu Menu) MenuFormatter {
	formater := MenuFormatter{
		ID:        menu.ID,
		MenuCode:  menu.MenuCode,
		MenuTitle: menu.MenuTitle,
		Url:       menu.Url,
		MenuIcons: menu.MenuIcons,
	}
	return formater
}
