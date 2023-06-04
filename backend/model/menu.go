package model

type Menu struct {
	MenuID      int     `json:"menu_id"`
	MenuName    string  `json:"menu_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Picture     string  `json:"picture"`
	CategoryID  int     `json:"category_id"`
}
